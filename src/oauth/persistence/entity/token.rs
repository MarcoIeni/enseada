use serde::{Deserialize, Serialize};
use crate::oauth::client::Client;
use crate::couchdb::guid::Guid;
use crate::oauth::Scope;
use url::Url;
use uuid::Uuid;
use crate::oauth::code::AuthorizationCode;
use crate::oauth::session::Session;
use crate::secure::SecureSecret;
use chrono::{DateTime, Utc, Duration};
use chrono::serde::ts_seconds;
use crate::oauth::token::{AccessToken, RefreshToken};

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct AccessTokenEntity {
    #[serde(rename = "_id")]
    id: Guid,
    #[serde(rename = "_rev", skip_serializing_if = "Option::is_none")]
    rev: Option<String>,
    session: Session,
    #[serde(with = "ts_seconds")]
    expiration: DateTime<Utc>,
}

impl AccessTokenEntity {
    pub fn build_guid(id: String) -> Guid {
        Guid::from(format!("access_token:{}", id))
    }
    pub fn new(sig: String, session: Session, expiration: DateTime<Utc>) -> AccessTokenEntity {
        let id = Self::build_guid(sig);
        AccessTokenEntity { id, rev: None::<String>, session, expiration, }
    }

    pub fn id(&self) -> &Guid {
        &self.id
    }

    pub fn rev(&self) -> Option<String> {
        self.rev.clone()
    }

    pub fn session(&self) -> &Session {
        &self.session
    }

    pub fn expiration(&self) -> &DateTime<Utc> {
        &self.expiration
    }

    pub fn expires_in(&self) -> Duration {
        self.expiration.signed_duration_since(Utc::now())
    }

    pub fn to_token(&self, token: SecureSecret) -> AccessToken {
        let session = self.session.clone();
        let expires_in = self.expires_in().num_seconds() as u64;
        AccessToken::new(token, session, expires_in)
    }

    pub fn to_empty_token(&self) -> AccessToken {
        self.to_token(SecureSecret::empty())
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct RefreshTokenEntity {
    #[serde(rename = "_id")]
    id: Guid,
    #[serde(rename = "_rev", skip_serializing_if = "Option::is_none")]
    rev: Option<String>,
    session: Session,
    #[serde(with = "ts_seconds")]
    expiration: DateTime<Utc>,
}

impl RefreshTokenEntity {
    pub fn build_guid(id: String) -> Guid {
        Guid::from(format!("refresh_token:{}", id))
    }
    pub fn new(sig: String, session: Session, expiration: DateTime<Utc>) -> RefreshTokenEntity {
        let id = Self::build_guid(sig);
        RefreshTokenEntity { id, rev: None::<String>, session, expiration, }
    }

    pub fn id(&self) -> &Guid {
        &self.id
    }

    pub fn rev(&self) -> Option<String> {
        self.rev.clone()
    }

    pub fn session(&self) -> &Session {
        &self.session
    }

    pub fn expiration(&self) -> &DateTime<Utc> {
        &self.expiration
    }

    pub fn expires_in(&self) -> Duration {
        self.expiration.signed_duration_since(Utc::now())
    }

    pub fn to_token(&self, token: SecureSecret) -> RefreshToken {
        let session = self.session.clone();
        let expires_in = self.expires_in().num_seconds() as u64;
        RefreshToken::new(token, session, expires_in)
    }

    pub fn to_empty_token(&self) -> RefreshToken {
        self.to_token(SecureSecret::empty())
    }
}