use anyhow::{anyhow, Result};
use flume::Sender;
use tracing::{debug, warn};
use std::sync::Arc;
use tokio::sync::RwLock;
use std::collections::HashMap;
use crate::{
    // cupti::CuptiTflopsCalculator,
    get_accounting_mode, monitor::{MonitoringSession, LogEntry}, set_accounting_mode, nvml::NVML
};
use std::time::Duration;

// The GPU service acts as a factory for creating monitoring and measurement instances
pub struct GpuService {
    resource_provider_id: String,
    log_tx: Sender<LogEntry>,
    // Wrap RwLocks in Arc
    active_monitors: Arc<RwLock<HashMap<String, Arc<RwLock<MonitoringSession>>>>>,
    // active_measurements: Arc<RwLock<HashMap<String, Arc<RwLock<CuptiTflopsCalculator>>>>>,
}

impl GpuService {
    pub async fn init(log_tx: Sender<LogEntry>) -> Result<Self> {
        debug!("Checking GPU accounting mode");
        let accounting_mode = get_accounting_mode()?;

        if accounting_mode != "Enabled" {
            debug!("Enabling GPU accounting mode");
            set_accounting_mode()?;
            // Wait a bit for the changes to take effect
            tokio::time::sleep(Duration::from_millis(100)).await;
            
            // Verify accounting mode was enabled
            let mode = get_accounting_mode()?;
            if mode != "Enabled" {
                return Err(anyhow!("Failed to enable GPU accounting mode"));
            } else {
                debug!("GPU accounting mode enabled");
            }

            // Explicitly enable accounting on the device
            if let Ok(nvml) = NVML.as_ref() {
                if let Ok(mut device) = nvml.device_by_index(0) {
                    if let Err(e) = device.set_accounting(true) {
                        warn!("Failed to enable accounting on device: {}. Will try with sudo.", e);
                        // Try with sudo if direct call fails
                        std::process::Command::new("sudo")
                            .arg("nvidia-smi")
                            .arg("-i")
                            .arg("0")
                            .arg("--accounting-mode=1")
                            .output()
                            .map_err(|e| anyhow!("Failed to enable device accounting: {}", e))?;
                    }
                }
            }
        }

        debug!("Initializing GPU service");
        Ok(Self {
            resource_provider_id: "".to_string(),
            log_tx,
            active_monitors: Arc::new(RwLock::new(HashMap::new())),
            // active_measurements: Arc::new(RwLock::new(HashMap::new())),
        })
    }

    pub async fn set_resource_provider_id(&mut self, resource_provider_id: String) {
        self.resource_provider_id = resource_provider_id;
    }

    pub async fn create_monitoring_session(&self, deal_id: String) -> Result<Arc<RwLock<MonitoringSession>>> {
        let mut monitors = self.active_monitors.write().await;
        let key = format!("{}:{}", self.resource_provider_id, deal_id);
        
        if monitors.contains_key(&key) {
            return Err(anyhow::anyhow!("Monitoring session already exists for this deal"));
        }

        let session = Arc::new(RwLock::new(MonitoringSession::new(
            self.log_tx.clone(),
            self.resource_provider_id.clone(),
            deal_id.clone(),
        )));
        
        monitors.insert(key, session.clone());
        Ok(session)
    }

    // pub async fn create_measurement_session(&self, deal_id: String) -> Result<Arc<RwLock<CuptiTflopsCalculator>>> {
    //     let calculator = Arc::new(RwLock::new(CuptiTflopsCalculator::new(
    //         self.log_tx.clone(),
    //         self.resource_provider_id.clone(),  
    //         deal_id.clone(),
    //     )));
    //     let key = format!("{}:{}", self.resource_provider_id, deal_id);
    //     measurements.insert(key, calculator.clone());
    //     Ok(calculator)
    // }

    pub async fn get_monitoring_session(&self, deal_id: String) -> Result<Arc<RwLock<MonitoringSession>>> {
        let monitors = self.active_monitors.read().await;
        let key = format!("{}:{}", self.resource_provider_id, deal_id);
        Ok(monitors.get(&key)
            .ok_or_else(|| anyhow::anyhow!("No active monitoring session found"))?
            .clone())
    }

    // pub async fn get_measurement_session(&self, deal_id: String) -> Result<Arc<RwLock<CuptiTflopsCalculator>>> {
    //     let measurements = self.active_measurements.read().await;
    //     let key = format!("{}:{}", self.resource_provider_id, deal_id);
    //     Ok(measurements.get(&key)
    //         .ok_or_else(|| anyhow::anyhow!("No active measurement session found"))?
    //         .clone())
    // }

    pub async fn stop_monitoring(&self, deal_id: String) -> Result<()> {
        let mut monitors = self.active_monitors.write().await;
        let key = format!("{}:{}", self.resource_provider_id, deal_id);
        
        if let Some(session) = monitors.remove(&key) {
            session.write().await.stop().await;
            Ok(())
        } else {
            Err(anyhow::anyhow!("No active monitoring session found"))
        }
    }

    // pub async fn stop_measurement(&self, deal_id: String) -> Result<()> {
    //     let mut measurements = self.active_measurements.write().await;
    //     let key = format!("{}:{}", self.resource_provider_id, deal_id);
        
    //     measurements.remove(&key)
    //         .ok_or_else(|| anyhow::anyhow!("No active measurement session found"))?;
    //     Ok(())
    // }
}

impl Drop for GpuService {
    fn drop(&mut self) {
        // Clone what we need before the move
        let monitors = self.active_monitors.clone();
        // let measurements = self.active_measurements.clone();

        if let Ok(handle) = tokio::runtime::Handle::try_current() {
            let _ = handle.spawn_blocking(move || {
                let rt = tokio::runtime::Runtime::new().unwrap();
                rt.block_on(async {
                    let mut monitors = monitors.write().await;
                    for (_, session) in monitors.drain() {
                        let _ = session.write().await.stop().await;
                    }

                    // let mut measurements = measurements.write().await;
                    // for (_, calculator) in measurements.drain() {
                    //     drop(calculator.write().await);
                    // }
                });
            });
        }
    }
}

impl GpuService {
    // Add cleanup method for graceful shutdown
    pub async fn shutdown(&self) -> Result<()> {
        // Stop all monitoring sessions
        let mut monitors = self.active_monitors.write().await;
        for (_, session) in monitors.drain() {
            session.write().await.stop().await;
        }

        // Stop all measurement sessions
        // let mut measurements = self.active_measurements.write().await;
        // measurements.clear();

        Ok(())
    }
}
