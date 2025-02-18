// External crates
use candle_core::{quantized::gguf_file, Device, Tensor};
use candle_examples::token_output_stream::TokenOutputStream;
use candle_transformers::{
    generation::{LogitsProcessor, Sampling},
    models::quantized_phi3::ModelWeights,
};
use futures::Stream;
use hf_hub::{api::sync::Api, Repo, RepoType};
use tokenizers::Tokenizer;
use tokio_stream::wrappers::ReceiverStream;
// Standard library
use std::marker::PhantomData;
use tokio::sync::mpsc;

const MAX_HISTORY_SIZE: usize = 1024;

#[derive(Clone)]
pub struct Model<T: ModelArgs + Clone + 'static> {
    model: ModelWeights,
    pub tokenizer: Tokenizer,
    pub device: Device,
    chat_history: Vec<u32>,
    _phantom: PhantomData<T>,
}

fn device() -> anyhow::Result<Device> {
    match Device::cuda_if_available(0) {
        Ok(_) => Ok(Device::new_cuda(0)?),
        Err(_) => anyhow::bail!("CUDA device required but not available"),
    }
}

pub trait ModelArgs: Send + Sync {
    fn default_model_path() -> anyhow::Result<std::path::PathBuf> {
        let api = Api::new().map_err(anyhow::Error::msg)?;
        let api = api.repo(Repo::with_revision(
            "microsoft/Phi-3-mini-4k-instruct-gguf".to_string(),
            RepoType::Model,
            "main".to_string(),
        ));
        api.get("Phi-3-mini-4k-instruct-q4.gguf")
            .map_err(anyhow::Error::msg)
    }

    fn default_tokenizer_path() -> anyhow::Result<std::path::PathBuf> {
        let api = Api::new().map_err(anyhow::Error::msg)?;
        let api = api.model("microsoft/Phi-3-mini-4k-instruct".to_string());
        api.get("tokenizer.json").map_err(anyhow::Error::msg)
    }

    fn model_path(&self) -> anyhow::Result<std::path::PathBuf> {
        Self::default_model_path()
    }

    fn tokenizer_path(&self) -> anyhow::Result<std::path::PathBuf> {
        Self::default_tokenizer_path()
    }

    fn sample_len(&self) -> usize {
        1000
    }

    fn temperature(&self) -> f64 {
        0.8
    }

    fn top_k(&self) -> Option<usize> {
        None
    }

    fn top_p(&self) -> Option<f64> {
        None
    }

    fn seed(&self) -> u64 {
        299792458
    }

    fn split_prompt(&self) -> bool {
        false
    }

    fn repeat_penalty(&self) -> f32 {
        1.1
    }

    fn repeat_last_n(&self) -> usize {
        64
    }
}

impl<T: ModelArgs + Clone + 'static> Model<T> {
    pub async fn new(args: &T) -> anyhow::Result<Self> {
        let model_path = args.model_path()?;
        let mut file = std::fs::File::open(&model_path)?;
        let device = device()?;

        let model = {
            let model = gguf_file::Content::read(&mut file)
                .map_err(|e| anyhow::anyhow!("Failed to read model: {}", e))?;
            ModelWeights::from_gguf(false, model, &mut file, &device)
                .map_err(|e| anyhow::anyhow!("Failed to load model weights: {}", e))?
        };

        let tokenizer = Tokenizer::from_file(args.tokenizer_path()?)
            .map_err(|e| anyhow::anyhow!("Failed to load tokenizer: {}", e))?;

        Ok(Model {
            model,
            tokenizer,
            device,
            chat_history: Vec::with_capacity(1024),
            _phantom: PhantomData,
        })
    }

    fn get_history_text(&self) -> anyhow::Result<String> {
        if self.chat_history.is_empty() {
            return Ok(String::new());
        }
        self.tokenizer
            .decode(&self.chat_history, true)
            .map_err(anyhow::Error::msg)
    }

    pub fn update_history(&mut self, tokens: &[u32]) {
        self.chat_history.extend_from_slice(tokens);
        if self.chat_history.len() > MAX_HISTORY_SIZE {
            let new_start = self.chat_history.len() - MAX_HISTORY_SIZE;
            self.chat_history.drain(0..new_start);
        }
    }

    pub async fn generate_stream_from_prompt(
        &mut self,
        prompt: String,
        args: &T,
    ) -> anyhow::Result<impl Stream<Item = String>> {
        let (tx, rx) = mpsc::channel(32);

        // Clone only what's needed
        let model = self.model.clone();
        let device = self.device.clone();
        let args = args.clone();
        let tokenizer = self.tokenizer.clone();

        // Format and encode prompt
        let history_text = self.get_history_text()?;
        tracing::debug!("History text: {}", history_text);
        let formatted_prompt = if history_text.is_empty() {
            format!("<|user|>{}<|end|><|assistant|>", prompt)
        } else {
            format!("{}<|user|>{}<|end|><|assistant|>", history_text, prompt)
        };

        let tokens = tokenizer
            .encode(formatted_prompt, true)
            .map_err(|e| anyhow::anyhow!("{}", e))?;
        let prompt_tokens = tokens.get_ids().to_vec();

        // Update history with prompt tokens
        self.update_history(&prompt_tokens);

        // Spawn independent generation task
        tokio::spawn(async move {
            let _ = run_generation(model, device, tokenizer, prompt_tokens, args, tx).await;
        });

        Ok(ReceiverStream::new(rx))
    }
}

