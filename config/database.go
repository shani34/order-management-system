package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./orders.db?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		log.Println("Database connection error:", err)
	}

		// Enable WAL mode
		_, err = db.Exec("PRAGMA journal_mode=WAL;")
		if err != nil {
			log.Fatal("Failed to enable WAL mode:", err)
		}
	
		// Set a busy timeout (in ms) to retry if the database is locked
		_, err = db.Exec("PRAGMA busy_timeout = 5000;")
		if err != nil {
			log.Fatal("Failed to set busy timeout:", err)
		}
	return db
}
