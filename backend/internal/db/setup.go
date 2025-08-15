package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	connStr := "postgres://dustinmeyer@localhost:5432/cwf"

	var err error
	Pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("unable to create Postgres pool: %v", err)
	}

	Pool.Config().MaxConns = 25
	Pool.Config().MinConns = 5

	err = Pool.Ping(context.Background())

	if err != nil {
		log.Fatal("unable to ping database: %v", err)
	}

	log.Println("Postgres connection pool initialized")

	createTables()

}

func createTables() {
	tables := []string{"account/001_CREATE_TABLE_ACCOUNT.UP.SQL", "hive/002_CREATE_TABLE_HIVE.UP.SQL", "chat/003_CREATE_TABLE_CHAT.UP.SQL", "message/004_CREATE_TABLE_MESSAGE.UP.SQL"}

	query := []byte{}

	for _, file := range tables {
		data, _ := os.ReadFile(fmt.Sprintf("internal/db/migrations/%s", file))
		query = append(query, data...)
	}

	pg, err := Pool.Exec(context.Background(), string(query))

	fmt.Println(pg, err)
}
