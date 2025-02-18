use anyhow::{anyhow, Result};
use chrono::Utc;
use cupti_sys::*;
use flume::Sender;
use std::collections::HashMap;
use std::process::Command;
use libc;
use nvml_wrapper::{
    enum_wrappers::device::Clock,
    struct_wrappers::device::AccountingStats,
};
use std::time::{Duration, Instant};
use std::{ffi::c_void, ptr};
use tracing::{debug, error, warn};
use nvml_wrapper::enums::device::UsedGpuMemory;
use crate::{get_gpu_info, GpuMetrics, LogEntry, GPU_MONITOR_TOPIC, NVML, SerializableGpuMetrics, SerializableGpuDetails};
use tokio;
use serde_json;
use tokio::sync::oneshot;

const CUPTI_START_INFO: &str = "cupti_start_info";
const CUPTI_METRICS: &str = "cupti_metrics";

impl std::fmt::Debug for CuptiTflopsCalculator {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("CuptiTflopsCalculator")
            .field("total_flops", &self.total_flops)
            .field("start_time", &self.start_time)
            .field("event_count", &self.event_ids.len())
            .finish()
    }
}

// SAFETY: The CUDA context is managed safely within this struct:
// 1. Context is only accessed in synchronized methods
// 2. All CUDA operations are properly synchronized
// 3. The context is managed by CUDA driver's own thread safety mechanisms
unsafe impl Send for CuptiTflopsCalculator {}
unsafe impl Sync for CuptiTflopsCalculator {}

#[derive(Clone, Copy)]
struct CudaContext(CUcontext);
unsafe impl Send for CudaContext {}
unsafe impl Sync for CudaContext {}

impl CudaContext {
    unsafe fn push(&self) -> Result<(), CuptiError> {
        let result = cuCtxPushCurrent_v2(self.0);
        if result != cudaError_enum::CUDA_SUCCESS {
            return Err(CuptiError::MeasurementError(CUptiResult::CUPTI_ERROR_NOT_READY));
        }
        Ok(())
    }
}

#[derive(Clone, Copy)]
struct CudaEventGroup(CUpti_EventGroup);
unsafe impl Send for CudaEventGroup {}
unsafe impl Sync for CudaEventGroup {}

pub struct CuptiTflopsCalculator {
    context: CUcontext,
    event_groups: Vec<CudaEventGroup>,
    event_ids: Vec<CUpti_EventID>,
    total_flops: u64,
    start_time: Option<Instant>,
    log_tx: Sender<LogEntry>,
    resource_provider_id: String,
    deal_id: String,
    process_id: u32,
}

#[derive(Debug)]
pub enum CuptiError {
    InitError(CUptiResult),
    MeasurementError(CUptiResult),
    NoContext,
    InsufficientPrivileges(&'static str),
}

impl std::error::Error for CuptiError {}

impl std::fmt::Display for CuptiError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}

impl CuptiTflopsCalculator {
    fn check_cupti_privileges() -> Result<(), CuptiError> {
        // CUPTI requires root privileges
        unsafe {
            if libc::geteuid() != 0 {
                return Err(CuptiError::InsufficientPrivileges(
                    "CUPTI requires root privileges. Please run with sudo.",
                ));
            }
        }
        Ok(())
    }
    fn set_accounting_mode() -> Result<String, anyhow::Error> {
        let output = Command::new("nvidia-smi")
            .arg("-pm")
            .arg("1")
            .output()
            .map_err(|e| anyhow!("Failed to set process mode: {}", e))?;

        if !output.status.success() {
            let error = String::from_utf8_lossy(&output.stderr);
            return Err(anyhow!("nvidia-smi failed: {}", error));
        }
        let output = Command::new("nvidia-smi")
            .arg("-am")
            .arg("1")
            .output()
            .map_err(|e| anyhow!("Failed to set accounting mode: {}", e))?;

        if !output.status.success() {
            let error = String::from_utf8_lossy(&output.stderr);
            return Err(anyhow!("nvidia-smi failed: {}", error));
        }

        let info = String::from_utf8(output.stdout)
            .map_err(|e| anyhow!("Invalid UTF-8 in nvidia-smi output: {}", e))?;

        Ok(info.trim().to_string())
    }

    fn get_accounting_mode() -> Result<String> {
        let output = Command::new("nvidia-smi")
            .arg("--query-gpu=accounting.mode")
            .arg("--format=csv,noheader")
            .output()
            .map_err(|e| anyhow!("Failed to get accounting mode: {}", e))?;

        if !output.status.success() {
            let error = String::from_utf8_lossy(&output.stderr);
            return Err(anyhow!("nvidia-smi failed: {}", error));
        }

        let info = String::from_utf8(output.stdout)
            .map_err(|e| anyhow!("Invalid UTF-8 in nvidia-smi output: {}", e))?;

        Ok(info.trim().to_string())
    }

