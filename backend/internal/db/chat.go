package db

import (
	"context"
)

func SaveMessage(msg []byte) error {
	queryString := "INSERT INTO chat (msg) VALUES ($1)"

	_, err := Pool.Exec(context.Background(), queryString, msg)

	if err != nil {
		return err
	}

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
