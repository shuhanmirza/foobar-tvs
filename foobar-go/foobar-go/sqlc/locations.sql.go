// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: locations.sql

package db

import (
	"context"
)

const createLocation = `-- name: CreateLocation :one
INSERT INTO locations ("name")
VALUES ($1)
RETURNING id, name, created_at
`

func (q *Queries) CreateLocation(ctx context.Context, name string) (Locations, error) {
	row := q.db.QueryRowContext(ctx, createLocation, name)
	var i Locations
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteLocationById = `-- name: DeleteLocationById :exec
DELETE
    FROM locations
    WHERE id = $1
`

func (q *Queries) DeleteLocationById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLocationById, id)
	return err
}

const getLocationById = `-- name: GetLocationById :one
SELECT id, name, created_at
from locations
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetLocationById(ctx context.Context, id int64) (Locations, error) {
	row := q.db.QueryRowContext(ctx, getLocationById, id)
	var i Locations
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listLocations = `-- name: ListLocations :many
SELECT id, name, created_at
FROM locations
ORDER BY id
`

func (q *Queries) ListLocations(ctx context.Context) ([]Locations, error) {
	rows, err := q.db.QueryContext(ctx, listLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Locations{}
	for rows.Next() {
		var i Locations
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const updateLocation = `-- name: UpdateLocation :one
UPDATE locations
SET "name" = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateLocationParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) (Locations, error) {
	row := q.db.QueryRowContext(ctx, updateLocation, arg.ID, arg.Name)
	var i Locations
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