async fn run_generation<T: ModelArgs + Send + 'static>(
    model: ModelWeights,
    device: Device,
    tokenizer: Tokenizer,
    prompt_tokens: Vec<u32>,
    args: T,
    tx: mpsc::Sender<String>,
) -> anyhow::Result<Vec<u32>> {
    let mut tos = TokenOutputStream::new(tokenizer);
    let tokens = generate_tokens(
        &model,
        &device,
        &prompt_tokens[..],
        args,
        Some(tx.clone()),
        &mut tos,
    )
    .await
    .map_err(|e| anyhow::anyhow!("{}", e))?;

    tracing::debug!("Generation completed successfully");
    Ok(tokens)
}

async fn generate_tokens<T: ModelArgs>(
    model: &ModelWeights,
    device: &Device,
    prompt_tokens: &[u32],
    args: T,
    tx: Option<mpsc::Sender<String>>,
    tos: &mut TokenOutputStream,
) -> Result<Vec<u32>, Box<dyn std::error::Error + Send + Sync + 'static>> {
    let mut model = model.clone();
    let mut all_tokens = Vec::with_capacity(args.sample_len());
    let to_sample = args.sample_len().saturating_sub(1);

    // Get the EOS token ID once
    let eos_token = *tos
        .tokenizer()
        .get_vocab(true)
        .get("<|endoftext|>")
        .unwrap();

    let mut logits_processor = {
        let temperature = args.temperature();
        let sampling = if temperature <= 0. {
            Sampling::ArgMax
        } else {
            match (args.top_k(), args.top_p()) {
                (None, None) => Sampling::All { temperature },
                (Some(k), None) => Sampling::TopK { k, temperature },
                (None, Some(p)) => Sampling::TopP { p, temperature },
                (Some(k), Some(p)) => Sampling::TopKThenTopP { k, p, temperature },
            }
        };
        LogitsProcessor::from_sampling(args.seed(), sampling)
    };

    // Pre-allocate tensors where possible
    let input = Tensor::new(prompt_tokens, device)?.unsqueeze(0)?;
    let next_token = {
        let logits = model.forward(&input, 0)?;
        logits_processor.sample(&logits.squeeze(0)?)?
    };

    all_tokens.push(next_token);
    if let Some(t) = tos.next_token(next_token)? {
        if let Some(tx) = &tx {
            tx.send(t)
                .await
                .map_err(|e| anyhow::anyhow!("Send error: {}", e))?;
        }
    }

    // Generate remaining tokens
    let mut next_token = next_token; // Keep track of the current token

    for index in 0..to_sample {
        let input = Tensor::new(&[next_token], device)?.unsqueeze(0)?;
        let next_pos = prompt_tokens.len() + index;

        next_token = {
            let logits = model.forward(&input, next_pos)?;
            logits_processor.sample(&logits.squeeze(0)?)?
        };

        all_tokens.push(next_token);
        if let Some(t) = tos.next_token(next_token)? {
            if let Some(tx) = &tx {
                tx.send(t)
                    .await
                    .map_err(|e| anyhow::anyhow!("Send error: {}", e))?;
            }
        }

        if next_token == eos_token {
            break;
        }
    }

    Ok(all_tokens)
}
