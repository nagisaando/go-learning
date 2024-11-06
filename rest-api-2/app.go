package main

import (
	"net/http"
	"strconv"

	"example.com/rest-api-2/db"
	"example.com/rest-api-2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", createEvent)

	router.Run("localhost:8080")
}

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
	id, err := strconv.ParseFloat(stringId, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Tey again later"})
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
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}
