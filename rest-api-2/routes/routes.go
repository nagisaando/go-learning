package routes

import (
	"example.com/rest-api-2/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	// [Adding middleware]
	// Option 1:
	// with Gin, multiple handlers can be registered and it will execute from left to right.
	// so middleware Authenticate can run before createEvent
	server.POST("/events", middleware.Authenticate, createEvent)

	// Option 2:
	// or we can add middleware by group:
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)    // setting middleware
	authenticated.PUT("/events/:id", updateEvent) // now middleware will be executed before these handler executes
	authenticated.DELETE("events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
