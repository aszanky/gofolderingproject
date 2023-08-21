package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	PostgresqlHost     = "localhost"
	PostgresqlPort     = "5432"
	PostgresqlUser     = "postgres"
	PostgresqlPassword = "postgres"
	PostgresqlDbname   = "ewallet"
	PostgresqlSslmode  = false
	PgDriver           = "postgres"
	maxOpenConns       = 60
	connMaxLifetime    = 120
	maxIdleConns       = 30
	connMaxIdleTime    = 20
)

//NewDatabase() PostgreSQL initialization
func NewDatabase() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		PostgresqlHost,
		PostgresqlPort,
		PostgresqlUser,
		PostgresqlDbname,
		PostgresqlPassword,
	)

	db, err := sqlx.Connect(PgDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
