package main

import (
	"auth_service/config"
	"auth_service/database"
	"auth_service/routes"
	"auth_service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}

	db, err := database.InitDB(cfg.DBConnection)
	if err != nil {
		panic("Failed to connect database")
	}

	authService := services.NewAuthService(db, cfg.JWTSecret)

	r := gin.Default()
	routes.AuthRoutes(r, authService)

	r.Run(":" + cfg.Port)
}
