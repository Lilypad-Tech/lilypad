// External crates
use anyhow::{anyhow, Result};
use chrono::Utc;
use flume::{unbounded, Receiver, Sender};
use nvml_wrapper::{
    enum_wrappers::device::{
        Brand, Clock, ComputeMode, PerformanceState, TemperatureSensor, TemperatureThreshold,
    },
    enums::device::{DeviceArchitecture, PowerSource},
    error::NvmlError,
    high_level::{Event, EventLoopProvider},
    struct_wrappers::device::{AccountingStats, MemoryInfo, PciInfo, Utilization},
    structs::device::{PowerManagementConstraints, UtilizationInfo},
    Device as NvmlDevice,
};
use serde::{Deserialize, Serialize};
use serde_json;
use std::{
    collections::HashMap,
    sync::{Arc, LazyLock},
};
use tokio::{
    sync::oneshot,
    task::JoinHandle,
    time::{sleep, Duration},
};
use tracing::{debug, error, warn};

use crate::nvml::NVML;

pub static LOG_CHANNEL: LazyLock<LogChannel> = LazyLock::new(|| LogChannel::new());

const GPU_MONITOR_TOPIC: &str = "gpu-monitor";

#[derive(Debug, Serialize, Deserialize)]
pub struct LogEntry {
    // internal monitoring/logging job id
    pub job_id: String,

    pub topic: String, // Kafka topic
    pub key: String,   // Kafka key
    // Message content
    pub timestamp: String,
    pub action: MonitoringAction,
    pub source: String,
    pub details: String,

    #[serde(default)]
    pub metadata: HashMap<String, String>, // Additional headers for filtering
}

impl LogEntry {
    pub fn new(
        job_id: String,
        resource_provider_id: String,
        deal_id: String,
        action: MonitoringAction,
        source: String,
        details: String,
    ) -> Self {
        let mut metadata = HashMap::new();
        metadata.insert(
            "resource_provider_id".to_string(),
            resource_provider_id.clone(),
        );
        metadata.insert("deal_id".to_string(), deal_id.clone());
        metadata.insert("job_id".to_string(), job_id.clone());
        Self {
            job_id,
            topic: GPU_MONITOR_TOPIC.to_string(),
            key: format!("{}:{}", resource_provider_id, deal_id),
            timestamp: Utc::now().to_rfc3339(),
            action,
            source,
            details,
            metadata,
        }
    }

    pub async fn send(self, sender: &Sender<LogEntry>) -> Result<()> {
        sender
            .send_async(self)
            .await
            .map_err(|e| anyhow::anyhow!("Failed to send log entry: {}", e))
    }
}

impl std::fmt::Display for LogEntry {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "Job ID: {}\nTopic: {}\nKey: {}\nTime: {}\nSource: {}\nAction: {:?}\nDetails:\n{}\nMetadata: {:?}", 
            &self.job_id,
            &self.topic,
            &self.key,
            &self.timestamp,
            &self.source,
            &self.action,
            &self.details,
            &self.metadata
        )
    }
}

#[derive(Clone)]
pub struct LogChannel {
    tx: Sender<LogEntry>,
    rx: Receiver<LogEntry>,
}

impl LogChannel {
    pub fn new() -> Self {
        let (tx, rx) = unbounded();
        Self { tx, rx }
    }

    pub fn get_sender(&self) -> Sender<LogEntry> {
        self.tx.clone()
    }

