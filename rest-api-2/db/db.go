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

	// UNIQUE => it allows one unique address in a table (different user can not have the same email address)

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL 
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	// CREATE TABLE IF NOT EXISTS => SQL command
	// event => table name
	// PRIMARY KEY => enforce each event will have unique id

	// FOREIGN KEY(user_id) REFERENCES users(id) => user_id will be referenced by users table, specifically id
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
		eventId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id),
		FOREIGN KEY(eventId) REFERENCES events(id) 
	)`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table")
	}
}
