package pg

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 10
	connMaxLifetime = 30 * time.Minute
	connMaxIdleTime = 5 * time.Minute
)

type Config struct {
	Database string `env:"PG_DB" env-required:"true"`
	Username string `env:"PG_USER" env-required:"true"`
	Password string `env:"PG_PASSWORD" env-required:"true"`
	Host     string `env:"PG_HOST" env-required:"true"`
	Port     string `env:"PG_PORT" env-required:"true"`
}

func (c Config) Dsn() string {
	return fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		c.Database,
		c.Username,
		c.Password,
		c.Host,
		c.Port,
	)
}

func New(cfg Config) (*pgsql.DB, func(), error) {
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

	return &pgsql.DB{DB: db}, closer, nil
}
