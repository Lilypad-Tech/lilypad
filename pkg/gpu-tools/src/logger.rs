// External crates
use anyhow::{self, Result};
use rdkafka::{
    message::{Header, OwnedHeaders},
    producer::{FutureProducer, FutureRecord},
    ClientConfig,
    config::RDKafkaLogLevel,
    producer::Producer,
};
use tracing::{debug, warn, error};
use flume::Receiver;
use std::sync::LazyLock;
use std::time::Duration;
use std::sync::Mutex;
use tokio::task::JoinHandle;
use std::sync::Arc;

use crate::monitor::{LogEntry, LOG_CHANNEL};
use crate::kafka::CustomContext;

pub static LOGGER: LazyLock<Result<Logger>> = LazyLock::new(|| {
    let kafka_broker = std::env::var("KAFKA_BROKER").unwrap_or_else(|_| "broker".to_string());
    let kafka_port = std::env::var("KAFKA_BROKER_PORT").unwrap_or_else(|_| "9094".to_string());
    let broker_address = format!("{}:{}", kafka_broker, kafka_port);
    Logger::new(&broker_address)
});

pub struct Logger {
    log_rx: Receiver<LogEntry>,
    producer: Arc<Mutex<FutureProducer<CustomContext>>>,
    logging_task: Arc<Mutex<Option<JoinHandle<()>>>>,
    broker: String,
}

impl Logger {
    pub fn new(broker: &str) -> Result<Self> {
        debug!("Initializing logger with broker: {}", broker);

        let config = ClientConfig::new();

        let context = CustomContext {
            token: String::new(),
        };

        let producer: FutureProducer<CustomContext> = config.create_with_context(context)?;

        let logger = Self {
            log_rx: LOG_CHANNEL.get_receiver(),
            producer: Arc::new(Mutex::new(producer)),
            logging_task: Arc::new(Mutex::new(None)),
            broker: broker.to_string(),
        };
        debug!("Logger initialized successfully");
        
        Ok(logger)
    }
    
    pub fn start_logging(&self) {
        let log_rx = self.log_rx.clone();
        let producer = self.producer.lock().unwrap().clone();
        debug!("Starting logging task");
        
        let handle = tokio::spawn(async move {
            debug!("Logging task started, waiting for messages");
            while let Ok(entry) = log_rx.recv_async().await {
                match Self::send_log_internal(producer.clone(), entry).await {
                    Ok(_) => (),
                    Err(e) => {
                        warn!("Failed to process log entry (broker may be down): {}", e);
                        continue;
                    }
                }
            }
            debug!("Logging task channel closed");
        });

        *self.logging_task.lock().unwrap() = Some(handle);
    }

    pub async fn stop_logging(&self) {
        if let Some(handle) = self.logging_task.lock().unwrap().take() {
            handle.abort();
            let _ = handle.await;
            debug!("Logging task stopped");
        }
    }

    // Helper function to create producer config
    fn create_config(&self, jwt_token: Option<&str>) -> ClientConfig {
        let mut config = ClientConfig::new();
        
        config.set("bootstrap.servers", &self.broker)
            .set("message.timeout.ms", "30000")
            .set("socket.timeout.ms", "30000")
            .set("socket.connection.setup.timeout.ms", "30000")
            .set("request.timeout.ms", "30000")
            .set("allow.auto.create.topics", "true")
            .set("message.send.max.retries", "0")
            .set("enable.idempotence", "false");

        if let Some(_) = jwt_token {
            debug!("Configuring SASL with JWT token");
            config
                .set("security.protocol", "SASL_PLAINTEXT")
                .set("sasl.mechanism", "OAUTHBEARER")
                .set("enable.ssl.certificate.verification", "false")
                .set_log_level(RDKafkaLogLevel::Debug);
        }

        config
    }

