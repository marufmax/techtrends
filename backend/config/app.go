package config

import "time"

type app struct {
	AppVersion              string
	ApiDocPath              string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	GzipLevel               int
	DBMaxIdleConnection     int
	DBMaxConnectionLifeTime time.Duration
	DBMaxOpenConnection     int
}

var App *app

// CreateAppConfig creates the application configuration
func CreateAppConfig() {
	if App == nil {
		App = new(app)
	}

	// Application
	App.AppVersion = "v1.0.0"
	App.ApiDocPath = "/docs/*"
	App.ReadTimeout = 15 * time.Second
	App.WriteTimeout = 15 * time.Second
	App.GzipLevel = 5

	// Database
	App.DBMaxOpenConnection = 10
	App.DBMaxIdleConnection = 100
	App.DBMaxConnectionLifeTime = time.Hour
}
