use nvml_wrapper::{error::NvmlError, Nvml};
use std::sync::LazyLock;

pub static NVML: LazyLock<Result<Nvml, NvmlError>> = LazyLock::new(|| Nvml::init());
