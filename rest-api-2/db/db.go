package db

import (
	// underline before import path tells GO to keep the path though it is not used directly
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// api.db is local file and it will be created automatically if it does not exist
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// CREATE TABLE IF NOT EXISTS => SQL command
	// event => table name
	// PRIMARY KEY => enforce each event will have unique id
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}
}
