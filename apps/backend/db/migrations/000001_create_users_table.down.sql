ALTER TABLE user_tokens DROP CONSTRAINT FK_ut__users;

DROP TABLE user_tokens;

ALTER TABLE users DROP CONSTRAINT uq_email;

DROP TABLE users;