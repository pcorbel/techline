package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetEvents is an handler to return all events stored in the database.
// The returned list will be ordered by creation timestamp.
func GetEvents(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var events []Event
		db.Order("created_at desc").Find(&events)
		return c.JSON(http.StatusOK, events)
	}
}

// GetThemes is an handler to return all the possible themes for an event.
func GetThemes(themes Themes) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, themes)
	}
}

// PostEvent is an handler to insert/update an event into the database.
func PostEvent(db *gorm.DB, logger *log.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		evt := new(Event)
		if err := c.Bind(evt); err != nil {
			logger.Errorf("cannot bind event: %+v", err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(evt); err != nil {
			logger.Errorf("cannot bind event: %+v", err)
			return c.NoContent(http.StatusBadRequest)
		}
		db.Save(&evt)
		return c.NoContent(http.StatusOK)
	}
}

// DeleteEvent is an handler to soft delete an event into the database.
func DeleteEvent(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var evt Event
		db.Where("id = ?", c.Param("id")).Find(&evt)
		db.Delete(&evt)
		return c.NoContent(http.StatusOK)
	}
}
