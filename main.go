package main

import (
	"booking-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = uuid.New()
	event.CreatorID = uuid.New()
	context.JSON(http.StatusCreated, gin.H{"message": "success", "event": event})
	models.SaveEvent(event)
}