    fn find_gpu_process(&self) -> u32 {
        // Get our real process path
        let our_path = std::fs::read_link("/proc/self/exe")
            .unwrap_or_else(|_| std::path::PathBuf::new());
        let our_name = our_path.file_name()
            .and_then(|n| n.to_str())
            .unwrap_or("unknown");
        
        debug!("Our process: {}", our_name);
        
        match NVML.as_ref() {
            Ok(nvml) => {
                if let Ok(device) = nvml.device_by_index(0) {
                    debug!("=== GPU Process List ===");
                    
                    // Check accounting stats
                    if let Ok(pids) = device.accounting_pids() {
                        debug!("Processes with accounting enabled: {}", pids.len());
                        for &pid in &pids {
                            let proc_path = std::fs::read_link(format!("/proc/{}/exe", pid))
                                .unwrap_or_else(|_| std::path::PathBuf::new());
                            let name = proc_path.file_name()
                                .and_then(|n| n.to_str())
                                .map(|s| s.to_string())
                                .unwrap_or_else(|| {
                                    std::fs::read_to_string(format!("/proc/{}/comm", pid))
                                        .unwrap_or_else(|_| String::from("unknown"))
                                        .trim()
                                        .to_string()
                                });
                            
                            if let Ok(stats) = device.accounting_stats_for(pid) {
                                debug!("PID {}: {} ({}) - GPU util: {}%, Memory: {}MB", 
                                    pid,
                                    name,
                                    proc_path.display(),
                                    stats.gpu_utilization.unwrap_or(0),
                                    stats.max_memory_usage.unwrap_or(0) / 1024 / 1024);
                                
                                // Check if this is our process
                                if proc_path == our_path {
                                    debug!("Found our process!");
                                    return pid;
                                }
                            }
                        }
                    }

                    // List compute processes
                    let compute_procs = device.running_compute_processes().unwrap_or_default();
                    for proc in &compute_procs {
                        let proc_path = std::fs::read_link(format!("/proc/{}/exe", proc.pid))
                            .unwrap_or_else(|_| std::path::PathBuf::new());
                        let name = proc_path.file_name()
                            .and_then(|n| n.to_str())
                            .map(|s| s.to_string())
                            .unwrap_or_else(|| {
                                std::fs::read_to_string(format!("/proc/{}/comm", proc.pid))
                                    .unwrap_or_else(|_| String::from("unknown"))
                                    .trim()
                                    .to_string()
                            });
                        let memory_mb = match proc.used_gpu_memory {
                            UsedGpuMemory::Used(bytes) => bytes / 1024 / 1024,
                            UsedGpuMemory::Unavailable => 0,
                        };
                        debug!("Compute PID {}: {} ({}) - Memory: {}MB", 
                            proc.pid, name, proc_path.display(), memory_mb);
                        
                        if proc_path == our_path {
                            debug!("Found our process!");
                            return proc.pid;
                        }
                    }

                    // List graphics processes
                    let graphics_procs = device.running_graphics_processes().unwrap_or_default();
                    for proc in &graphics_procs {
                        let proc_path = std::fs::read_link(format!("/proc/{}/exe", proc.pid))
                            .unwrap_or_else(|_| std::path::PathBuf::new());
                        let name = proc_path.file_name()
                            .and_then(|n| n.to_str())
                            .map(|s| s.to_string())
                            .unwrap_or_else(|| {
                                std::fs::read_to_string(format!("/proc/{}/comm", proc.pid))
                                    .unwrap_or_else(|_| String::from("unknown"))
                                    .trim()
                                    .to_string()
                            });
                        let memory_mb = match proc.used_gpu_memory {
                            UsedGpuMemory::Used(bytes) => bytes / 1024 / 1024,
                            UsedGpuMemory::Unavailable => 0,
                        };
                        debug!("Graphics PID {}: {} ({}) - Memory: {}MB", 
                            proc.pid, name, proc_path.display(), memory_mb);
                    }
                    debug!("=====================");
                }
            }
            Err(e) => debug!("NVML error: {}", e),
        }
        
        let current_pid = std::process::id();
        warn!("Falling back to current process ID: {}", current_pid);
        current_pid
    }


