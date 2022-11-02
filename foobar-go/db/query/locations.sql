-- name: CreateLocation :one
INSERT INTO locations ("name")
VALUES ($1)
RETURNING *;

-- name: GetLocationById :one
SELECT *
from locations
WHERE id = $1
LIMIT 1;

-- name: ListLocations :many
SELECT *
FROM locations
ORDER BY id;

-- name: DeleteLocationById :exec
DELETE
    FROM locations
    WHERE id = $1;

-- name: UpdateLocation :one
UPDATE locations
SET "name" = $2
WHERE id = $1
RETURNING *;
