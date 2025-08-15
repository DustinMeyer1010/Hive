package db

import (
	"context"
)

// Save the provided message to the database
func SaveMessage(msg []byte) error {
	queryString := "INSERT INTO chat (msg) VALUES ($1)"

	_, err := pool.Exec(context.Background(), queryString, msg)

	if err != nil {
		return err
	}

	return nil
}
