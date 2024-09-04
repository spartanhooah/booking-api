package db

import (
	"booking-api/models"
	"booking-api/utils"
	"database/sql"
)

func Save(user models.User) error {
	query := "INSERT INTO users(email, password, salt) VALUES (?, ?, ?)"

	statement, err := DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		sqlErr := statement.Close()
		if sqlErr != nil {
			err = sqlErr
		}
	}(statement)

	hashedPw, salt, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPw, salt)

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}
