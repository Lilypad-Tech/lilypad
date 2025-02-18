use rdkafka::client::{ClientContext, OAuthToken};
use rdkafka::producer::ProducerContext;
use std::error::Error;
use std::time::{SystemTime, UNIX_EPOCH};
use tracing::{debug, error};

pub struct CustomContext {
    pub token: String,
}

impl ProducerContext for CustomContext {
    type DeliveryOpaque = ();

    fn delivery(
        &self,
        delivery_result: &rdkafka::producer::DeliveryResult,
        _delivery_opaque: Self::DeliveryOpaque,
    ) {
        match delivery_result {
            Ok(_) => debug!("Message delivered successfully"),
            Err((err, _)) => error!("Failed to deliver message: {}", err),
        }
    }
}

impl ClientContext for CustomContext {
    const ENABLE_REFRESH_OAUTH_TOKEN: bool = true;

    fn generate_oauth_token(&self, _config: Option<&str>) -> Result<OAuthToken, Box<dyn Error>> {
        let now = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis() as i64;

        Ok(OAuthToken {
            token: self.token.clone(),
            principal_name: "".to_string(),
            lifetime_ms: now + (3600 * 1000), // current time + 1 hour
        })
    }
}