    pub fn get_receiver(&self) -> Receiver<LogEntry> {
        self.rx.clone()
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub enum MonitoringAction {
    StartInfo,
    StartState,
    StopInfo,
    StopState,
    AccountingStats,
    AccountingStatsNoStats,
    Metrics,
    ClockChange,
    PowerStateChange,
}

impl MonitoringAction {
    pub const START_INFO: &'static str = "monitoring_start_info";
    pub const START_STATE: &'static str = "monitoring_start_state";
    pub const STOP_INFO: &'static str = "monitoring_stop_info";
    pub const STOP_STATE: &'static str = "monitoring_stop_state";
    pub const ACCOUNTING_STATS: &'static str = "monitoring_accounting_stats";
    pub const ACCOUNTING_STATS_NO_STATS: &'static str = "monitoring_accounting_stats_no_stats";
    pub const METRICS: &'static str = "monitoring_metrics";
    pub const CLOCK_CHANGE: &'static str = "monitoring_clock_change";
    pub const POWER_STATE_CHANGE: &'static str = "monitoring_power_state_change";
}

impl ToString for MonitoringAction {
    fn to_string(&self) -> String {
        match self {
            Self::StartInfo => Self::START_INFO,
            Self::StartState => Self::START_STATE,
            Self::StopInfo => Self::STOP_INFO,
            Self::StopState => Self::STOP_STATE,
            Self::AccountingStats => Self::ACCOUNTING_STATS,
            Self::AccountingStatsNoStats => Self::ACCOUNTING_STATS_NO_STATS,
            Self::Metrics => Self::METRICS,
            Self::ClockChange => Self::CLOCK_CHANGE,
            Self::PowerStateChange => Self::POWER_STATE_CHANGE,
        }
        .to_string()
    }
}

pub struct MonitoringSession {
    job_id: String,
    monitor_handle: Option<JoinHandle<()>>,
    polling_metrics_handle: Option<JoinHandle<()>>,
    log_tx: Sender<LogEntry>,
    resource_provider_id: String,
    deal_id: String,
    nvml_device: Arc<NvmlDevice<'static>>,
    event_stop_tx: Option<oneshot::Sender<()>>,
    metrics_stop_tx: Option<oneshot::Sender<()>>,
    is_stopped: bool,
}

impl MonitoringSession {
    pub fn new(log_tx: Sender<LogEntry>, resource_provider_id: String, deal_id: String) -> Self {
        Self {
            job_id: nanoid::nanoid!(),
            monitor_handle: None,
            polling_metrics_handle: None,
            log_tx,
            resource_provider_id,
            deal_id,
            nvml_device: Arc::new(NVML.as_ref().unwrap().device_by_index(0).unwrap()),
            event_stop_tx: None,
            metrics_stop_tx: None,
            is_stopped: false,
        }
    }

    pub async fn start(&mut self) -> Result<()> {
        debug!("Cleaning up any existing handles");
        if self.monitor_handle.is_some() || self.polling_metrics_handle.is_some() {
            self.cleanup_handles_quiet().await?;
            self.is_stopped = false; // Reset stopped flag since we're starting again
        }

        debug!("Starting monitoring session");
        let gpu_info = match get_gpu_info(&self.nvml_device) {
            Ok(info) => info,
            Err(e) => {
                warn!("Failed to get GPU info: {}", e);
                return Ok(());
            }
        };

        debug!("Sending monitoring start log entry");
        let serializable_gpu_info = SerializableGpuDetails::from(&gpu_info);

        let details = match serde_json::to_string(&serializable_gpu_info) {
            Ok(json) => json,
            Err(e) => {
                warn!("Failed to serialize GPU info: {}", e);
                format!("{}", gpu_info) // Fallback to Display impl
            }
        };

        let log_entry = LogEntry::new(
            self.job_id.clone(),
            self.resource_provider_id.to_string(),
            self.deal_id.to_string(),
            MonitoringAction::StartInfo,
            self.nvml_device.uuid()?.to_string(),
            details,
        );
        log_entry.send(&self.log_tx).await?;

        let state = collect_device_state(&self.nvml_device)?;
        let serializable_state = SerializableDeviceState::from(&state);
        let details = match serde_json::to_string(&serializable_state) {
            Ok(json) => json,
            Err(e) => {
                warn!("Failed to serialize device state: {}", e);
                format!("{}", state) // Fallback to Display impl
            }
        };
        let log_entry = LogEntry::new(
            self.job_id.clone(),
            self.resource_provider_id.to_string(),
            self.deal_id.to_string(),
            MonitoringAction::StartState,
            self.nvml_device.uuid()?.to_string(),
            details,
        );
        log_entry.send(&self.log_tx).await?;

        // Create fresh oneshot channels
        let (event_stop_tx, event_stop_rx) = oneshot::channel();
        let (metrics_stop_tx, metrics_stop_rx) = oneshot::channel();
        self.event_stop_tx = Some(event_stop_tx);
        self.metrics_stop_tx = Some(metrics_stop_tx);

        self.start_event_monitor(event_stop_rx)?;
        self.start_metrics_polling(metrics_stop_rx)?;
        Ok(())
    }

    fn check_root_access() -> bool {
        unsafe { libc::geteuid() == 0 }
    }

    fn try_clear_accounting() -> Result<(), String> {
        // If we're root, let NVML handle it through the next device operation
        if Self::check_root_access() {
            return Ok(());
        }

        // If not root, try using sudo with nvidia-smi
        debug!("Not running as root, attempting to clear accounting with sudo");
        match std::process::Command::new("sudo")
            .arg("nvidia-smi")
            .arg("--accounting-mode=0")
            .output()
        {
            Ok(output) if output.status.success() => {
                debug!("Successfully cleared accounting with sudo");
                Ok(())
            }
            Ok(_) => Err("sudo command failed".to_string()),
            Err(e) => Err(format!("failed to execute sudo: {}", e)),
        }
    }

    fn get_accounting_stats(&self) -> Result<SerializableAccountingStats> {
        let device = Arc::clone(&self.nvml_device);
        if let Ok(accounting_pids) = device.accounting_pids() {
            debug!("Processes with accounting enabled: {}", accounting_pids.len());
            for pid in accounting_pids {
                if let Ok(stats) = device.accounting_stats_for(pid) {
                    let name = std::fs::read_to_string(format!("/proc/{}/comm", pid))
                        .map(|s| s.trim().to_string())
                        .unwrap_or_else(|_| String::from("unknown"));
                    debug!("Process name: {}, PID: {}", name, pid);
                    if name == "unknown" {
                        continue;
                    }
                    let gpu_util = stats.gpu_utilization.unwrap_or(0);
                    let mem_usage = stats.max_memory_usage.unwrap_or(0) / 1024 / 1024;
                    if gpu_util > 0 || mem_usage > 100 {
                        debug!("Retrieved accounting stats for PID {}: {} - GPU util: {}%, Memory: {}MB", pid, name, gpu_util, mem_usage);
                        return Ok(SerializableAccountingStats::new(pid, name, stats));
                    }
                } else {
                    warn!("Failed to get accounting stats for PID: {}", pid);
                }
            }
            warn!("No accounting stats found");
            Err(anyhow::anyhow!("No accounting stats found"))
        } else {
            warn!("Failed to get accounting pids");
            Err(anyhow::anyhow!("No accounting stats found"))
        }
    }

    pub fn start_event_monitor(&mut self, mut stop_rx: oneshot::Receiver<()>) -> Result<()> {
        let device = Arc::clone(&self.nvml_device);
        let log_tx = self.log_tx.clone();
        let job_id = self.job_id.clone();
        let resource_provider_id = self.resource_provider_id.to_string();
        let deal_id = self.deal_id.to_string();

        // Spawn the event loop in a blocking task to avoid blocking the async runtime
        let handle = tokio::task::spawn_blocking(move || {
            let rt = tokio::runtime::Handle::current();
            if let Err(e) = (|| {
                let mut event_loop = NVML
                    .as_ref()
                    .map_err(|e| anyhow::anyhow!("NVML error: {}", e))?
                    .create_event_loop(vec![&*device])?;

                event_loop.run_forever(|event_result, state| {
                    // Check if we received stop signal
                    if stop_rx.try_recv().is_ok() {
                        debug!("Received stop signal, interrupting event loop");
                        state.interrupt();
                        return;
                    }

                    match event_result {
                        Ok(event) => {
                            if let Err(e) = rt.block_on(handle_event(
                                job_id.clone(),
                                resource_provider_id.clone(),
                                deal_id.clone(),
                                event,
                                &log_tx,
                            )) {
                                warn!("Error handling event: {}", e);
                            }
                        }
                        Err(e) => match e {
                            NvmlError::Unknown => {}
                            _ => {
                                error!("NVML error: {}", e);
                                state.interrupt();
                            }
                        },
                    }
                });

                Ok::<(), anyhow::Error>(())
            })() {
                error!("Monitor error: {}", e);
            }
        });

        self.monitor_handle = Some(handle);
        Ok(())
    }

    pub fn start_metrics_polling(&mut self, mut stop_rx: oneshot::Receiver<()>) -> Result<()> {
        let device = Arc::clone(&self.nvml_device);
        let log_tx = self.log_tx.clone();
        let job_id = self.job_id.clone();
        let resource_provider_id = self.resource_provider_id.to_string();
        let deal_id = self.deal_id.to_string();

        // Spawn the polling loop in a blocking task to be consistent with event_loop handling
        let handle = tokio::task::spawn_blocking(move || {
            let rt = tokio::runtime::Handle::current();
            while !stop_rx.try_recv().is_ok() {
                // Use block_on since we're in a blocking task
                if let Err(e) = rt.block_on(poll_metrics(
                    job_id.clone(),
                    resource_provider_id.clone(),
                    deal_id.clone(),
                    &*device,
                    &log_tx,
                )) {
                    error!("Failed to poll metrics: {}", e);
                }
                // Sleep for 1 second
                std::thread::sleep(Duration::from_secs(1));
            }
            debug!("Received stop signal, interrupting metrics polling");
        });

        self.polling_metrics_handle = Some(handle);
        Ok(())
    }

    async fn cleanup_handles_quiet(&mut self) -> Result<()> {
        debug!("Performing quiet cleanup of handles");
        // Send stop signals if we have senders
        if let Some(stop_tx) = self.event_stop_tx.take() {
            let _ = stop_tx.send(());
        }
        if let Some(stop_tx) = self.metrics_stop_tx.take() {
            let _ = stop_tx.send(());
        }

        // Wait a bit for clean shutdown
        sleep(Duration::from_millis(100)).await;

        // Clean up handles
        if let Some(handle) = self.monitor_handle.take() {
            debug!("Cleaning up event monitor");
            handle.abort();
            let _ = handle.await;
        }
        if let Some(handle) = self.polling_metrics_handle.take() {
            debug!("Cleaning up polling metrics");
            handle.abort();
            let _ = handle.await;
        }

        // Wait a bit more to ensure everything is cleaned up
        sleep(Duration::from_millis(100)).await;
        Ok(())
    }

    async fn cleanup_handles(&mut self) -> Result<()> {
        debug!("Cleaning up handles with state logging");
        if let Ok(stats) = self.get_accounting_stats() {
            debug!("Accounting stats: {:?}", stats);
            let log_entry = LogEntry::new(
                self.job_id.clone(),
                self.resource_provider_id.to_string(),
                self.deal_id.to_string(),
                MonitoringAction::AccountingStats,
                self.nvml_device.uuid()?.to_string(),
                serde_json::to_string(&stats)?,
            );
            log_entry.send(&self.log_tx).await?;    
        } else {
            let log_entry = LogEntry::new(
                self.job_id.clone(),
                self.resource_provider_id.to_string(),
                self.deal_id.to_string(),
                MonitoringAction::AccountingStatsNoStats,
                self.nvml_device.uuid()?.to_string(),
                "No accounting stats found".to_string(),
            );
            log_entry.send(&self.log_tx).await?;
            warn!("Failed to get accounting stats");
        };

        // Create GPU info
        let gpu_info = match get_gpu_info(&self.nvml_device) {
            Ok(info) => info,
            Err(e) => {
                warn!("Failed to get GPU info: {}", e);
                return Ok(());
            }
        };

        let serializable_gpu_info = SerializableGpuDetails::from(&gpu_info);
        let details = match serde_json::to_string(&serializable_gpu_info) {
            Ok(json) => json,
            Err(e) => {
                warn!("Failed to serialize GPU info: {}", e);
                format!("{}", gpu_info)
            }
        };

        let log_entry = LogEntry::new(
            self.job_id.clone(),
            self.resource_provider_id.to_string(),
            self.deal_id.to_string(),
            MonitoringAction::StopInfo,
            self.nvml_device.uuid()?.to_string(),
            details,
        );
        log_entry.send(&self.log_tx).await?;

        // Perform the actual cleanup of handles
        self.cleanup_handles_quiet().await?;

        // Try to clear accounting stats
        match Self::try_clear_accounting() {
            Ok(_) => debug!("Successfully cleared GPU accounting stats"),
            Err(e) => warn!(
                "Failed to clear GPU accounting stats. Root access required: {}",
                e
            ),
        }

        let state = collect_device_state(&self.nvml_device)?;
        let serializable_state = SerializableDeviceState::from(&state);
        let log_entry = LogEntry::new(
            self.job_id.clone(),
            self.resource_provider_id.to_string(),
            self.deal_id.to_string(),
            MonitoringAction::StopState,
            self.nvml_device.uuid()?.to_string(),
            serde_json::to_string(&serializable_state)?,
        );
        log_entry.send(&self.log_tx).await?;

        Ok(())
    }

    pub async fn stop(&mut self) {
        if self.is_stopped {
            debug!("MonitoringSession already stopped");
            return;
        }
        debug!("Stopping monitoring session");
        if let Err(e) = self.cleanup_handles().await {
            error!("Error during cleanup: {}", e);
        }
        self.is_stopped = true;
    }

    pub fn stop_sync(&mut self) {
        debug!("Stopping monitoring session (sync)");
        if self.is_stopped {
            debug!("MonitoringSession already stopped");
            return;
        }
        // Send stop signals if we have senders
        if let Some(stop_tx) = self.event_stop_tx.take() {
            let _ = stop_tx.send(());
        }
        if let Some(stop_tx) = self.metrics_stop_tx.take() {
            let _ = stop_tx.send(());
        }

        // Clean up handles
        if let Some(handle) = self.monitor_handle.take() {
            debug!("Cleaning up event monitor");
            handle.abort();
        }
        if let Some(handle) = self.polling_metrics_handle.take() {
            debug!("Cleaning up polling metrics");
            handle.abort();
        }
        self.is_stopped = true;
    }
}

impl Drop for MonitoringSession {
    fn drop(&mut self) {
        self.stop_sync();
    }
}

impl std::fmt::Debug for MonitoringSession {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("MonitoringSession")
            .field("is_monitoring", &self.monitor_handle.is_some())
            .field("is_polling", &self.polling_metrics_handle.is_some())
            .finish()
    }
}
// Helper function for polling metrics
async fn poll_metrics(
    job_id: String,
    resource_provider_id: String,
    deal_id: String,
    device: &NvmlDevice<'static>,
    log_tx: &Sender<LogEntry>,
) -> Result<()> {
    let state = collect_device_state(device)?;
    let serializable_state = SerializableDeviceState::from(&state);
    let details = match serde_json::to_string(&serializable_state) {
        Ok(json) => json,
        Err(e) => {
            warn!("Failed to serialize device state: {}", e);
            format!("{}", state) // Fallback to Display impl
        }
    };
    let log_entry = LogEntry::new(
        job_id,
        resource_provider_id,
        deal_id,
        MonitoringAction::Metrics,
        device.uuid()?.to_string(),
        details,
    );

    log_entry.send(log_tx).await?;
    Ok(())
}

// Helper function for event handling
async fn handle_event<'a>(
    job_id: String,
    resource_provider_id: String,
    deal_id: String,
    event: Event<'a>,
    log_tx: &Sender<LogEntry>,
) -> Result<()> {
    match event {
        Event::ClockChange(device) => {
            let device_state = collect_device_state(&device)?;
            let serializable_state = SerializableDeviceState::from(&device_state);
            let details = match serde_json::to_string(&serializable_state) {
                Ok(json) => json,
                Err(e) => {
                    warn!("Failed to serialize device state: {}", e);
                    format!("{}", device_state) // Fallback to Display impl
                }
            };
            let log_entry = LogEntry::new(
                job_id,
                resource_provider_id,
                deal_id,
                MonitoringAction::ClockChange,
                device.uuid()?.to_string(),
                details,
            );
            log_entry.send(log_tx).await?;
        }
        Event::PowerStateChange(device) => {
            let device_state = collect_device_state(&device)?;
            let serializable_state = SerializableDeviceState::from(&device_state);
            let details = match serde_json::to_string(&serializable_state) {
                Ok(json) => json,
                Err(e) => {
                    warn!("Failed to serialize device state: {}", e);
                    format!("{}", device_state) // Fallback to Display impl
                }
            };
            let log_entry = LogEntry::new(
                job_id,
                resource_provider_id,
                deal_id,
                MonitoringAction::PowerStateChange,
                device.uuid()?.to_string(),
                details,
            );
            log_entry.send(log_tx).await?;
        }
        _ => warn!("Unhandled event: {:?}", event),
    }
    Ok(())
}

