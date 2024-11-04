package event

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var events = []event{
	{ID: "1", Title: "Event 1", Description: "This is event 1 description.", CreatedAt: time.Now()},
	{ID: "2", Title: "Event 2", Description: "This is event 2 description.", CreatedAt: time.Now()},
	{ID: "3", Title: "Event 3", Description: "This is event 3 description.", CreatedAt: time.Now()},
}

// *gin.Context carries request details =, validates and serializes JSON and more.
// it will be set by gin automatically when we set this function as handler
func GetEvents(c *gin.Context) {
	// creates JSON from the slice of event structs and writes (pretty) JSON into the response
	// this is recommended to do only for development since pretty JSON is expensive (it can be replace with Context.JSON to send more compact JSON)
	// c.JSON(http.StatusOK, events)
	c.IndentedJSON(http.StatusOK, events)
}

func PostEvent(c *gin.Context) {
	var newEvent event

	// BindJSON is to bind received JSON to newEvent
	// Gin is forgiving framework and if no value is filled for the field, (for example no id) gin will return the null value without throwing error
	err := c.BindJSON(&newEvent)
	// or err = c.ShouldBindJSON(&newEvent)
	newEvent.CreatedAt = time.Now()

	if err != nil {
		// gin.H{} is to create map
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	for _, val := range events {
		fmt.Println(id)
		fmt.Println(val.ID)
		if id == val.ID {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}
