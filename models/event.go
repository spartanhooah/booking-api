package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID          uuid.UUID
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	CreatorID   uuid.UUID
}

var events []Event

func (e *Event) Save() {
	// TODO: save to DB
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

func SaveEvent(event Event) {
	events = append(events, event)
}
