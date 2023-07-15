CREATE TABLE hotel
(
    id       BIGSERIAL PRIMARY KEY,
    name     VARCHAR(64) NOT NULL,
    phone    VARCHAR(64) NOT NULL,
    location VARCHAR(64) NOT NULL
);
