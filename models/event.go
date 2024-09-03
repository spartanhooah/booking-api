package models

import (
	"booking-api/db"
	"database/sql"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	CreatorID   int64
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		sqlErr := rows.Close()
		if sqlErr != nil {
			err = sqlErr
		}
	}(rows)

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.CreatorID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id string) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.CreatorID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func SaveEvent(event Event) error {
	query := `
	INSERT INTO events(name, description, location, date_time, creator_id)
	VALUES (?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(query)

	defer func(statement *sql.Stmt) {
		sqlErr := statement.Close()
		if sqlErr != nil {
			err = sqlErr
		}
	}(statement)

	if err != nil {
		return err
	}

	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.CreatorID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func UpdateEvent(id string, event Event) error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, date_time = ?, creator_id = ? WHERE id = ?"

	statement, err := db.DB.Prepare(query)

	defer func(statement *sql.Stmt) {
		sqlErr := statement.Close()
		if sqlErr != nil {
			err = sqlErr
		}
	}(statement)

	if err != nil {
		return err
	}

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.CreatorID, id)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	query := "DELETE FROM events WHERE id = ?"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		sqlErr := statement.Close()
		if sqlErr != nil {
			err = sqlErr
		}
	}(statement)

	_, err = statement.Exec(id)

	return err
}