// GPU monitoring structures
#[derive(Debug, Clone)]
pub struct GpuDetails {
    pub name: String,
    pub uuid: String,
    pub pci_info: PciInfo,
    pub graphics_clock: u32,
    pub memory_clock: u32,
    pub sm_clock: u32,
    pub memory_info: MemoryInfo,
    pub perf_state: PerformanceState,
    pub architecture: String,
    pub cuda_cores: u32,
    pub max_pcie_link_gen: u32,
    pub max_pcie_link_width: u32,
    pub max_pcie_link_speed: Option<u64>,
    pub is_multi_gpu_board: bool,
    pub vbios_version: String,
    pub memory_bus_width: u32,
    pub power_management_limit_default: u32,
    pub power_management_limit: u32,
    pub power_management_limit_constraints: PowerManagementConstraints,
    pub brand: Brand,
    pub board_id: Option<String>,
    pub serial: Option<String>,
}

// Serializable version of GpuDetails for JSON output
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SerializableGpuDetails {
    pub name: String,
    pub uuid: String,
    pub graphics_clock: u32,
    pub memory_clock: u32,
    pub sm_clock: u32,
    pub perf_state: String,
    pub architecture: String,
    pub cuda_cores: u32,
    pub max_pcie_link_gen: u32,
    pub max_pcie_link_width: u32,
    pub max_pcie_link_speed: Option<u64>,
    pub is_multi_gpu_board: bool,
    pub vbios_version: String,
    pub memory_bus_width: u32,
    pub power_management_limit_default: u32,
    pub power_management_limit: u32,
    pub brand: String,
    pub board_id: Option<String>,
    pub serial: Option<String>,
}
impl From<&GpuDetails> for SerializableGpuDetails {
    fn from(gpu: &GpuDetails) -> Self {
        Self {
            name: gpu.name.clone(),
            uuid: gpu.uuid.clone(),
            graphics_clock: gpu.graphics_clock,
            memory_clock: gpu.memory_clock,
            sm_clock: gpu.sm_clock,
            perf_state: format!("{:?}", gpu.perf_state),
            architecture: gpu.architecture.clone(),
            cuda_cores: gpu.cuda_cores,
            max_pcie_link_gen: gpu.max_pcie_link_gen,
            max_pcie_link_width: gpu.max_pcie_link_width,
            max_pcie_link_speed: gpu.max_pcie_link_speed,
            is_multi_gpu_board: gpu.is_multi_gpu_board,
            vbios_version: gpu.vbios_version.clone(),
            memory_bus_width: gpu.memory_bus_width,
            power_management_limit_default: gpu.power_management_limit_default,
            power_management_limit: gpu.power_management_limit,
            brand: format!("{:?}", gpu.brand),
            board_id: gpu.board_id.clone(),
            serial: gpu.serial.clone(),
        }
    }
}

