use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::transport::Server;
use tracing::{debug, info};
use tracing_subscriber::{fmt, EnvFilter};
use tokio::signal::unix::{signal, SignalKind};

use gpu_tools::{
    gpu_tools::gpu_tools_server::GpuToolsServer, 
    server::GrpcServer,
    service::GpuService,
    monitor::LOG_CHANNEL
};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Initialize tracing with default environment filter
    fmt()
        .with_env_filter(
            EnvFilter::from_default_env()
                .add_directive(tracing::Level::DEBUG.into())
                .add_directive("h2=error".parse()?)
        )
        .init();

    debug!("Initializing GPU Tools service...");

    let log_tx = LOG_CHANNEL.get_sender();
    let gpu_service = GpuService::init(
        log_tx.clone(),
    ).await?;

    let server = GrpcServer {
        gpu_service: Arc::new(RwLock::new(gpu_service))
    };

    // Start gRPC server - listen on all interfaces
    let addr = "[::0]:50051".parse()?;
    info!("Starting gRPC server on {}", addr);
    
    Server::builder()
        .add_service(GpuToolsServer::new(server.clone()))
        .serve_with_shutdown(addr, async {
            let mut sigterm = match signal(SignalKind::terminate()) {
                Ok(sig) => sig,
                Err(e) => {
                    debug!("Failed to register SIGTERM handler: {}", e);
                    return;
                }
            };
            let mut sigint = match signal(SignalKind::interrupt()) {
                Ok(sig) => sig,
                Err(e) => {
                    debug!("Failed to register SIGINT handler: {}", e);
                    return;
                }
            };

            tokio::select! {
                _ = sigterm.recv() => {
                    debug!("SIGTERM received");
                }
                _ = sigint.recv() => {
                    debug!("SIGINT received");
                }
            }
            debug!("Shutdown signal received");
            server.shutdown().await;
        })
        .await?;

    debug!("Server shutdown complete");
    Ok(())
}
