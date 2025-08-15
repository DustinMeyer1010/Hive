package db

import "context"

// Reads all the chat message from the database
func ReadAllChat() string {
	var alldata string
	rows, _ := pool.Query(context.Background(), "SELECT * FROM chat")

	for rows.Next() {
		var data []byte

		rows.Scan(&data)
		alldata += "\n"
		alldata += string(data)
	}

	return alldata

}
