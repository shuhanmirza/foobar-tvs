-- name: CreateEvent :one
INSERT INTO events ("name", location_id, datetime)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEventById :one
SELECT *
FROM events
WHERE id = $1
LIMIT 1;

-- name: GetEventWithLocationById :one
SELECT events.id         as id,
       events.name       as name,
       "datetime",
       location_id,
       locations.name    as location,
       events.created_at as created_at
FROM events
         INNER JOIN locations on locations.id = events.location_id
WHERE events.id = $1
LIMIT 1;

-- name: GetEventListByOffsetLimit :many
SELECT events.id         as id,
       events.name       as name,
       "datetime",
       location_id,
       locations.name    as location,
       events.created_at as created_at
FROM events
         INNER JOIN locations on locations.id = events.location_id
ORDER BY events.created_at
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