package config

import "fmt"

type PG struct {
	Database string `env:"PG_DB" env-required:"true"`
	Username string `env:"PG_USER" env-required:"true"`
	Password string `env:"PG_PASSWORD" env-required:"true"`
	Host     string `env:"PG_HOST" env-required:"true"`
	Port     string `env:"PG_PORT" env-required:"true"`
}

func (p PG) Dsn() string {
	return fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		p.Database,
		p.Username,
		p.Password,
		p.Host,
		p.Port,
	)
}
