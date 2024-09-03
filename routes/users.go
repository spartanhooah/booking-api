package routes

import (
	"booking-api/db"
	"booking-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Save(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}
