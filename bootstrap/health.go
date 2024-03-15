package bootstrap

import (
	"time"

	"github.com/lengocson131002/go-clean-core/config"
	health "github.com/lengocson131002/go-clean-core/health"
)

func NewHealthChecker(srv *ServerConfig, cfg config.Configure) health.HealthChecker {

	var (
		grRunningThreshold    = cfg.GetInt("HEALTH_CHECK_GR_RUNNING_THRESHOLD")
		gcMaxPauseThresholdms = cfg.GetInt("HEALTH_CHECK_GC_PAUSE_THRESHOLD_MS")
		envPath               = ".env"
	)

	// Init health
	healthChecker := health.NewHealthChecker(srv.Name, srv.AppVersion)

	// check Garbage Collector
	gcChecker := health.NewGarbageCollectionMaxChecker(time.Millisecond * time.Duration(gcMaxPauseThresholdms))
	healthChecker.AddLivenessCheck("garbage collector check", gcChecker)

	// check Goroutine
	grChecker := health.NewGoroutineChecker(grRunningThreshold)
	healthChecker.AddLivenessCheck("goroutine checker", grChecker)

	// check env file
	envFileChecker := health.NewEnvChecker(envPath)
	healthChecker.AddReadinessCheck("env file checker", envFileChecker)

	// check network
	pingChecker := health.NewPingChecker("http://google.com", "GET", time.Millisecond*time.Duration(200), nil, nil)
	healthChecker.AddReadinessCheck("ping check", pingChecker)

	return healthChecker
}