    pub fn new(log_tx: Sender<LogEntry>, resource_provider_id: String, deal_id: String) -> Result<Self, CuptiError> {
        // Check privileges first
        // Self::check_cupti_privileges()?;

        // Check if accounting mode is enabled, if not, enable it
        match Self::get_accounting_mode() {
            Ok(mode) if mode != "Enabled" => {
                if let Err(e) = Self::set_accounting_mode() {
                    warn!(
                        "Failed to enable GPU accounting mode: {}. Some metrics may not be available.",
                        e
                    );
                }
            }
            Err(e) => {
                warn!(
                    "Failed to check GPU accounting mode: {}. Some metrics may not be available.",
                    e
                );
            }
            _ => {}
        }

        unsafe {
            // Initialize CUDA first
            let result = cuInit(0);
            if result != cudaError_enum::CUDA_SUCCESS {
                error!("Failed to initialize CUDA: {:?}", result);
                return Err(CuptiError::InitError(
                    CUptiResult::CUPTI_ERROR_NOT_INITIALIZED,
                ));
            }
            debug!("CUDA initialized successfully");

            // First create/get CUDA context
            let mut context = ptr::null_mut();
            let result = cuCtxGetCurrent(&mut context);
            if result != cudaError_enum::CUDA_SUCCESS || context.is_null() {
                debug!("No current CUDA context, creating one");
                let mut device = 0;
                let result = cuDeviceGet(&mut device, 0);
                if result != cudaError_enum::CUDA_SUCCESS {
                    error!("Failed to get CUDA device: {:?}", result);
                    return Err(CuptiError::NoContext);
                }

                let result = cuCtxCreate_v2(&mut context, 0, device);
                if result != cudaError_enum::CUDA_SUCCESS {
                    error!("Failed to create CUDA context: {:?}", result);
                    return Err(CuptiError::NoContext);
                }
                debug!("Created new CUDA context: {:?}", context);
            } else {
                debug!("Using existing CUDA context: {:?}", context);
            }

            // Initialize CUPTI by getting a timestamp
            let mut timestamp: u64 = 0;
            let result = cuptiGetTimestamp(&mut timestamp);
            if result != CUptiResult::CUPTI_SUCCESS {
                error!("Failed to initialize CUPTI: {:?}", result);
                return Err(CuptiError::InitError(result));
            }
            debug!("CUPTI initialized successfully (timestamp: {})", timestamp);

            let mut event_groups = Vec::new();
            let mut event_ids = Vec::new();

            // Modify the event keywords to include only the metrics we track
            let event_keywords = [
                "fp16",         // Used for FP16 operations
                "fp32",         // Used for FP32 operations
                "tensor",       // Used for tensor operations
                "flop",         // Used for direct FLOP counts
                "active_cycles", // Used for activity metrics
                "active_warps", // Used for warp activity
                "inst_executed", // Used for instruction counts
                "sm_efficiency", // Used for SM efficiency
                // Remove these as they aren't used in our metrics:
                // "fb_subp",           // Memory subsystem events
                // "l2_subp",           // L2 cache events     
                // "elapsed_cycles_sm",  // Basic SM cycle counting
                // "gpu_busy",          // GPU utilization
                // "single_precision",   // Redundant with fp32
                // "double_precision",   // Not tracking DP ops
                // "tensor_pipe",        // Redundant with tensor
            ];
            let mut target_events = Vec::new();

            // Get available domains
            let mut num_domains = 0u32;
            let result = cuptiDeviceGetNumEventDomains(0, ptr::addr_of_mut!(num_domains) as *mut _);
            if result != CUptiResult::CUPTI_SUCCESS {
                error!("Failed to get number of domains: {:?}", result);
                return Err(CuptiError::InitError(result));
            }
            debug!("Found {} event domains", num_domains);

            let mut domains = vec![0u32; num_domains as usize];
            let result = cuptiDeviceEnumEventDomains(
                0,
                ptr::addr_of_mut!(num_domains) as *mut _,
                domains.as_mut_ptr(),
            );
            if result != CUptiResult::CUPTI_SUCCESS {
                error!("Failed to enumerate domains: {:?}", result);
                return Err(CuptiError::InitError(result));
            }

            debug!("Available events by domain:");
            // For each domain, enumerate events and look for our target events
            for &domain_id in &domains {
                // Skip invalid domain ID
                if domain_id == 0 {
                    continue;
                }

                // Get domain name for debugging
                let mut name = [0i8; 256];
                let mut name_size = name.len() as u32;
                let result = cuptiEventDomainGetAttribute(
                    domain_id,
                    CUpti_EventDomainAttribute::CUPTI_EVENT_DOMAIN_ATTR_NAME,
                    ptr::addr_of_mut!(name_size) as *mut _,
                    name.as_mut_ptr() as *mut c_void,
                );

                let domain_name = if result == CUptiResult::CUPTI_SUCCESS {
                    std::ffi::CStr::from_ptr(name.as_ptr())
                        .to_str()
                        .unwrap_or("invalid_name")
                } else {
                    "unknown_domain"
                };

                let mut num_events = 0u32;
                let result = cuptiEventDomainGetNumEvents(
                    domain_id,
                    ptr::addr_of_mut!(num_events) as *mut _,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    warn!(
                        "Failed to get number of events for domain {}: {:?}",
                        domain_id, result
                    );
                    continue;
                }
                debug!(
                    "Domain {} ({}) has {} events",
                    domain_id, domain_name, num_events
                );

                let mut events = vec![0u32; num_events as usize];
                let result = cuptiEventDomainEnumEvents(
                    domain_id,
                    ptr::addr_of_mut!(num_events) as *mut _,
                    events.as_mut_ptr(),
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    warn!(
                        "Failed to enumerate events for domain {}: {:?}",
                        domain_id, result
                    );
                    continue;
                }

                debug!("Events in domain {}:", domain_name);
                // List all events in this domain
                for &event_id in &events {
                    // Skip invalid event ID
                    if event_id == 0 {
                        continue;
                    }

                    let mut name = [0i8; 256];
                    let mut name_size = name.len() as u32;
                    let mut desc = [0i8; 1024];
                    let mut desc_size = desc.len() as u32;

                    let name_result = cuptiEventGetAttribute(
                        event_id,
                        CUpti_EventAttribute::CUPTI_EVENT_ATTR_NAME,
                        ptr::addr_of_mut!(name_size) as *mut _,
                        name.as_mut_ptr() as *mut c_void,
                    );

                    let desc_result = cuptiEventGetAttribute(
                        event_id,
                        CUpti_EventAttribute::CUPTI_EVENT_ATTR_LONG_DESCRIPTION,
                        ptr::addr_of_mut!(desc_size) as *mut _,
                        desc.as_mut_ptr() as *mut c_void,
                    );

                    let name_str = if name_result == CUptiResult::CUPTI_SUCCESS {
                        std::ffi::CStr::from_ptr(name.as_ptr())
                            .to_str()
                            .unwrap_or("invalid_name")
                    } else {
                        "unknown_event"
                    };

                    let desc_str = if desc_result == CUptiResult::CUPTI_SUCCESS {
                        std::ffi::CStr::from_ptr(desc.as_ptr())
                            .to_str()
                            .unwrap_or("no description")
                    } else {
                        "no description"
                    };

                    debug!("  Event {}: {} - {}", event_id, name_str, desc_str);

                    // Check if this event matches any of our keywords
                    if event_keywords
                        .iter()
                        .any(|&keyword| name_str.contains(keyword))
                    {
                        debug!("Found relevant event: {} ({})", name_str, event_id);
                        target_events.push(name_str.to_string());
                    }
                }
            }

            debug!("Selected target events: {:?}", target_events);

            // Create an event group
            let mut event_group = ptr::null_mut();
            let result = cuptiEventGroupCreate(context, &mut event_group, 0);
            if result != CUptiResult::CUPTI_SUCCESS {
                error!("Failed to create event group: {:?}", result);
                return Err(CuptiError::InitError(result));
            }
            debug!("Created event group: {:?}", event_group);

            // Try to add each event that matches our target names
            for &domain_id in &domains {
                if domain_id == 0 {
                    continue;
                }

                let mut num_events = 0u32;
                let result = cuptiEventDomainGetNumEvents(
                    domain_id,
                    ptr::addr_of_mut!(num_events) as *mut _,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    continue;
                }

                let mut events = vec![0u32; num_events as usize];
                let result = cuptiEventDomainEnumEvents(
                    domain_id,
                    ptr::addr_of_mut!(num_events) as *mut _,
                    events.as_mut_ptr(),
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    continue;
                }

                for &event_id in &events {
                    if event_id == 0 {
                        continue;
                    }

                    let mut name = [0i8; 256];
                    let mut name_size = name.len() as u32;

                    let result = cuptiEventGetAttribute(
                        event_id,
                        CUpti_EventAttribute::CUPTI_EVENT_ATTR_NAME,
                        ptr::addr_of_mut!(name_size) as *mut _,
                        name.as_mut_ptr() as *mut c_void,
                    );

                    if result != CUptiResult::CUPTI_SUCCESS {
                        continue;
                    }

                    let name_str = std::ffi::CStr::from_ptr(name.as_ptr())
                        .to_str()
                        .unwrap_or("invalid_name");

                    // Check if this is one of our target events
                    if target_events.iter().any(|target| name_str == target) {
                        debug!(
                            "Adding target event: {} ({}) from domain {}",
                            name_str, event_id, domain_id
                        );

                        // Add event to group
                        let result = cuptiEventGroupAddEvent(event_group, event_id);
                        if result != CUptiResult::CUPTI_SUCCESS {
                            warn!(
                                "Failed to add event {} ({}) to group: {:?}",
                                event_id, name_str, result
                            );
                            continue;
                        }

                        // Store event ID for later use
                        event_ids.push(event_id);
                        debug!("Added event {} ({}) to group", event_id, name_str);
                    }
                }
            }

            // Only add the event group if we found any events
            if !event_ids.is_empty() {
                event_groups.push(CudaEventGroup(event_group));
                debug!("Successfully set up {} event groups", event_groups.len());
            } else {
                warn!("No target events were found and added to groups");
                let _ = cuptiEventGroupDestroy(event_group);
            }

            Ok(Self {
                context,
                event_groups,
                event_ids,
                start_time: None,
                total_flops: 0,
                log_tx,
                resource_provider_id,
                deal_id,
                process_id: 0,
            })
        }
    }

