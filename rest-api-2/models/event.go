package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
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

func (event Event) Save() {
	// [TODO]: add it to database
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
