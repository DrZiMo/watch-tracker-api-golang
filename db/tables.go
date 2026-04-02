package db

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	userTable := `
	CREATE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		last_login TEXT NOT NULL
	)
	`

	_, err := db.Exec(userTable)

	if err != nil {
		return err
	}

	return nil
}