    pub async fn start_measurement(&mut self) -> Result<(), CuptiError> {
        self.process_id = self.find_gpu_process();
        debug!("Starting CUPTI measurement");
        
        let gpu_info = &get_gpu_info().unwrap()[0];
        let serializable_gpu_info = SerializableGpuDetails::from(gpu_info);
        let details = match serde_json::to_string(&serializable_gpu_info) {
            Ok(json) => json,
            Err(e) => {
                warn!("Failed to serialize GPU info: {}", e);
                format!("{}", gpu_info) // Fallback to Display impl
            }
        };
        if let Err(e) = self.log_tx.send_async(LogEntry {
            topic: GPU_MONITOR_TOPIC.to_string(),
            resource_provider_id: self.resource_provider_id.clone(),
            deal_id: self.deal_id.clone(),
            metadata: HashMap::new(),
            timestamp: Utc::now().to_rfc3339(),
            action: CUPTI_START_INFO.to_string(),
            source: gpu_info.uuid.to_string(),
            details,
        }).await {
            warn!("Failed to send log entry: {}", e);
        }

        unsafe {
            // Push the context before starting measurement
            let result = cuCtxPushCurrent_v2(self.context);
            if result != cudaError_enum::CUDA_SUCCESS {
                error!("Failed to push CUDA context: {:?}", result);
                return Err(CuptiError::MeasurementError(
                    CUptiResult::CUPTI_ERROR_NOT_READY,
                ));
            }

            // Configure and enable each group
            for (i, group) in self.event_groups.iter().enumerate() {
                // Enable profiling for all domain instances
                let profile_all = 1u32;
                let size = std::mem::size_of::<u32>();
                let result = cuptiEventGroupSetAttribute(
                    group.0,
                    CUpti_EventGroupAttribute::CUPTI_EVENT_GROUP_ATTR_PROFILE_ALL_DOMAIN_INSTANCES,
                    size,
                    &profile_all as *const _ as *mut c_void,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    error!("Failed to set profile all attribute: {:?}", result);
                    return Err(CuptiError::MeasurementError(result));
                }

                // Get number of instances and events for debugging
                let mut num_instances = 0u32;
                let mut size = std::mem::size_of::<u32>();
                let _ = cuptiEventGroupGetAttribute(
                    group.0,
                    CUpti_EventGroupAttribute::CUPTI_EVENT_GROUP_ATTR_INSTANCE_COUNT,
                    ptr::addr_of_mut!(size),
                    &mut num_instances as *mut _ as *mut c_void,
                );
                debug!("Group {} has {} instances", i, num_instances);

                let mut num_events = 0u32;
                let mut size = std::mem::size_of::<u32>();
                let result = cuptiEventGroupGetAttribute(
                    group.0,
                    CUpti_EventGroupAttribute::CUPTI_EVENT_GROUP_ATTR_NUM_EVENTS,
                    ptr::addr_of_mut!(size),
                    ptr::addr_of_mut!(num_events) as *mut c_void,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    error!("Failed to get number of events: {:?}", result);
                    continue;
                }
                debug!("Group {} has {} events", i, num_events);

                // Enable collection
                let result = cuptiEventGroupEnable(group.0);
                if result != CUptiResult::CUPTI_SUCCESS {
                    error!("Failed to enable event group {}: {:?}", i, result);
                    return Err(CuptiError::MeasurementError(result));
                }
                debug!("Enabled event group {}", i);
            }
        }
        self.start_time = Some(Instant::now());
        Ok(())
    }

