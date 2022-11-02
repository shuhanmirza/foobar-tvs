-- name: CreateEvent :one
INSERT INTO events ("name", location_id, datetime)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEventById :one
SELECT *
from events
WHERE id = $1
LIMIT 1;

-- name: GetEventListByOffsetLimit :many
SELECT *
from events
ORDER BY id
OFFSET $1 LIMIT $2;

-- name: UpdateEvent :one
UPDATE events
SET "name"=$2,
    location_id=$3,
    datetime=$2
WHERE id = $1
RETURNING *;


-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1;