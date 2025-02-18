use std::time::Duration;
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response, Status};
use tracing::debug;

// Import from the library crate
use crate::gpu_tools::{ 
    gpu_tools_server::GpuTools, InitRequest, InitResponse, StartProfilingRequest, StartProfilingResponse,
    StopProfilingRequest, StopProfilingResponse, UpdateJwtTokenRequest, UpdateJwtTokenResponse,
};
use crate::logger::LOGGER;
use crate::service::GpuService;

#[derive(Clone)]
pub struct GrpcServer {
    pub gpu_service: Arc<RwLock<GpuService>>,
}

impl GrpcServer {
    pub async fn shutdown(&self) {
        if let Err(e) = self.gpu_service.write().await.shutdown().await {
            debug!("Error during shutdown: {}", e);
        }
    }
}

#[tonic::async_trait]
impl GpuTools for GrpcServer {
    async fn init(
        &self,
        request: Request<InitRequest>,
    ) -> Result<Response<InitResponse>, Status> {
        let req = request.into_inner();
        let jwt_token = req.jwt_token;
        let rp_id = req.rp_id;

        // call gpu service set_resource_provider_id
        self.gpu_service.write().await.set_resource_provider_id(rp_id).await;

        // call logger init_with_token
        if let Ok(logger) = &*LOGGER {
            if let Err(e) = logger.init_with_token(Some(&jwt_token)) {
                debug!("Error initializing logger with token: {}", e);
                return Err(Status::internal("Failed to initialize logging"));
            }
        }

        Ok(Response::new(InitResponse { success: true }))
    }
    
    async fn update_jwt_token(
        &self,
        request: Request<UpdateJwtTokenRequest>,
    ) -> Result<Response<UpdateJwtTokenResponse>, Status> {
        let new_token = request.into_inner().jwt_token;
        debug!(
            "Updating JWT token (last 7 chars): '...{}'",
            new_token
                .chars()
                .rev()
                .take(7)
                .collect::<String>()
                .chars()
                .rev()
                .collect::<String>()
        );

        // call logger update_jwt_token
        if let Ok(logger) = &*LOGGER {
            if let Err(e) = logger.update_jwt_token(Some(&new_token)) {
                debug!("Error updating JWT token: {}", e);
                return Err(Status::internal("Failed to update JWT token"));
            }
        }

        Ok(Response::new(UpdateJwtTokenResponse { success: true }))
    }

    async fn start_profiling(
        &self,
        request: Request<StartProfilingRequest>,
    ) -> Result<Response<StartProfilingResponse>, Status> {
        debug!("Starting profiling");
        let req = request.into_inner();
        let deal_id = req.deal_id;

        // Create and start monitoring session
        let monitor_session = self.gpu_service
            .write()
            .await
            .create_monitoring_session(deal_id.clone())
            .await
            .map_err(|e| Status::internal(format!("Failed to create monitoring session: {}", e)))?;
            
        let mut monitor = monitor_session.write().await;
        monitor.start()
            .await
            .map_err(|e| Status::internal(format!("Failed to start monitoring: {}", e)))?;
        debug!("Monitoring started");

        // Create and start measurement session
        // let measurement_session = self.gpu_service
        //     .write()            .await
        //     .create_measurement_session(deal_id.clone())
        //     .await
        //     .map_err(|e| Status::internal(format!("Failed to create measurement session: {}", e)))?;
            
        // let mut measurement = measurement_session.write().await;
        // measurement.start_measurement()
        //     .await
        //     .map_err(|e| Status::internal(format!("Failed to start measurement: {}", e)))?;
        // debug!("Measurement started");

        Ok(Response::new(StartProfilingResponse { success: true }))
    }

    async fn stop_profiling(
        &self,
        request: Request<StopProfilingRequest>,
    ) -> Result<Response<StopProfilingResponse>, Status> {
        debug!("Stopping profiling");
        tokio::time::sleep(Duration::from_secs(1)).await; 
        let deal_id = request.into_inner().deal_id;

        // find the measurement session and stop it
        // let measurement_session = self.gpu_service
        //     .write()
        //     .await
        //     .get_measurement_session(deal_id.clone())
        //     .await
        //     .map_err(|e| Status::internal(format!("Failed to get measurement session: {}", e)))?;

        // let mut measurement = measurement_session.write().await;
        // if let Err(e) = measurement.stop_measurement().await {
        //     debug!("Error stopping measurement: {}", e);
        //     // Continue execution to ensure we still try to stop monitoring and flush logs
        // }

        // find the monitoring session and stop it
        let monitoring_session = self.gpu_service
            .write()
            .await
            .get_monitoring_session(deal_id.clone())
            .await
            .map_err(|e| Status::internal(format!("Failed to get monitoring session: {}", e)))?;

        let mut monitoring = monitoring_session.write().await;
        monitoring.stop().await;

        // flush the logs
        if let Ok(logger) = &*LOGGER {
            if let Err(e) = logger.flush(Duration::from_secs(1)).await {
                debug!("Error flushing logs: {}", e);
            }
        }
        Ok(Response::new(StopProfilingResponse { success: true }))
    }
}