    async fn run_on_cuda_thread<F, T>(&self, f: F) -> Result<T, CuptiError>
    where
        F: FnOnce() -> Result<T, CuptiError> + Send + 'static,
        T: Send + 'static,
    {
        let (tx, rx) = oneshot::channel();
        
        // Spawn the blocking operation on a dedicated thread
        tokio::task::spawn_blocking(move || {
            let result = f();
            let _ = tx.send(result);
        })
        .await
        .unwrap();

        rx.await.unwrap()
    }

    pub async fn stop_measurement(&mut self) -> Result<GpuMetrics, CuptiError> {
        debug!("Stopping CUPTI measurement");
        let elapsed = self.start_time.take().map(|start| start.elapsed())
            .unwrap_or(Duration::from_secs(1));

        let context = CudaContext(self.context);
        let event_groups = self.event_groups.clone();
        let event_ids = self.event_ids.clone();
        let process_id = self.process_id;
        
        let metrics = self.run_on_cuda_thread(move || unsafe {
            context.push()?;
            // Synchronize context
            let cuda_result = cuCtxSynchronize();
            if cuda_result != cudaError_enum::CUDA_SUCCESS {
                error!("Failed to synchronize CUDA context: {:?}", cuda_result);
                // Continue anyway to try to collect what data we can
            }

            let mut total_flops = 0u64;
            let mut active_cycles = 0u64;
            let mut active_warps = 0u64;
            let mut fp16_ops = 0u64;
            let mut fp32_ops = 0u64;
            let mut tensor_cycles = 0u64;
            let mut sm_efficiency = 0u64;
            let mut inst_executed = 0u64;

            // Read events while context is active
            for (group_idx, group) in event_groups.iter().enumerate() {
                debug!("Reading events from group {}", group_idx);

                // Get current context for debugging
                let mut current = ptr::null_mut();
                let _ = cuCtxGetCurrent(&mut current);
                debug!("Reading events with context: {:?}", current);

                // Get the number of events and instances
                let mut num_events = 0u32;
                let mut size: usize = std::mem::size_of::<u32>();

                let result = cuptiEventGroupGetAttribute(
                    group.0,
                    CUpti_EventGroupAttribute::CUPTI_EVENT_GROUP_ATTR_NUM_EVENTS,
                    ptr::addr_of_mut!(size),
                    ptr::addr_of_mut!(num_events) as *mut c_void,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    error!("Failed to get number of events: {:?}", result);
                    continue;
                }

                let mut num_instances = 0u32;
                let mut size: usize = std::mem::size_of::<u32>();

                let result = cuptiEventGroupGetAttribute(
                    group.0,
                    CUpti_EventGroupAttribute::CUPTI_EVENT_GROUP_ATTR_INSTANCE_COUNT,
                    ptr::addr_of_mut!(size),
                    ptr::addr_of_mut!(num_instances) as *mut c_void,
                );
                if result != CUptiResult::CUPTI_SUCCESS {
                    error!("Failed to get number of instances: {:?}", result);
                    continue;
                }

                debug!(
                    "Group has {} events and {} instances",
                    num_events, num_instances
                );

                if num_events == 0 {
                    warn!("No events in group {}, skipping", group_idx);
                    continue;
                }

                // Read each event with FORCE_INT flag
                for &event_id in &event_ids {
                    let mut event_values = vec![0u64; num_instances as usize];
                    let mut size = std::mem::size_of::<u64>() * num_instances as usize;

                    let result = cuptiEventGroupReadEvent(
                        group.0,
                        CUpti_ReadEventFlags::CUPTI_EVENT_READ_FLAG_FORCE_INT,
                        event_id,
                        ptr::addr_of_mut!(size),
                        event_values.as_mut_ptr() as *mut u64,
                    );

                    if result == CUptiResult::CUPTI_SUCCESS {
                        // Get event name for better logging
                        let mut name = [0i8; 256];
                        let mut name_size = name.len() as u32;
                        let name_result = cuptiEventGetAttribute(
                            event_id,
                            CUpti_EventAttribute::CUPTI_EVENT_ATTR_NAME,
                            ptr::addr_of_mut!(name_size) as *mut _,
                            name.as_mut_ptr() as *mut c_void,
                        );

                        let name_str = if name_result == CUptiResult::CUPTI_SUCCESS {
                            std::ffi::CStr::from_ptr(name.as_ptr())
                                .to_str()
                                .unwrap_or("invalid_name")
                        } else {
                            "unknown_event"
                        };

                        // Log each non-zero instance value
                        for (idx, &value) in event_values.iter().enumerate() {
                            if value > 0 {
                                debug!(
                                    "Event {} ({}) instance {}: {}",
                                    event_id, name_str, idx, value
                                );
                            }
                        }

                        let total_value: u64 = event_values.iter().sum();
                        debug!(
                            "Event {} ({}) total across instances: {}",
                            event_id, name_str, total_value
                        );

                        match name_str {
                            // FP16 operations
                            s if s.contains("fp16") || s.contains("half") => {
                                fp16_ops += total_value;
                                total_flops += total_value;
                                debug!("FP16 operations: {}", total_value);
                            }

                            // FP32 operations
                            s if s.contains("fp32") || s.contains("single_precision") => {
                                fp32_ops += total_value;
                                total_flops += total_value;
                                debug!("FP32 operations: {}", total_value);
                            }

                            // Tensor operations
                            s if s.contains("tensor") || s.contains("tensor_pipe") => {
                                tensor_cycles = total_value;
                                debug!("Tensor core cycles: {}", total_value);
                            }

                            // Direct FLOP counts
                            s if s.contains("flop") => {
                                total_flops += total_value;
                                debug!("FLOP count: {}", total_value);
                            }

                            // Activity metrics
                            s if s.contains("active_cycles") => {
                                active_cycles = total_value;
                                debug!("Active cycles: {}", total_value);
                            }
                            s if s.contains("active_warps") => {
                                active_warps = total_value;
                                debug!("Active warps: {}", total_value);
                            }

                            // Instruction and efficiency metrics
                            s if s.contains("inst_executed") => {
                                inst_executed = total_value;
                                debug!("Instructions executed: {}", total_value);
                            }
                            s if s.contains("sm_efficiency") => {
                                sm_efficiency = total_value;
                                debug!("SM efficiency: {}", total_value);
                            }

                            _ => {
                                debug!("Unhandled event: {} = {}", name_str, total_value);
                            }
                        }
                    } else {
                        error!("Failed to read event {}: {:?}", event_id, result);
                    }
                }
            }

            // Enhanced debug output with all metrics
            debug!("Collected metrics:");
            if active_cycles > 0 {
                debug!("  Active cycles: {}", active_cycles);
            }
            if active_warps > 0 {
                debug!("  Active warps: {}", active_warps);
            }
            if fp16_ops > 0 {
                debug!("  FP16 operations: {}", fp16_ops);
            }
            if fp32_ops > 0 {
                debug!("  FP32 operations: {}", fp32_ops);
            }
            if tensor_cycles > 0 {
                debug!("  Tensor core cycles: {}", tensor_cycles);
            }
            if sm_efficiency > 0 {
                debug!("  SM efficiency: {}", sm_efficiency);
            }
            if inst_executed > 0 {
                debug!("  Instructions executed: {}", inst_executed);
            }

            // If we have direct FLOP counts, show the breakdown
            if total_flops > 0 {
                debug!("Direct FLOP counts breakdown:");
                let mut breakdown = Vec::new();
                if fp16_ops > 0 {
                    let percent = (fp16_ops as f64 / total_flops as f64 * 100.0).round();
                    breakdown.push(format!("FP16: {} ({:.0}%)", fp16_ops, percent));
                }
                if fp32_ops > 0 {
                    let percent = (fp32_ops as f64 / total_flops as f64 * 100.0).round();
                    breakdown.push(format!("FP32: {} ({:.0}%)", fp32_ops, percent));
                }
                if tensor_cycles > 0 {
                    // Tensor cores can do significantly more ops per cycle
                    // For Ampere: 256 FP16 FMA ops per cycle per tensor core
                    let tensor_ops = tensor_cycles * 256; // This varies by architecture
                    let percent = (tensor_ops as f64 / total_flops as f64 * 100.0).round();
                    breakdown.push(format!("Tensor: {} ({:.0}%)", tensor_ops, percent));
                }
                debug!("  {}", breakdown.join(", "));
                debug!("  Total FLOPS: {}", total_flops);
            }
            // If we have no direct FLOP counts but have activity metrics, estimate FLOPS
            if total_flops == 0 && active_cycles > 0 && active_warps > 0 {
                // Get GPU capabilities from NVML
                let max_threads_per_warp = 32; // Default for most GPUs
                let mut max_ops_per_cycle = 2; // Conservative default
                let mut clock_mhz = 1500.0; // Default clock speed
                let mut sm_count = 1; // Default SM count

                let nvml = NVML
                    .as_ref()
                    .map_err(|e| anyhow::anyhow!("NVML error: {}", e))
                    .unwrap();
                if let Ok(device) = nvml.device_by_index(0) {
                    // Get compute capability
                    if let Ok(attrs) = device.cuda_compute_capability() {
                        debug!("GPU Compute Capability: {}.{}", attrs.major, attrs.minor);
                        // Ampere (SM 8.x) and later can do 2x FP32 ops per cycle
                        max_ops_per_cycle = if attrs.major >= 8 { 4 } else { 2 };
                    }

                    // Get actual clock speed
                    if let Ok(clock) = device.clock_info(Clock::SM) {
                        clock_mhz = clock as f64;
                        debug!("GPU SM Clock: {} MHz", clock_mhz);
                    }

                    // Get SM count
                    if let Ok(count) = device.num_cores() {
                        sm_count = count as i32;
                    }

                    // Get memory info
                    if let Ok(info) = device.memory_info() {
                        debug!(
                            "GPU Memory: Total {}MB, Used {}MB",
                            info.total / 1024 / 1024,
                            info.used / 1024 / 1024
                        );
                    }
                }

                // Calculate average warps per cycle
                let avg_warps_per_cycle = (active_warps as f64) / (active_cycles as f64);
                // Calculate total threads active per cycle
                let avg_threads_per_cycle = avg_warps_per_cycle * (max_threads_per_warp as f64);
                // Calculate total operations, factoring in SM count and clock speed
                let total_ops = avg_threads_per_cycle * (max_ops_per_cycle as f64);
                let clock_factor = clock_mhz / 1500.0; // Normalize to baseline clock
                let estimated_flops = ((total_ops * active_cycles as f64 * sm_count as f64)
                    * clock_factor)
                    .round() as u64;

                debug!("No direct FLOP counters available. Estimating from activity metrics:");
                debug!("  Active cycles: {}", active_cycles);
                debug!("  Active warps: {}", active_warps);
                debug!("  Average warps per cycle: {:.2}", avg_warps_per_cycle);
                debug!("  Average threads per cycle: {:.2}", avg_threads_per_cycle);
                debug!("  Operations per cycle per thread: {}", max_ops_per_cycle);
                debug!("  SM Count: {}", sm_count);
                debug!(
                    "  Clock Speed: {:.1} MHz (factor: {:.2})",
                    clock_mhz, clock_factor
                );
                debug!("  Formula: FLOPS = avg_threads_per_cycle ({:.2}) * ops_per_cycle ({}) * active_cycles ({}) * sm_count ({}) * clock_factor ({:.2})", 
                      avg_threads_per_cycle, max_ops_per_cycle, active_cycles, sm_count, clock_factor);
                debug!("  Estimated FLOPS: {}", estimated_flops);

                total_flops = estimated_flops;
            }

            // Calculate FLOPS before potentially popping context
            debug!("Total FLOPS: {}, elapsed: {:?}", total_flops, elapsed);

            // Make sure to pop the context if we pushed it
            if cuda_result == cudaError_enum::CUDA_SUCCESS {
                let mut popped = ptr::null_mut();
                let _ = cuCtxPopCurrent_v2(&mut popped);
            }

            let mut metrics = GpuMetrics {
                measured_flops: if fp16_ops > 0 || fp32_ops > 0 {
                    total_flops
                } else {
                    0
                },
                fp16_ops,
                fp32_ops,
                tensor_ops: tensor_cycles * 256,
                active_cycles,
                active_warps,
                sm_efficiency: sm_efficiency as f32,
                estimated_flops: 0,
                measured_tflops: 0.0,
                estimated_tflops: 0.0,
                sm_count: 1,
                clock_mhz: 1500.0,
                compute_capability: (0, 0),
                memory_bandwidth_gbps: 0.0,
                memory_utilization: 0.0,
                tensor_core_utilization: 0.0,
                accounting_stats: None,
            };

            // Get hardware info from NVML
            if let Ok(nvml) = NVML.as_ref() {
                if let Ok(device) = nvml.device_by_index(0) {
                    if let Ok(attrs) = device.cuda_compute_capability() {
                        metrics.compute_capability = (attrs.major as u32, attrs.minor as u32);
                    }
                    if let Ok(clock) = device.clock_info(Clock::SM) {
                        metrics.clock_mhz = clock as f64;
                    }
                    if let Ok(count) = device.num_cores() {
                        metrics.sm_count = count as u32;
                    }
                }
            }

            // Calculate measured TFLOPS if we have direct measurements
            if metrics.measured_flops > 0 {
                metrics.measured_tflops =
                    (metrics.measured_flops as f64) / (elapsed.as_secs_f64() * 1e12);
            }

            // Always calculate estimated TFLOPS if we have activity metrics
            if metrics.active_cycles > 0 && metrics.active_warps > 0 {
                let max_threads_per_warp = 32;
                let max_ops_per_cycle = if metrics.compute_capability.0 >= 8 {
                    4
                } else {
                    2
                };

                let avg_warps_per_cycle =
                    (metrics.active_warps as f64) / (metrics.active_cycles as f64);
                let avg_threads_per_cycle = avg_warps_per_cycle * (max_threads_per_warp as f64);
                let total_ops = avg_threads_per_cycle * (max_ops_per_cycle as f64);
                let clock_factor = metrics.clock_mhz / 1500.0;

                metrics.estimated_flops =
                    ((total_ops * metrics.active_cycles as f64 * metrics.sm_count as f64)
                        * clock_factor)
                        .round() as u64;
                metrics.estimated_tflops =
                    (metrics.estimated_flops as f64) / (elapsed.as_secs_f64() * 1e12);
            }

            debug!("Final metrics: {:?}", metrics);
            if metrics.measured_flops > 0 {
                debug!("  Measured TFLOPS: {:.2}", metrics.measured_tflops);
                debug!("  FP16 ops: {}", metrics.fp16_ops);
                debug!("  FP32 ops: {}", metrics.fp32_ops);
                debug!("  Tensor ops: {}", metrics.tensor_ops);
            }
            if metrics.estimated_flops > 0 {
                debug!("  Estimated TFLOPS: {:.2}", metrics.estimated_tflops);
                debug!("  Active cycles: {}", metrics.active_cycles);
                debug!("  Active warps: {}", metrics.active_warps);
                debug!("  SM efficiency: {}", metrics.sm_efficiency);
            }

            // Get accounting stats with better error reporting
            let accounting_stats = match NVML.as_ref() {
                Ok(nvml) => match nvml.device_by_index(0) {
                    Ok(device) => {
                        if !device.is_accounting_enabled().unwrap_or(false) {
                            warn!("GPU accounting is not enabled - no per-process stats will be available");
                            None
                        } else {
                            match device.accounting_stats_for(process_id).ok() {
                                Some(stats) => {
                                    debug!("Got accounting stats: GPU util {}%, Memory util {}%, Max memory {}MB",
                                        stats.gpu_utilization.unwrap_or(0),
                                        stats.memory_utilization.unwrap_or(0),
                                        stats.max_memory_usage.unwrap_or(0) / 1024 / 1024);
                                    Some(stats)
                                },
                                None => {
                                    warn!("No accounting stats available for process {}", process_id);
                                    None
                                }
                            }
                        }
                    },
                    Err(e) => {
                        warn!("Failed to get GPU device: {}", e);
                        None
                    }
                },
                Err(e) => {
                    warn!("NVML not initialized: {}", e);
                    None
                }
            };

            metrics.accounting_stats = accounting_stats;
            Ok(metrics)
        }).await?;

        // Set total_flops after the CUDA thread completes
        self.total_flops = metrics.measured_flops;

        // Do async operations after CUDA thread completes
        let gpu_info = &get_gpu_info().unwrap()[0];
        let serializable_metrics = SerializableGpuMetrics::from(&metrics);
        let details = serde_json::to_string(&serializable_metrics)
            .unwrap_or_else(|e| {
                warn!("Failed to serialize metrics: {}", e);
                format!("{}", metrics)
            });

        let log_entry = LogEntry {
            topic: GPU_MONITOR_TOPIC.to_string(),
            resource_provider_id: self.resource_provider_id.clone(),
            deal_id: self.deal_id.clone(),
            metadata: HashMap::new(),
            timestamp: Utc::now().to_rfc3339(),
            action: CUPTI_METRICS.to_string(),
            source: gpu_info.uuid.to_string(),
            details,
        };
        
        if let Err(e) = self.log_tx.send_async(log_entry).await {
            warn!("Failed to send log entry: {}", e);
        }

        tokio::time::sleep(Duration::from_millis(100)).await;
        Ok(metrics)
    }

    pub fn get_accounting_stats(&self) -> Option<AccountingStats> {
        if let Ok(nvml) = NVML.as_ref() {
            if let Ok(device) = nvml.device_by_index(0) {
                if device.is_accounting_enabled().unwrap_or(false) {
                    return device.accounting_stats_for(self.process_id).ok();
                }
            }
        }
        None
    }

    // Add a method to check context validity
    fn check_context(&self) -> bool {
        unsafe {
            let mut current = ptr::null_mut();
            let result = cuCtxGetCurrent(&mut current);
            result == cudaError_enum::CUDA_SUCCESS && !current.is_null()
        }
    }
}

impl Drop for CuptiTflopsCalculator {
    fn drop(&mut self) {
        unsafe {
            // Try to clean up event groups
            for group in &self.event_groups {
                let _ = cuptiEventGroupDestroy(group.0);  // Use .0 to access raw pointer
            }

            // Try to clean up context if it's still valid
            if self.check_context() {
                let _ = cuCtxDestroy_v2(self.context);
            }
        }
    }
}