package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type environment struct {
	AppEnv       string
	DatabaseURL  string
	BugSnagKey   string
	AppPort      string
	AllowOrigins []string
}

var Env *environment

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok && defaultValue != "" {
		return defaultValue
	}

	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(environment)
	}

	Env.DatabaseURL = getEnv("DB_URL", "host=127.0.0.1 user=techtrend password=trendy dbname=techtrend_local port=5435 sslmode=disable")
	Env.AppEnv = getEnv("APP_ENV", "production")
	Env.BugSnagKey = getEnv("BUGSNAG_KEY", "")
	Env.AppPort = getEnv("APP_PORT", "5898")
	Env.AllowOrigins = strings.Split(getEnv("ALLOW_DOMAINS", "*"), ",")
}

func LoadEnvironmentFile(file string) {
	if err := godotenv.Load(file); err != nil {
		fmt.Printf("Error on load environment file: %s", file)
	}

	LoadEnvironment()
}
