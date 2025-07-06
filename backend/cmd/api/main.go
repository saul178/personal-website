package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/config"
	"github.com/saul178/personal-website/internal/server"
	"github.com/saul178/personal-website/middleware"
)

func main() {
	middleware.InitLogger()

	if os.Getenv("GIN_MODE") == "release" {
		middleware.Logger.Info("Setting up for release mode")
		gin.SetMode(gin.ReleaseMode)
	} else if os.Getenv("GIN_MODE") == "debug" {
		middleware.Logger.Info("Setting up for dev mode")
		gin.SetMode(gin.DebugMode)
		// if err := godotenv.Load(".env.dev"); err != nil {
		// 	log.Panic("failed to load .env.dev file: ", err)
		// }

	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic("failed to load configuration: ", err)
	}
	routes := server.NewServer(cfg)
	server.SetupRoutes(routes, cfg)
	routes.Run()
}
