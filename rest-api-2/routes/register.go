package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api-2/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error": "Could not parse id"})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Try again later"})
		return
	}

	userId := context.GetInt64("userId")

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could register user for the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})

}

func cancelForRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error": "Could not parse id"})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Try again later"})
		return
	}

	userId := context.GetInt64("userId")

	affectedRows, err := event.DeleteRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could register user for the event"})
		return
	}

	if affectedRows == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "No registration found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration deleted!"})

}