// Keep Display implementation for human-readable output
impl std::fmt::Display for GpuDetails {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "Device ID: {}\n\
            Name: {} \n\
            Architecture: {}\n\
            CUDA Cores: {}\n\
            PCIe: Gen {} x{} @ {}\n\
            Multi-GPU Board: {}\n\
            Brand: {}\n\
            VBIOS Version: {}\n\
            Memory Bus Width: {} bits\n\
            Total Memory: {} MB\n\
            Power Limits:\n\
              Default: {}W\n\
              Current: {}W\n\
              Range: {}W - {}W",
            self.uuid,
            self.name,
            self.architecture,
            self.cuda_cores,
            self.max_pcie_link_gen,
            self.max_pcie_link_width,
            self.max_pcie_link_speed
                .map(|x| format!("{} GT/s", x / 1_000_000))
                .unwrap_or_else(|| String::from("<unknown>")),
            if self.is_multi_gpu_board { "Yes" } else { "No" },
            format!("{:?}", self.brand),
            self.vbios_version,
            self.memory_bus_width,
            self.memory_info.total / 1024 / 1024,
            self.power_management_limit_default / 1000,
            self.power_management_limit / 1000,
            self.power_management_limit_constraints.min_limit / 1000,
            self.power_management_limit_constraints.max_limit / 1000
        )
    }
}

