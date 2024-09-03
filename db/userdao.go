package db

import (
	"booking-api/models"
	"database/sql"
)

func Save(user models.User) error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

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

	result, err := statement.Exec(user.Email, user.Password)

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}
