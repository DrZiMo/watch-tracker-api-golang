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
		last_login TEXT NOT NULL,

		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := db.Exec(userTable)

	if err != nil {
		panic(fmt.Sprintf("Could't create users table, err: %e", err))
	}

	watchTable := `
	CREATE TABLE IF NOT EXISTS watches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status TEXT NOT NULL,
		rating REAL CHECK(rating >= 0 AND rating <= 10),
		user_id INTEGER NOT NULL,
		
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = db.Exec(watchTable)

	if err != nil {
		panic(fmt.Sprintf("Could't create watches table, err: %e", err))
	}

	return nil
}
