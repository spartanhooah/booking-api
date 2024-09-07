package routes

import (
	"booking-api/db"
	"booking-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(context *gin.Context) {
	userId, event, err := getUserIdAndEventId(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Register(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"event": event})
}

func cancelRegistration(context *gin.Context) {
	userId, event, err := getUserIdAndEventId(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.CancelRegistration(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration canceled"})
}

func getUserIdAndEventId(context *gin.Context) (int64, *models.Event, error) {
	userId := context.GetInt64("userId")
	eventId := context.Param("id")

	event, err := db.GetEvent(eventId)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			context.JSON(http.StatusNotFound, gin.H{"error": "No event found for that id"})
			return 0, nil, nil
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, nil, nil
	}
	return userId, event, err
}
