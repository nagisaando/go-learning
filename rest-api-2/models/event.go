package models

import (
	"time"

	"example.com/rest-api-2/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

// type Event struct {
// 	ID          int
// 	Name        string `binding:"required"`
// 	Description string `binding:"required"`
// 	Location    string
// 	DateTime    time.Time
// 	UserID      int
// }

var events = []Event{}

func (event *Event) Save() error {

	// ? is parameter in SQLite
	query := `
	INSERT INTO events (name, description, location, dateTime, userId)
	VALUES(?, ?, ?, ?, ?)
	`

	// Prepare compiles the SQL command but does not execute it immediately. It creates a "template" that can be executed with parameters
	// it can use the both db.DB.prepare and db.DB.execute() but
	// Prepare() can lead to better performance in certain situations such as when we want to execute the same query multiple times with different values
	// UNLESS statement.Close() is not called between the executions. otherwise there is no advantages
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close() // this will be executed at the end of Save function

	// Exec()  runs the prepares statement
	// Exec() is typically used to update a data while Query() is used to read a data
	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // get auto generated id

	if err != nil {
		return err
	}

	event.ID = id
	events = append(events, *event)

	return nil
}

func GetAllEvents() ([]Event, error) {

	// we can query all keys using * but it is recommended for testing purpose only: https://www.sqlitetutorial.net/sqlite-select/

	query := `
	SELECT
		id,
		name,
		description,
		location,
		dateTime,
		userId
	FROM
		events
	`

	var events []Event
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		// we are reading data from row and assigning them to event struct
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	// if there is any issue during iteration of rows.Next(),
	// error will be stored in rows.Err()
	if err = rows.Err(); err != nil {
		return events, err
	}

	return events, nil
}

// returning pointer for memory efficiency
func GetEventByID(id int64) (*Event, error) {

	// we can query all keys using * but it is recommended for only testing purpose: https://www.sqlitetutorial.net/sqlite-select/
	query := `SELECT * FROM events WHERE id = ?`
	// QueryRow() returns only one row
	// no need to close the connection
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() (int64, error) {

	query := `
	UPDATE 
		events 
	SET
		name = ?,
		description = ?,
		location = ? ,
		dateTime = ?
	WHERE
		id = ?
	`

	result, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected() // it will return 0 if no row is updated
}

func (event Event) DELETE() (int64, error) {
	query := `
	DELETE FROM events
	WHERE id = ?
	`

	result, err := db.DB.Exec(query, event.ID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
