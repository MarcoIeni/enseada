use actix_web::get;
use actix_web::web::ServiceConfig;
use actix_web::web::{Data, Json};
use serde::Serialize;

use couchdb;
use couchdb::status::Status;

use crate::http::error::ApiError;
use crate::http::error::ApiError::ServiceUnavailable;
use crate::responses;

pub fn mount(cfg: &mut ServiceConfig) {
    cfg.service(get);
}

#[derive(Debug, Serialize, PartialEq)]
pub struct HealthResponse {
    pub status: String,
}

#[get("/health")]
pub async fn get(couch: Data<couchdb::Couch>) -> Result<Json<HealthResponse>, ApiError> {
    match couch.status().await {
        Ok(Status { status }) => responses::ok(HealthResponse { status }),
        Err(err) => {
            log::error!("{}", err);
            Err(ServiceUnavailable(
                "database connection refused".to_string(),
            ))
        }
    }
}
