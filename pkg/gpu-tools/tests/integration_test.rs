use anyhow::Result;
use clap::Parser;
use futures::StreamExt;
use gpu_tools::{GpuService, monitor::LOG_CHANNEL};
use tokio::time::Duration;
use tracing::{debug, warn};
use tracing_subscriber::{filter::LevelFilter, EnvFilter};

mod models;
use models::{model, phi3::ModelArgs};

#[derive(Clone)]
struct TestArgs;

impl ModelArgs for TestArgs {}

#[derive(Parser, Debug, Clone)]
struct InferenceArgs {
    /// The prompt to send to the model
    #[arg(long)]
    prompt: Option<String>,

    #[clap(flatten)]
    model_args: model::Args,
}

impl ModelArgs for InferenceArgs {
    fn model_path(&self) -> Result<std::path::PathBuf> {
        self.model_args.model_path()
    }

    fn tokenizer_path(&self) -> Result<std::path::PathBuf> {
        self.model_args.tokenizer_path()
    }

    fn sample_len(&self) -> usize {
        self.model_args.sample_len()
    }

    fn temperature(&self) -> f64 {
        self.model_args.temperature()
    }

    fn top_k(&self) -> Option<usize> {
        self.model_args.top_k()
    }

    fn top_p(&self) -> Option<f64> {
        self.model_args.top_p()
    }

    fn seed(&self) -> u64 {
        self.model_args.seed()
    }

    fn split_prompt(&self) -> bool {
        self.model_args.split_prompt()
    }

    fn repeat_penalty(&self) -> f32 {
        self.model_args.repeat_penalty()
    }

    fn repeat_last_n(&self) -> usize {
        self.model_args.repeat_last_n()
    }
}

#[cfg(test)]
mod tests {
    use gpu_tools::logger::LOGGER;

    use super::*;
    use std::io::Write;

    async fn setup_logger() -> Result<()> {
        debug!("Starting logger setup");
        dotenv::from_path("tests/.env").ok();

        let jwt_token = std::env::var("JWT_TESTRP_TOKEN").map_err(|e| anyhow::anyhow!("Failed to get JWT token: {}", e))?;

        match &*LOGGER {
            Ok(logger) => {
                if let Err(e) = logger.init_with_token(Some(&jwt_token)) {
                    debug!("Error initializing logger with token: {}", e);
                    return Err(anyhow::anyhow!("Failed to initialize logging"));
                } else {
                    debug!("Logger initialized successfully");
                    Ok(())
                }
            }
            Err(e) => {
                debug!("Error initializing logger: {}", e);
                return Err(anyhow::anyhow!("Failed to initialize logging"));
            }
        }
    }

