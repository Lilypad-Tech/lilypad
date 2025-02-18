use chrono;
use gpu_tools::gpu_tools::{
    gpu_tools_client::GpuToolsClient, InitRequest, StartProfilingRequest, StopProfilingRequest,
};
use gpu_tools::nvml::NVML;
use rdkafka::consumer::{Consumer, StreamConsumer};
use rdkafka::{ClientConfig, Message, TopicPartitionList};
use serde_json::Value;
use std::env;
use std::fs;
use std::io::Write;
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::time::Duration;
use std::time::Instant;
use tokio::sync::mpsc;
use tonic::transport::Channel;
use clap::Parser;
use anyhow::Result;
use futures::StreamExt;

mod models;
use models::{model, phi3::ModelArgs};

const KAFKA_BROKER: &str = "broker:9092";

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

async fn setup_client() -> GpuToolsClient<Channel> {
    GpuToolsClient::connect("http://localhost:50051")
        .await
        .expect("Failed to create client")
}

fn get_jwt_token() -> String {
    dotenv::from_path("tests/.env").ok();
    env::var("JWT_TESTRP_TOKEN").expect("JWT_TESTRP_TOKEN must be set in .env file")
}

async fn start_consumer() -> StreamConsumer {
    // First create base consumer for metadata check
    let client: rdkafka::consumer::BaseConsumer = ClientConfig::new()
        .set("bootstrap.servers", KAFKA_BROKER)
        .create()
        .expect("Failed to create client");

    // Check topic exists...
    let topic = "gpu-monitor";
    let metadata = client.fetch_metadata(Some(topic), Duration::from_secs(5))
        .expect("Failed to fetch metadata");

    if metadata.topics().is_empty() {
        println!("Topic {} does not exist, creating...", topic);
        // Create a separate admin client for topic creation
        let admin_client: rdkafka::admin::AdminClient<_> = ClientConfig::new()
            .set("bootstrap.servers", KAFKA_BROKER)
            .create()
            .expect("Failed to create admin client");

        let new_topic = rdkafka::admin::NewTopic::new(topic, 1, rdkafka::admin::TopicReplication::Fixed(1));
        match admin_client.create_topics(
            &[new_topic],
            &rdkafka::admin::AdminOptions::new(),
        ).await {
            Ok(_) => println!("Topic {} created successfully", topic),
            Err(e) => println!("Failed to create topic: {}", e),
        }
    }

    // Create the StreamConsumer for actual consumption
    let consumer: StreamConsumer = ClientConfig::new()
        .set("group.id", format!("test-consumer-{}", chrono::Utc::now().timestamp_millis()))
        .set("bootstrap.servers", KAFKA_BROKER)
        .set("enable.partition.eof", "false")
        .set("session.timeout.ms", "6000")
        .set("enable.auto.commit", "false")
        .create()
        .expect("Failed to create consumer");

    // Use consumer for remaining operations
    let mut tpl = TopicPartitionList::new();
    tpl.add_partition(topic, 0);
    consumer.assign(&tpl).expect("Failed to assign partition");
    
    // Wait for assignment to be ready
    tokio::time::sleep(Duration::from_secs(1)).await;

    // Try to fetch watermarks, but don't fail if we can't
    match consumer.fetch_watermarks(topic, 0, Duration::from_secs(5)) {
        Ok((low, high)) => {
            println!("Consumer ready - partition watermarks - low: {}, high: {}", low, high);
        }
        Err(e) => {
            println!("Warning: Could not fetch watermarks: {}", e);
            // Continue anyway as the topic might be empty
        }
    }

    // Set position to beginning
    let mut tpl = TopicPartitionList::new();
    tpl.add_partition_offset(topic, 0, rdkafka::Offset::Beginning)
        .expect("Failed to add partition offset");
    consumer.assign(&tpl).expect("Failed to assign with offset");

    consumer
}

async fn create_message_stream(
    consumer: StreamConsumer,
    is_running: Arc<AtomicBool>,
) -> mpsc::Receiver<Value> {
    let (tx, rx) = mpsc::channel(100);
    
    tokio::spawn(async move {
        println!("Starting message stream...");
        let mut message_count = 0;
        
        while is_running.load(Ordering::Relaxed) {
            match consumer.recv().await {
                Ok(msg) => {
                    if let Some(payload) = msg.payload() {
                        match serde_json::from_slice::<Value>(payload) {
                            Ok(json) => {
                                message_count += 1;
                                println!("Message {} received from offset {}", 
                                    message_count, msg.offset());
                                if tx.send(json).await.is_err() {
                                    break;
                                }
                            }
                            Err(e) => println!("Failed to parse message: {}", e),
                        }
                    }
                }
                Err(e) => {
                    println!("Error receiving message: {}", e);
                    tokio::time::sleep(Duration::from_millis(100)).await;
                }
            }
        }
        println!("Message stream stopped after {} messages", message_count);
    });

    rx
}

