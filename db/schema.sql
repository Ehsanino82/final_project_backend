CREATE TABLE hotel
(
    id       BIGSERIAL PRIMARY KEY,
    name     VARCHAR(64) NOT NULL,
    phone    VARCHAR(64) NOT NULL,
    location VARCHAR(64) NOT NULL
);

CREATE TABLE room
(
    id   BIGSERIAL PRIMARY KEY,
    size INTEGER,
    beds INTEGER,
    floor INTEGER
);

CREATE TABLE person
(
    id   BIGSERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE reservation
(
    room_id INTEGER REFERENCES person(id),
    person_id INTEGER REFERENCES room(id),
    CONSTRAINT rooms_person_pk PRIMARY KEY(room_id,person_id),
    dates Integer[]
);
