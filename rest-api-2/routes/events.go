package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api-2/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	stringId := context.Param("id")
	id, err := strconv.ParseInt(stringId, 10, 64)

	if err != nil {
		// bad request because maybe the value that is not convertible to integer is added
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Tey again later"})
		return
	}

	// context.JSOn will handle the pointer automatically
	// hence no need to dereference event beforehand
	context.JSON(http.StatusOK, event)
}
func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// dummy data for now
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}

// this will support partial update (e.g. just name)
func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	// check if the event exists in our database
	event, err := models.GetEventByID(id) // getting existing event

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Tey again later"})
		return
	}

	var updatedEvent map[string]interface{}

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// for inner function, it has to use anonymous function syntax
	assignStringField := func(fieldName string, target *string) bool {
		// updates["name"].(string) is a type assertion to check if we are getting string as a value
		if value, ok := updatedEvent[fieldName].(string); ok {
			*target = value // assigning updated field value to event struct
			return true
		} else if _, exists := updatedEvent[fieldName]; exists {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type for field '" + fieldName + "'"})
		}
		return true
	}

	// accessing a field directly (e.g., event.Name) does not return a pointer to that field.
	// Instead, event.Name provides the value of the field, not a pointer to it. To get a pointer to the Name field itself, you need to use &event.Name.
	if !assignStringField("name", &event.Name) || !assignStringField("description", &event.Description) || !assignStringField("location", &event.Location) {
		return // Exit if any field has an invalid type
	}

	if dateTimeStr, ok := updatedEvent["date_time"].(string); ok {
		parseDate, parseError := time.Parse(time.RFC3339, dateTimeStr)
		if parseError != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format for field 'date_time'"})
			return
		}
		event.DateTime = parseDate
	} else if _, exists := updatedEvent["date_time"]; exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type for field 'date_time'"})
		return
	}

	// dummy data for now
	event.ID = id

	_, err = event.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event. Try again later"})
		return
	}

	// http.StatusNoContent can be used if we don't want to send any message!
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

}

// this one expects the request body has all the necessary key of event (if something is missing it will return error.)
// func updateEvent(context *gin.Context) {
// 	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
// 		return
// 	}

// 	// check if the event exists in our database
// 	_, err = models.GetEventByID(id)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Tey again later"})
// 		return
// 	}

// 	var updatedEvent models.Event

// 	err = context.ShouldBindJSON(&updatedEvent)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	// dummy data for now
// 	updatedEvent.ID = id
// 	updatedEvent.UserID = 1

// 	_, err = updatedEvent.Update()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event. Try again later"})
// 		return
// 	}

// 	// http.StatusNoContent can be used if we don't want to send any message!
// 	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

// }

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch event. Tey again later"})
	}

	_, err = event.DELETE()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event. Try again later"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
