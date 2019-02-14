package main

import (
	"time"

	"github.com/go-playground/validator"
)

// Event is a struct containing all data relative to an event.
// It will be used by GORM as the schema for the "events" table.
type Event struct {
	ID        uint64     `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Color     string     `json:"color" validate:"required"`
	Icon      string     `json:"icon" validate:"required"`
	Message   string     `json:"message" validate:"required"`
	Author    string     `json:"author" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// EventValidator is used to validate an event post on the API is correct.
type EventValidator struct {
	validator *validator.Validate
}

// Validate is used to validate an event post on the API is correct.
func (ev *EventValidator) Validate(i interface{}) error {
	return ev.validator.Struct(i)
}
