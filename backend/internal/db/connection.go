package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

type Connection struct {
	Config pgx.ConnConfig
	conn   *pgx.Conn
}

var db Connection

func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
		fmt.Println("Disconnected from PostgeSQL")
	}
}

func (c *Connection) Connect() {
	var err error

	c.conn, err = pgx.Connect(c.Config)

	if err != nil {
		log.Fatal("Error opening DB: ", err)
	}

	err = c.conn.Ping(context.Background())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}

func CreateConnection(config pgx.ConnConfig) error {
	db.Config = config
	db.Connect()

	return nil
}

func Close() error {
	db.Close()
	return nil
}