    #[tokio::test]
    async fn test_inference_with_monitoring_and_measurement() {
        let filter = EnvFilter::from_default_env()
            .add_directive(LevelFilter::DEBUG.into())
            .add_directive("tokenizers=error".parse().unwrap());

        tracing_subscriber::fmt().with_env_filter(filter).init();

        debug!("Setting up logger");
        setup_logger().await.unwrap();
        // Create default arguments for testing using the trait defaults
        let test_args = TestArgs;
        let args = InferenceArgs {
            prompt: None,
            model_args: model::Args {
                model_path: None,     // Will use default_model_path from trait
                tokenizer_path: None, // Will use default_tokenizer_path from trait
                tracing: false,
                sample_len: test_args.sample_len(),
                temperature: test_args.temperature(),
                top_k: test_args.top_k(),
                top_p: test_args.top_p(),
                repeat_penalty: test_args.repeat_penalty(),
                repeat_last_n: test_args.repeat_last_n(),
                seed: test_args.seed(),
                split_prompt: test_args.split_prompt(),
            },
        };

        debug!(
            "avx: {}, neon: {}, simd128: {}, f16c: {}",
            candle_core::utils::with_avx(),
            candle_core::utils::with_neon(),
            candle_core::utils::with_simd128(),
            candle_core::utils::with_f16c()
        );
        debug!(
            "temp: {:.2} repeat-penalty: {:.2} repeat-last-n: {}",
            args.model_args.temperature,
            args.model_args.repeat_penalty,
            args.model_args.repeat_last_n
        );

        debug!("Getting model instance");
        let model = model::ModelManager::instance()
            .get_model("inference-service-session".to_string(), &args.model_args)
            .await
            .unwrap();
        debug!("Got model instance");
        let mut model = model.write().await;
        debug!("Got write lock on model");

        let prompt =
            "Explain the usage of prompt in the context of LLM text generation".to_string();

        let log_tx = LOG_CHANNEL.get_sender();

        let resource_provider_id = "sample_rp".to_string();
        let deal_id = "sample_deal".to_string();
        let mut gpu_service = GpuService::init(log_tx.clone())
            .await
            .unwrap();
        gpu_service.set_resource_provider_id(resource_provider_id).await;

        // Create monitoring session
        let monitor_session = gpu_service.create_monitoring_session(deal_id.clone()).await.unwrap();
        let mut monitor = monitor_session.write().await;
        monitor.start().await.unwrap();
        debug!("Monitoring started");

        // let calculator_session = gpu_service.create_measurement_session(deal_id.clone()).await.unwrap();
        // let mut calculator = calculator_session.write().await;
        // calculator.start_measurement().await.unwrap();
        // debug!("Measurement started");
        // Run inference
        debug!("Starting inference");
        let mut stream = model
            .generate_stream_from_prompt(prompt.clone(), &args.model_args)
            .await
            .unwrap();
        debug!("Got stream for inference run");

        while let Some(token) = stream.next().await {
            print!("{}", token);
            std::io::stdout().flush().unwrap();
        }
        println!();
        debug!("Inference completed");

        monitor.stop().await;
        debug!("Monitoring stopped");

        // let metrics = calculator.stop_measurement().await.unwrap();
        // debug!("Metrics: {:?}", metrics);
        // Ensure all logs are flushed
        if let Ok(logger) = LOGGER.as_ref() {
            if let Err(e) = logger.flush(Duration::from_secs(5)).await {
                warn!("Failed to flush logs: {}", e);
            }
        }
        // debug!("Measurement stopped");
        
        // // Drop the locks before stopping the sessions
        // drop(calculator);
        drop(monitor);
        
        // Properly stop sessions through service layer
        if let Err(e) = gpu_service.stop_monitoring(deal_id.clone()).await {
            warn!("Failed to stop monitoring session: {}", e);
        }
        // if let Err(e) = gpu_service.stop_measurement(deal_id.clone()).await {
        //     warn!("Failed to stop measurement session: {}", e);
        // }
        gpu_service.shutdown().await.unwrap();
        drop(gpu_service);
        debug!("Service sessions stopped");
    }

    #[tokio::test]
    async fn test_phi3() {
        let filter = EnvFilter::from_default_env()
            .add_directive(LevelFilter::DEBUG.into())
            .add_directive("tokenizers=error".parse().unwrap());

        tracing_subscriber::fmt().with_env_filter(filter).init();
        // Create default arguments for testing using the trait defaults
        let test_args = TestArgs;
        let args = InferenceArgs {
            prompt: None,
            model_args: model::Args {
                model_path: None,     // Will use default_model_path from trait
                tokenizer_path: None, // Will use default_tokenizer_path from trait
                tracing: false,
                sample_len: test_args.sample_len(),
                temperature: test_args.temperature(),
                top_k: test_args.top_k(),
                top_p: test_args.top_p(),
                repeat_penalty: test_args.repeat_penalty(),
                repeat_last_n: test_args.repeat_last_n(),
                seed: test_args.seed(),
                split_prompt: test_args.split_prompt(),
            },
        };

        debug!(
            "avx: {}, neon: {}, simd128: {}, f16c: {}",
            candle_core::utils::with_avx(),
            candle_core::utils::with_neon(),
            candle_core::utils::with_simd128(),
            candle_core::utils::with_f16c()
        );
        debug!(
            "temp: {:.2} repeat-penalty: {:.2} repeat-last-n: {}",
            args.model_args.temperature,
            args.model_args.repeat_penalty,
            args.model_args.repeat_last_n
        );

        let model = model::ModelManager::instance()
            .get_model("inference-service-session".to_string(), &args.model_args)
            .await
            .unwrap();
        let mut model = model.write().await;

        let prompt =
            "Explain the usage of prompt in the context of LLM text generation".to_string();

        let mut stream = model
            .generate_stream_from_prompt(prompt.clone(), &args.model_args)
            .await
            .unwrap();

        while let Some(token) = stream.next().await {
            print!("{}", token);
            std::io::stdout().flush().unwrap();
        }
        println!();
    }
}
