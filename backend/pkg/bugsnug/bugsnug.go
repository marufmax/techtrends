package bugsnug

import (
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/marufmax/techtrends/api/pkg/config"
)

func Configure() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          config.Env.BugSnagKey,
		ReleaseStage:    config.Env.AppEnv,
		ProjectPackages: []string{"main", "Run", "github.com/marufmax/techtrends/api"},
		AppVersion:      config.Env.AppVersion,
	})
}
