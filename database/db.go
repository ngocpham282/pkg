package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	db *pgx.Conn
}

func (db *DB) InTx(ctx context.Context, isoLevel pgx.TxIsoLevel, fn func(tx pgx.Tx) error) error {
	tx, err := db.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: isoLevel})
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	return tx.Commit(ctx)
}
