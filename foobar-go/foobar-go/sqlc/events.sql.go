// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: events.sql

package db

import (
	"context"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events ("name", location_id, datetime)
VALUES ($1, $2, $3)
RETURNING id, location_id, name, datetime, created_at
`

type CreateEventParams struct {
	Name       string `json:"name"`
	LocationID int64  `json:"location_id"`
	Datetime   int64  `json:"datetime"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Events, error) {
	row := q.db.QueryRowContext(ctx, createEvent, arg.Name, arg.LocationID, arg.Datetime)
	var i Events
	err := row.Scan(
		&i.ID,
		&i.LocationID,
		&i.Name,
		&i.Datetime,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEventById = `-- name: GetEventById :one
SELECT id, location_id, name, datetime, created_at
from events
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEventById(ctx context.Context, id int64) (Events, error) {
	row := q.db.QueryRowContext(ctx, getEventById, id)
	var i Events
	err := row.Scan(
		&i.ID,
		&i.LocationID,
		&i.Name,
		&i.Datetime,
		&i.CreatedAt,
	)
	return i, err
}

const getEventListByOffsetLimit = `-- name: GetEventListByOffsetLimit :many
SELECT id, location_id, name, datetime, created_at
from events
ORDER BY id
OFFSET $1 LIMIT $2
`

type GetEventListByOffsetLimitParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetEventListByOffsetLimit(ctx context.Context, arg GetEventListByOffsetLimitParams) ([]Events, error) {
	rows, err := q.db.QueryContext(ctx, getEventListByOffsetLimit, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Events{}
	for rows.Next() {
		var i Events
		if err := rows.Scan(
			&i.ID,
			&i.LocationID,
			&i.Name,
			&i.Datetime,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEvent = `-- name: UpdateEvent :one
UPDATE events
SET "name"=$2,
    location_id=$3,
    datetime=$2
WHERE id = $1
RETURNING id, location_id, name, datetime, created_at
`

type UpdateEventParams struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LocationID int64  `json:"location_id"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Events, error) {
	row := q.db.QueryRowContext(ctx, updateEvent, arg.ID, arg.Name, arg.LocationID)
	var i Events
	err := row.Scan(
		&i.ID,
		&i.LocationID,
		&i.Name,
		&i.Datetime,
		&i.CreatedAt,
	)
	return i, err
}