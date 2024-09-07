package routes

import (
	"booking-api/db"
	"booking-api/models"
	"booking-api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Save(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ValidateCredentials(&user)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func ValidateCredentials(u *models.User) error {
	query := "SELECT id, password, salt FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	var salt string
	err := row.Scan(&u.ID, &hashedPassword, &salt)

	if err != nil {
		return errors.New("Invalid credentials.")
	}

	if !utils.HashesMatch(u.Password, salt, hashedPassword) {
		return errors.New("Invalid credentials.")
	}

	return nil
}
