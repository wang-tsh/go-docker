package main

import (
	"github.com/gin-gonic/gin"
	"go-docker/config"
	"go-docker/log"
	"go-docker/routes"
)

func main() {
	appConfig := config.AppConfig

	gin.SetMode(gin.ReleaseMode)
	log.ConfigLocalFilesystemLogger(appConfig.LogPath, appConfig.LogFileName, appConfig.MaxAge, appConfig.RotationTime)
	routes.Api().Run(":8088")
}