async fn try_gpu_operations() -> Result<(), Box<dyn std::error::Error + Send>> {
    println!("Entering try_gpu_operations");
    if let Ok(nvml) = NVML.as_ref() {
        println!("Got NVML reference");
        if let Ok(mut device) = nvml.device_by_index(0) {
            println!("Got GPU device");
            if let Err(e) = device.set_accounting(true) {
                println!("Error setting accounting: {:?}", e);
                let output = std::process::Command::new("nvidia-smi")
                    .arg("-i")
                    .arg("0")
                    .arg("--accounting-mode=1")
                    .output();
                println!("nvidia-smi command result: {:?}", output);
            }
        }
    } else {
        println!("Failed to get NVML reference");
    }

    let pid = std::process::id();
    let name = std::fs::read_to_string(format!("/proc/{}/comm", pid))
        .map(|s| s.trim().to_string())
        .unwrap_or_else(|_| String::from("unknown"));
    println!("Process name, pid: {}, {}", name, pid);

    // Create test args for Phi-3 model
    println!("Creating test args for Phi-3 model");
    let test_args = TestArgs;
    let args = InferenceArgs {
        prompt: Some("Explain the usage of prompt in the context of LLM text generation".to_string()),
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

    println!("Initializing Phi-3 model...");
    let model = model::ModelManager::instance()
        .get_model("test-session".to_string(), &args.model_args)
        .await
        .unwrap();
    let mut model = model.write().await;

    println!("Starting text generation...");
    let mut stream = model
        .generate_stream_from_prompt(args.prompt.unwrap(), &args.model_args)
        .await
        .unwrap();

    println!("\nGenerated text:");
    while let Some(token) = stream.next().await {
        print!("{}", token);
        std::io::stdout().flush().unwrap();
    }

    println!("\nGeneration completed.");

    Ok(())
}

#[tokio::test]
async fn test_profiling_workflow() {
    println!("Starting consumer setup...");
    let consumer = start_consumer().await;

    // Setup client and initialize
    let mut client = setup_client().await;
    let jwt_token = get_jwt_token();
    let unique_deal_id = format!("test-deal-{}", chrono::Utc::now().timestamp_millis());

    println!("Starting test with deal ID: {}", unique_deal_id);

    let init_response = client
        .init(InitRequest {
            rp_id: "test-rp-123".to_string(),
            jwt_token,
        })
        .await
        .expect("Failed to call init method");
    assert!(init_response.into_inner().success);

    // Create message stream
    let is_running = Arc::new(AtomicBool::new(true));
    let mut message_stream = create_message_stream(consumer, is_running.clone()).await;

    // Start profiling
    let start_response = client
        .start_profiling(StartProfilingRequest {
            deal_id: unique_deal_id.clone(),
        })
        .await
        .expect("Failed to start profiling");
    assert!(start_response.into_inner().success);
    println!("Profiling started successfully");

    // Run GPU operations and collect messages
    let mut messages = Vec::new();
    let mut gpu_task = Some(tokio::spawn(async move {
        println!("Starting GPU operations...");
        let result = try_gpu_operations().await;
        println!("GPU operations completed");
        tokio::time::sleep(Duration::from_secs(5)).await;
        result
    }));

    // Collect messages with better timeout handling
    let start_time = Instant::now();
    let timeout = Duration::from_secs(30); // Longer timeout for testing

    loop {
        tokio::select! {
            Some(msg) = message_stream.recv() => {
                messages.push(msg);
                println!("Collected {} messages", messages.len());
            }
            _result = gpu_task.as_mut().unwrap(), if gpu_task.is_some() => {
                println!("GPU task completed");
                // Wait a bit more for any final messages
                tokio::time::sleep(Duration::from_secs(2)).await;
                // Collect any remaining messages
                while let Ok(Some(msg)) = tokio::time::timeout(
                    Duration::from_millis(100),
                    message_stream.recv()
                ).await {
                    messages.push(msg);
                    println!("Collected {} messages", messages.len());
                }
                gpu_task = None;  // Clear the task
                break;
            }
            _ = tokio::time::sleep(Duration::from_secs(1)) => {
                if start_time.elapsed() > timeout {
                    println!("Test timeout reached");
                    break;
                }
            }
        }
    }

    // Stop profiling before stopping message collection
    let stop_response = client
        .stop_profiling(StopProfilingRequest {
            deal_id: unique_deal_id.clone(),
        })
        .await
        .expect("Failed to stop profiling");
    assert!(stop_response.into_inner().success);

    // Wait a bit and collect any final messages after stopping profiling
    tokio::time::sleep(Duration::from_secs(2)).await;
    while let Ok(Some(msg)) = tokio::time::timeout(
        Duration::from_millis(100),
        message_stream.recv()
    ).await {
        messages.push(msg);
        println!("Collected final message, total: {}", messages.len());
    }

    // Now safe to stop the message stream
    is_running.store(false, Ordering::Relaxed);

    let output_path = format!("test_results_{}.json", unique_deal_id);
    fs::write(
        &output_path,
        serde_json::to_string_pretty(&messages).unwrap(),
    ).expect("Failed to write messages to file");

    println!("Test completed:");
    println!("- Messages collected: {}", messages.len());
    println!("- Results saved to: {}", output_path);
}

