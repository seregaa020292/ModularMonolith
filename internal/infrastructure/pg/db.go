package pg

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg/transaction"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 10
	connMaxLifetime = 30 * time.Minute
	connMaxIdleTime = 5 * time.Minute
)

type DB struct {
	*sql.DB
}

func New(cfg config.PG) (*DB, func(), error) {
	db, err := sql.Open("postgres", cfg.Dsn())
	if err != nil {
		return nil, nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	closer := func() {
		_ = db.Close()
	}

	return &DB{DB: db}, closer, nil
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
