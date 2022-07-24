package app

import (
	"github.com/marufmax/techtrends/api/pkg/bugsnug"
	"github.com/marufmax/techtrends/api/pkg/config"
)

func Run() {
	bugsnug.Configure()
	e := New()

	e.Logger.Fatal(e.Start(":" + config.Env.AppPort))
}
