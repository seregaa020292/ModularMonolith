package pgsql

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql/transaction"
)

type DB struct {
	*sql.DB
}

func (d DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if tx := transaction.ExtractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return d.DB.ExecContext(ctx, query, args)
}

func (d DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if tx := transaction.ExtractTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return d.DB.QueryContext(ctx, query, args...)
}
