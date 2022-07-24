package main

import (
	"github.com/marufmax/techtrends/api/internal/app"
	"github.com/marufmax/techtrends/api/pkg/config"
)

func main() {
	config.LoadEnvironmentFile(".env")
	app.Run()
}
