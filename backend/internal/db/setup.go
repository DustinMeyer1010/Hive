package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	connStr := "postgres://dustinmeyer@localhost:5432/cwf"

	var err error
	Pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("Unable to create Postgres pool: %v", err)
	}

	Pool.Config().MaxConns = 25
	Pool.Config().MinConns = 5

	err = Pool.Ping(context.Background())

	if err != nil {
		log.Fatal("Unable to ping database: %v", err)
	}

	log.Println("Postgres connection pool initialized")

}
