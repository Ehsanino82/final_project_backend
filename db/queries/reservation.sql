-- name: GetRoomsReservedDates :many
SELECT dates
FROM reservation
WHERE room_id = $1;
