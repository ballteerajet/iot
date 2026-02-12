// @title Sensor API
// @version 1.0
// @description Role-based API with API Key Auth
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
package main

import (
	"iot/config"
	"iot/models"
	"iot/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "iot/docs"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	// migrate
	config.DB.AutoMigrate(
		&models.User{},
		&models.SensorData{},
	)

	// init admin
	models.InitAdmin()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