#[derive(Debug, Clone, Copy, Serialize, Deserialize)]
pub struct ClockSpeeds {
    pub graphics: u32,
    pub memory: u32,
    pub sm: u32,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TemperatureThresholds {
    pub slowdown: u32,
    pub shutdown: u32,
}

#[derive(Debug, Clone)]
pub struct DeviceState {
    pub utilization: Utilization,
    pub power: u32,
    pub temperature: u32,
    pub memory_info: MemoryInfo,
    pub clock_speeds: ClockSpeeds,
    pub performance_state: PerformanceState,
    pub timestamp: std::time::Instant,
    pub pcie_link_gen: u32,
    pub pcie_link_width: u32,
    pub pcie_link_speed: u64,
    pub power_usage: u32,
    pub fan_speed: Option<u32>,
    pub encoder_utilization: Option<u32>,
    pub decoder_utilization: Option<u32>,
    pub temperature_threshold: TemperatureThresholds,
    pub power_source: PowerSource,
    pub current_compute_mode: ComputeMode,
}

// Serializable version of DeviceState for JSON output
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SerializableDeviceState {
    pub gpu_utilization: u32,
    pub memory_utilization: u32,
    pub power: u32,
    pub temperature: u32,
    pub memory_total: u64,
    pub memory_used: u64,
    pub memory_free: u64,
    pub clock_speeds: ClockSpeeds,
    pub performance_state: String,
    pub timestamp: u64, // Unix timestamp in milliseconds
    pub pcie_link_gen: u32,
    pub pcie_link_width: u32,
    pub pcie_link_speed: u64,
    pub power_usage: u32,
    pub fan_speed: Option<u32>,
    pub encoder_utilization: Option<u32>,
    pub decoder_utilization: Option<u32>,
    pub temperature_threshold: TemperatureThresholds,
    pub power_source: String,
    pub current_compute_mode: String,
}

impl From<&DeviceState> for SerializableDeviceState {
    fn from(state: &DeviceState) -> Self {
        Self {
            gpu_utilization: state.utilization.gpu,
            memory_utilization: state.utilization.memory,
            power: state.power,
            temperature: state.temperature,
            memory_total: state.memory_info.total,
            memory_used: state.memory_info.used,
            memory_free: state.memory_info.free,
            clock_speeds: state.clock_speeds,
            performance_state: format!("{:?}", state.performance_state),
            timestamp: state.timestamp.elapsed().as_millis() as u64,
            pcie_link_gen: state.pcie_link_gen,
            pcie_link_width: state.pcie_link_width,
            pcie_link_speed: state.pcie_link_speed,
            power_usage: state.power_usage,
            fan_speed: state.fan_speed,
            encoder_utilization: state.encoder_utilization,
            decoder_utilization: state.decoder_utilization,
            temperature_threshold: state.temperature_threshold.clone(),
            power_source: format!("{:?}", state.power_source),
            current_compute_mode: format!("{:?}", state.current_compute_mode),
        }
    }
}

impl std::fmt::Display for DeviceState {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "GPU Utilization: {}%\n\
            Memory Utilization: {}%\n\
            Power: {}W\n\
            Temperature: {}°C\n\
            Memory: {:.1}GB used / {:.1}GB total\n\
            Clocks (MHz):\n\
              Graphics: {}\n\
              Memory: {}\n\
              SM: {}\n\
            PCIe: Gen {} x{} @ {} GT/s\n\
            Performance State: {:?}\n\
            Power Usage: {}W\n\
            Fan Speed: {}%\n\
            Encoder: {}%\n\
            Decoder: {}%\n\
            Temperature Thresholds:\n\
              Slowdown: {}°C\n\
              Shutdown: {}°C\n\
            Power Source: {:?}\n\
            Compute Mode: {:?}",
            self.utilization.gpu,
            self.utilization.memory,
            self.power,
            self.temperature,
            self.memory_info.used as f64 / 1024.0 / 1024.0 / 1024.0,
            self.memory_info.total as f64 / 1024.0 / 1024.0 / 1024.0,
            self.clock_speeds.graphics,
            self.clock_speeds.memory,
            self.clock_speeds.sm,
            self.pcie_link_gen,
            self.pcie_link_width,
            self.pcie_link_speed / 1_000_000,
            self.performance_state,
            self.power_usage,
            self.fan_speed.unwrap_or(0),
            self.encoder_utilization.unwrap_or(0),
            self.decoder_utilization.unwrap_or(0),
            self.temperature_threshold.slowdown,
            self.temperature_threshold.shutdown,
            self.power_source,
            self.current_compute_mode
        )
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SerializableAccountingStats {
    pub pid: u32,
    pub process_name: String,
    pub gpu_utilization: Option<u32>,
    pub is_running: bool,
    pub max_memory_usage: Option<u64>,
    pub memory_utilization: Option<u32>,
    pub start_time: u64,
    pub time: u64,
}

impl SerializableAccountingStats {
    pub fn new(pid: u32, process_name: String, stats: AccountingStats) -> Self {
        Self {
            pid,
            process_name,
            gpu_utilization: stats.gpu_utilization,
            is_running: stats.is_running,
            max_memory_usage: stats.max_memory_usage,
            memory_utilization: stats.memory_utilization,
            start_time: stats.start_time,
            time: stats.time,
        }
    }
}

// Helper functions for GPU monitoring
pub fn get_gpu_infos() -> anyhow::Result<Vec<GpuDetails>> {
    let nvml = NVML.as_ref().map_err(|e| anyhow!("NVML error: {}", e))?;
    let device_count = nvml.device_count()?;
    let mut gpu_info = Vec::new();

    for i in 0..device_count {
        let device = nvml.device_by_index(i)?;
        let details = GpuDetails {
            name: device.name()?,
            uuid: device.uuid()?,
            pci_info: device.pci_info()?,
            graphics_clock: device.clock_info(Clock::Graphics).unwrap_or_default(),
            memory_clock: device.clock_info(Clock::Memory).unwrap_or_default(),
            sm_clock: device.clock_info(Clock::SM).unwrap_or_default(),
            memory_info: device.memory_info()?,
            perf_state: device
                .performance_state()
                .unwrap_or(PerformanceState::Unknown),
            architecture: device
                .architecture()
                .unwrap_or(DeviceArchitecture::Unknown)
                .to_string(),
            cuda_cores: device.num_cores().unwrap_or_default(),
            max_pcie_link_gen: device.max_pcie_link_gen().unwrap_or_default(),
            max_pcie_link_width: device.max_pcie_link_width().unwrap_or_default(),
            max_pcie_link_speed: device
                .max_pcie_link_speed()
                .map(|x| x.as_integer().map(|y| y as u64))
                .unwrap_or(None),
            is_multi_gpu_board: device.is_multi_gpu_board().unwrap_or_default(),
            vbios_version: device.vbios_version().unwrap_or_default(),
            memory_bus_width: device.memory_bus_width().unwrap_or_default(),
            power_management_limit_default: device
                .power_management_limit_default()
                .unwrap_or_default(),
            power_management_limit: device.power_management_limit().unwrap_or_default(),
            power_management_limit_constraints: device
                .power_management_limit_constraints()
                .unwrap_or(PowerManagementConstraints {
                    min_limit: 0,
                    max_limit: 0,
                }),
            brand: device.brand().unwrap_or(Brand::Unknown),
            board_id: device.board_id().ok().map(|x| x.to_string()),
            serial: device.serial().ok().map(|x| x.to_string()),
        };
        gpu_info.push(details);
    }
    Ok(gpu_info)
}

pub fn get_gpu_info(device: &nvml_wrapper::Device) -> anyhow::Result<GpuDetails> {
    let details = GpuDetails {
        name: device.name()?,
        uuid: device.uuid()?,
        pci_info: device.pci_info()?,
        graphics_clock: device.clock_info(Clock::Graphics).unwrap_or_default(),
        memory_clock: device.clock_info(Clock::Memory).unwrap_or_default(),
        sm_clock: device.clock_info(Clock::SM).unwrap_or_default(),
        memory_info: device.memory_info()?,
        perf_state: device
            .performance_state()
            .unwrap_or(PerformanceState::Unknown),
        architecture: device
            .architecture()
            .unwrap_or(DeviceArchitecture::Unknown)
            .to_string(),
        cuda_cores: device.num_cores().unwrap_or_default(),
        max_pcie_link_gen: device.max_pcie_link_gen().unwrap_or_default(),
        max_pcie_link_width: device.max_pcie_link_width().unwrap_or_default(),
        max_pcie_link_speed: device
            .max_pcie_link_speed()
            .map(|x| x.as_integer().map(|y| y as u64))
            .unwrap_or(None),
        is_multi_gpu_board: device.is_multi_gpu_board().unwrap_or_default(),
        vbios_version: device.vbios_version().unwrap_or_default(),
        memory_bus_width: device.memory_bus_width().unwrap_or_default(),
        power_management_limit_default: device.power_management_limit_default().unwrap_or_default(),
        power_management_limit: device.power_management_limit().unwrap_or_default(),
        power_management_limit_constraints: device.power_management_limit_constraints().unwrap_or(
            PowerManagementConstraints {
                min_limit: 0,
                max_limit: 0,
            },
        ),
        brand: device.brand().unwrap_or(Brand::Unknown),
        board_id: device.board_id().ok().map(|x| x.to_string()),
        serial: device.serial().ok().map(|x| x.to_string()),
    };
    Ok(details)
}

pub fn collect_device_state(device: &nvml_wrapper::Device) -> anyhow::Result<DeviceState> {
    let temperature_threshold = TemperatureThresholds {
        slowdown: device
            .temperature_threshold(TemperatureThreshold::Slowdown)
            .unwrap_or_default(),
        shutdown: device
            .temperature_threshold(TemperatureThreshold::Shutdown)
            .unwrap_or_default(),
    };

    Ok(DeviceState {
        utilization: device.utilization_rates()?,
        power: device.power_usage().unwrap_or_default(),
        temperature: device
            .temperature(TemperatureSensor::Gpu)
            .unwrap_or_default(),
        memory_info: device.memory_info()?,
        clock_speeds: ClockSpeeds {
            graphics: device.clock_info(Clock::Graphics).unwrap_or_default(),
            memory: device.clock_info(Clock::Memory).unwrap_or_default(),
            sm: device.clock_info(Clock::SM).unwrap_or_default(),
        },
        performance_state: device
            .performance_state()
            .unwrap_or(PerformanceState::Unknown),
        timestamp: std::time::Instant::now(),
        pcie_link_gen: device.current_pcie_link_gen().unwrap_or_default(),
        pcie_link_width: device.current_pcie_link_width().unwrap_or_default(),
        pcie_link_speed: u64::from(device.pcie_link_speed().unwrap_or_default()),
        power_usage: device.power_usage().unwrap_or_default(),
        fan_speed: Some(device.fan_speed(0).unwrap_or(0) as u32),
        encoder_utilization: Some(
            device
                .encoder_utilization()
                .unwrap_or(UtilizationInfo {
                    utilization: 0,
                    sampling_period: 0,
                })
                .utilization,
        ),
        decoder_utilization: Some(
            device
                .decoder_utilization()
                .unwrap_or(UtilizationInfo {
                    utilization: 0,
                    sampling_period: 0,
                })
                .utilization,
        ),
        temperature_threshold,
        power_source: device.power_source().unwrap_or(PowerSource::Ac),
        current_compute_mode: device.compute_mode().unwrap_or(ComputeMode::Default),
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use tracing::{debug, error};

    #[test]
    fn test_get_gpu_info() -> Result<()> {
        // let _ = tracing_subscriber::fmt()
        //     .with_env_filter(
        //         tracing_subscriber::EnvFilter::from_default_env()
        //             .add_directive(tracing::Level::DEBUG.into()),
        //     )
        //     .try_init();
        // Check if we're running in a proper CUDA environment
        if let Err(e) = NVML.as_ref() {
            error!("Skipping test_get_gpu_info: NVML not available ({}).", e);
            error!("This is normal when running outside of Docker or without NVIDIA drivers.");
            return Err(anyhow::anyhow!("NVML not available: {}", e));
        }

        let gpu_infos = get_gpu_infos()?;
        assert!(!gpu_infos.is_empty(), "Should have at least one GPU");

        // Verify each GPU info
        for info in &gpu_infos {
            // Print formatted GPU details
            debug!("GPU Details: {:?}", info);
        }

        Ok(())
    }

    #[test]
    fn test_collect_device_state() -> Result<()> {
        // let _ = tracing_subscriber::fmt()
        //     .with_env_filter(
        //         tracing_subscriber::EnvFilter::from_default_env()
        //             .add_directive(tracing::Level::DEBUG.into()),
        //     )
        //     .try_init();
        // Check if we're running in a proper CUDA environment
        if let Err(e) = NVML.as_ref() {
            error!("Skipping test_get_gpu_info: NVML not available ({}).", e);
            error!("This is normal when running outside of Docker or without NVIDIA drivers.");
            return Err(anyhow::anyhow!("NVML not available: {}", e));
        }

        let device = NVML.as_ref().unwrap().device_by_index(0).unwrap();
        let state = collect_device_state(&device)?;

        debug!("Device State: {}", state);

        Ok(())
    }
}
