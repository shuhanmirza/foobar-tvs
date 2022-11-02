package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	db "foobar-go/sqlc"
)

// Store provides all functions to execute Db Queries and transactions
type Store struct {
	Queries *db.Queries
	Db      *sql.DB
}

// NewStore creates a new Store
func NewStore(sqlDb *sql.DB) *Store {
	return &Store{
		Queries: db.New(sqlDb),
		Db:      sqlDb,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := store.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v ---- rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
