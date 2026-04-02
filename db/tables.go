package db

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		last_login TEXT NOT NULL
	)
	`

	_, err := db.Exec(userTable)

	if err != nil {
		panic(fmt.Sprintf("Could't create users table, err: %e", err))
	}

	fmt.Print("user table created successfully")

	return nil
}
