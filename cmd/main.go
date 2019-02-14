package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	db     *gorm.DB
	config *Config
	logger *log.Logger
	themes Themes
	err    error
	e      *echo.Echo
)

func init() {

	// Init logger
	logger = log.New()
	logger.SetFormatter(&Formatter{})
	logger.SetLevel(log.InfoLevel)
	logger.Info("Starting techline server")

	// Load config
	config = &Config{Logger: logger}
	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		logger.Panic("cannot load config file because environment variable CONFIG_FILE has not been set")
	}
	err = config.LoadFromFile(configFile)
	if err != nil {
		logger.Panicf("cannot load config file: %+v", err)
	}

	// Init connection to DB
	logger.Infof("Connecting to Postgre database on %s:%d", config.DBHostname, config.DBPort)
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", config.DBHostname, config.DBPort, config.DBName, config.DBUser, config.DBPassword, config.DBSSLMode))
	if err != nil {
		logger.Panicf("cannot open DB: %+v", err)
	}
	db.LogMode(false)

	// Load preconfigured themes
	themesFile, ok := os.LookupEnv("THEMES_FILE")
	if !ok {
		logger.Panic("cannot load preconfigured themes file because environment variable THEMES_FILE has not been set")
	}
	err = themes.LoadFromFile(themesFile)
	if err != nil {
		logger.Panicf("cannot load themes: %+v", err)
	}

	// Init DB tables
	if config.DBReset {
		db.DropTableIfExists(&Event{})
	}
	db.AutoMigrate(&Event{})

	// Init server
	e = echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Server.Addr = fmt.Sprintf(":%d", config.ServerPort)

	// Add handlers to server
	v1 := e.Group("/api/v1")
	v1.GET("/events", GetEvents(db))
	v1.GET("/themes", GetThemes(themes))
	v1.POST("/events", PostEvent(db, logger))
	v1.DELETE("/events/:id", DeleteEvent(db))

	// Add validators to server
	e.Validator = &EventValidator{validator: validator.New()}

	// Add Frontend to server
	e.Static("/", config.WebDirectory)
}

func main() {

	defer db.Close()

	// Start server
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
