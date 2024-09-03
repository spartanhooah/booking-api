package routes

import (
	"booking-api/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(200, events)
}

func getEvent(context *gin.Context) {
	id := context.Param("id")

	event, err := models.GetEvent(id)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			context.JSON(http.StatusNotFound, gin.H{"error": "No event found with id " + id})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "success", "event": event})
	err = models.SaveEvent(event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func updateEvent(context *gin.Context) {
	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	err = models.UpdateEvent(id, updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")

	err := models.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successful deletion"})
}
