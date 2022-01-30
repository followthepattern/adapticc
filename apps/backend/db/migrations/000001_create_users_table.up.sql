CREATE TABLE users (
  id VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  first_name VARCHAR,
  last_name VARCHAR,
  password_hash VARCHAR,
  salt VARCHAR,
  active BOOLEAN NOT NULL,
  creation_user_id VARCHAR,
  created_at TIMESTAMP NOT NULL,
  update_user_id VARCHAR,
  updated_at TIMESTAMP,
  last_login_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE users ADD CONSTRAINT uq_email UNIQUE (email);

CREATE TABLE user_tokens (
  user_id VARCHAR NOT NULL,
  token VARCHAR NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

ALTER TABLE user_tokens ADD CONSTRAINT FK_ut__users FOREIGN KEY (user_id) REFERENCES users (id);