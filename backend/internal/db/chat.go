package db

import (
	"context"
	"fmt"
)

func SaveMessage(msg []byte) error {
	queryString := "INSERT INTO chat (msg) VALUES ($1)"

	pgCommand, err := Pool.Exec(context.Background(), queryString, msg)

	fmt.Println(pgCommand)
	fmt.Println(err)

	return nil
}

func ReadAllChat() string {
	var alldata string
	rows, _ := Pool.Query(context.Background(), "SELECT * FROM chat")

	for rows.Next() {
		var data []byte

		rows.Scan(&data)
		alldata += "\n"
		alldata += string(data)
	}

	return alldata

}
