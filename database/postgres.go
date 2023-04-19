package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PostgresDB struct {
	BunDB *bun.DB
}

func NewPostgresDB() *PostgresDB {
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return &PostgresDB{BunDB: db}
}

func (pg *PostgresDB) Close() error {
	return pg.BunDB.Close()
}
