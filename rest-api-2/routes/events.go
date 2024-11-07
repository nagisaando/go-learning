package routes

import (
	"net/http"
	"strconv"

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

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	// check if the event exists in our database
	_, err = models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Tey again later"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// dummy data for now
	updatedEvent.UserID = 1

	_, err = updatedEvent.Update(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event. Try again later"})
		return
	}

	// http.StatusNoContent can be used if we don't want to send any message!
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

}
