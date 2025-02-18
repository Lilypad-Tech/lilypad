pub mod gpu_tools {
    tonic::include_proto!("gpu_tools");
}

// pub mod cupti;
pub mod kafka;
pub mod logger;
pub mod monitor;
pub mod server;
pub mod service;
pub mod nvml;
use anyhow::{anyhow, Result};
pub use service::GpuService;
use std::process::Command;

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

// pub fn new(log_tx: Sender<LogEntry>, resource_provider_id: String, deal_id: String) -> Result<Self, CuptiError> {
//     // Check privileges first
//     Self::check_cupti_privileges()?;

//     // Check if accounting mode is enabled, if not, enable it
//     match Self::get_accounting_mode() {
//         Ok(mode) if mode != "Enabled" => {
//             if let Err(e) = Self::set_accounting_mode() {
//                 warn!(
//                     "Failed to enable GPU accounting mode: {}. Some metrics may not be available.",
//                     e
//                 );
//             }
//         }
//         Err(e) => {
//             warn!(
//                 "Failed to check GPU accounting mode: {}. Some metrics may not be available.",
//                 e
//             );
//         }
//         _ => {}
//     }
