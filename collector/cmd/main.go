package main

import (
	"github.com/bugsnag/bugsnag-go/v2"
	collector "github.com/marufmax/techtrends"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	configureBugReport()

	app := &cli.App{
		Name:    "collector",
		Usage:   "collect technology trends from BDjobs",
		Version: "v1.0.0",
		Action: func(ctx *cli.Context) error {
			collector.Collect()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func configureBugReport() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          "700b3e0e1b7e29f7ced38f84d1c554ef",
		ReleaseStage:    os.Getenv("APP_ENV"),
		ProjectPackages: []string{"main", "github.com/marufmax/techtrends"},
	})
}