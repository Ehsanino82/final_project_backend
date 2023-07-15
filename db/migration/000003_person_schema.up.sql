CREATE TABLE person
(
    id       BIGSERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL,
    credit   INTEGER DEFAULT 0
);
