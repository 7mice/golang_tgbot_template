package main

import (
	"github.com/gin-gonic/gin"
	"goTgTemplate/app"
	"goTgTemplate/internal/config"
	"goTgTemplate/pkg/logging"
)

func main() {
	if config.DefaultConfig.IsDebug {
		logging.DefaultLogger.Info().Msg("Starting application as debug")
	} else {
		gin.SetMode(gin.ReleaseMode)
		logging.DefaultLogger.Info().Msg("Starting application")
	}
	app.Start()
}
