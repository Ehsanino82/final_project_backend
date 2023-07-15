-- name: GetPersonByUsername :one
SELECT *
FROM person
WHERE username = $1;

-- name: CreatePerson :one
INSERT INTO person(username, password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCredit :exec
UPDATE person
SET credit = $2
WHERE username = $1;