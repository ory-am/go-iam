
CREATE TABLE "_hydra_oauth2_trusted_jwt_bearer_issuer_tmp"
(
    id             VARCHAR(36) PRIMARY KEY,
    issuer         VARCHAR(255) NOT NULL,
    subject        VARCHAR(255) NOT NULL,
    allowed_domain VARCHAR(255) NOT NULL,
    scope          TEXT         NOT NULL,
    key_set        varchar(255) NOT NULL,
    key_id         varchar(255) NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE (issuer, subject, allowed_domain, key_id),
    FOREIGN KEY (key_set, key_id) REFERENCES hydra_jwk (sid, kid) ON DELETE CASCADE
);

INSERT INTO "_hydra_oauth2_trusted_jwt_bearer_issuer_tmp"
SELECT id, issuer, subject, '', scope, key_set, key_id, created_at, expires_at FROM "hydra_oauth2_trusted_jwt_bearer_issuer";

DROP TABLE "hydra_oauth2_trusted_jwt_bearer_issuer";

ALTER TABLE "_hydra_oauth2_trusted_jwt_bearer_issuer_tmp" RENAME TO "hydra_oauth2_trusted_jwt_bearer_issuer";

CREATE INDEX hydra_oauth2_trusted_jwt_bearer_issuer_expires_at_idx ON hydra_oauth2_trusted_jwt_bearer_issuer (expires_at);
