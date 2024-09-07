package routes

import (
	"booking-api/db"
	"booking-api/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getEvents(context *gin.Context) {
	events, err := db.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(200, events)
}

func getEvent(context *gin.Context) {
	id := context.Param("id")

	event, err := db.GetEvent(id)

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

	event.CreatorID = context.GetInt64("userId")

	err = db.SaveEvent(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "success", "event": event})
}

func updateEvent(context *gin.Context) {
	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	err, done := validateUser(context, id)

	if done {
		return
	}

	err = db.UpdateEvent(id, updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "success", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")

	err, done := validateUser(context, id)

	if done {
		return
	}

	err = db.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successful deletion"})
}

func validateUser(context *gin.Context, id string) (error, bool) {
	existingEvent, err := db.GetEvent(id)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			context.JSON(http.StatusNotFound, gin.H{"error": "No event found with id " + id})
			return nil, true
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, true
	}

	if context.GetInt64("userId") != existingEvent.CreatorID {
		context.JSON(http.StatusForbidden, gin.H{"error": "You can't update this event"})
		return nil, true
	}
	return err, false
}
