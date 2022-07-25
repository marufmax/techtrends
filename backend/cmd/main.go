package main

import (
	"github.com/marufmax/techtrends/api/config"
	"github.com/marufmax/techtrends/api/internal/http"
)

func main() {
	config.LoadEnvironmentFile(".env")
	config.CreateAppConfig()
	http.Run()
}