    // Helper function to create a new producer
    fn create_producer(&self, jwt_token: Option<&str>) -> Result<FutureProducer<CustomContext>> {
        let config = self.create_config(jwt_token);
        let context = if let Some(token) = jwt_token {
            CustomContext {
                token: token.to_string(),
            }
        } else {
            CustomContext {
                token: String::new(),
            }
        };

        let producer: FutureProducer<CustomContext> = config.create_with_context(context)?;

        // Test connection
        producer
            .client()
            .fetch_metadata(None, std::time::Duration::from_secs(5))
            .map_err(|e| anyhow::anyhow!("Failed to connect to Kafka broker: {}", e))?;

        Ok(producer)
    }

    // Initial setup with token
    pub fn init_with_token(&self, jwt_token: Option<&str>) -> Result<()> {
        let producer = self.create_producer(jwt_token)?;
        *self.producer.lock().unwrap() = producer;
        self.start_logging();
        debug!("Logger initialized with token");
        Ok(())
    }

    // Update existing token and restart logging
    pub fn update_jwt_token(&self, new_token: Option<&str>) -> Result<()> {
        debug!("Updating JWT token");
        let new_producer = self.create_producer(new_token)?;
        
        // Use block_in_place since we're in a sync context but need to call async
        tokio::task::block_in_place(|| {
            let rt = tokio::runtime::Handle::current();
            rt.block_on(async {
                self.restart_logging(new_producer).await
            })
        })?;
        
        debug!("JWT token updated successfully");
        Ok(())
    }

    // Internal method to handle producer updates
    async fn restart_logging(&self, new_producer: FutureProducer<CustomContext>) -> Result<()> {
        // Only flush and stop if there's an existing task
        if self.logging_task.lock().unwrap().is_some() {
            self.flush(Duration::from_secs(2)).await?;
            self.stop_logging().await;
        }
        
        *self.producer.lock().unwrap() = new_producer;
        self.start_logging();
        Ok(())
    }

    async fn send_log_internal(producer: FutureProducer<CustomContext>, entry: LogEntry) -> Result<()> {
        let payload = serde_json::to_string(&entry).map_err(|e| {
            error!("Failed to serialize log entry: {}", e);
            error!("Entry details: job_id={}, action={:?}, source={}, details={}", 
                entry.job_id, entry.action, entry.source, entry.details);
            anyhow::anyhow!("Failed to serialize log entry: {}", e)
        })?;

        // Create headers
        let mut headers = OwnedHeaders::new();
        // Add metadata as a single header
        if !entry.metadata.is_empty() {
            let metadata_str = serde_json::to_string(&entry.metadata)?;
            headers = headers.insert(Header {
                key: "metadata",
                value: Some(metadata_str.as_str()),
            });
        }

        let record = FutureRecord::to(&entry.topic)
            .payload(&payload)
            .key(&entry.key)
            .headers(headers);

        let delivery_result = producer.send(record, Duration::from_secs(2)).await;
        
        match delivery_result {
            Ok(_) => Ok(()),
            Err((kafka_error, _)) => {
                // Always output the entry details when broker is down
                warn!(
                    "[KAFKA_FALLBACK] Broker error: {}. Log entry details:\n\
                     Job ID: {}\n\
                     Action: {:?}\n\
                     Source: {}\n\
                     Details: {}\n\
                     Topic: {}\n\
                     Metadata: {:#?}",
                    kafka_error,
                    entry.job_id,
                    entry.action,
                    entry.source,
                    entry.details,
                    entry.topic,
                    entry.metadata
                );
                
                // Return error but continue processing
                Err(anyhow::anyhow!("Kafka error: {}", kafka_error))
            }
        }
    }

    pub async fn flush(&self, timeout: Duration) -> Result<()> {
        debug!("Flushing logger with timeout {:?}", timeout);
        match self.producer.lock().unwrap().flush(timeout) {
            Ok(_) => {
                debug!("Successfully flushed all messages");
                Ok(())
            },
            Err(e) => {
                warn!("Failed to flush messages: {}", e);
                Err(anyhow::anyhow!("Failed to flush messages: {}", e))
            }
        }
    }
}

impl Drop for Logger {
    fn drop(&mut self) {
        debug!("Logger being dropped, flushing messages");
        // Use blocking flush instead of async
        if let Err(e) = self.producer.lock().unwrap().flush(Duration::from_secs(1)) {
            warn!("Failed to flush messages during shutdown: {}", e);
        }
    }
}
