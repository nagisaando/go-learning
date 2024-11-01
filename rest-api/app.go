package main

import (
	"example.com/rest-api/event"
	"github.com/gin-gonic/gin"
)

func main() {
	// Engin that configures HTTP servers with Logger and Recovery middleware
	// it is initializing router
	router := gin.Default()

	router.GET("/events", event.GetEvents)
	router.GET("/events/:id", event.GetEventByID)
	router.POST("/events", event.PostEvent)

	router.Run("localhost:8020")

}
