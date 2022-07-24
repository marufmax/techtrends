package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Enviroment struct {
	AppEnv     string
	DBDSN      string
	BugSnagKey string
	AppPort    int
	AppVersion string
}

var Env *Enviroment

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok && defaultValue != "" {
		return defaultValue
	}

	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Enviroment)
	}

	Env.DBDSN = getEnv("DB_DSN", "host=127.0.0.1 user=techtrend password=trendy dbname=techtrend_local port=5435 sslmode=disable")
	Env.AppEnv = getEnv("APP_ENV", "production")
	Env.BugSnagKey = getEnv("BUGSNAG_KEY", "")
	Env.AppPort, _ = strconv.Atoi(getEnv("APP_PORT", "5898"))
	Env.AppVersion = "v1.0.0"
}

func LoadEnvironmentFile(file string) {
	if err := godotenv.Load(file); err != nil {
		fmt.Printf("Error on load environment file: %s", file)
	}

	LoadEnvironment()
}
