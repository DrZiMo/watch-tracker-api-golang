package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite", "api.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	CreateTables(DB)

	fmt.Println("Database initialized successfully")
}
