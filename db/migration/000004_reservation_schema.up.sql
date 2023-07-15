CREATE TABLE reservation
(
    room_id INTEGER REFERENCES person(id),
    person_id INTEGER REFERENCES room(id),
    CONSTRAINT rooms_person_pk PRIMARY KEY(room_id,person_id),
    dates Integer[]
);