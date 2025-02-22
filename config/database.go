package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Println("Database connection error:", err)
	}
	return db
}
