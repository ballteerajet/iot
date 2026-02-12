package routes

import (
	"iot/controllers"
	"iot/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth.Use(middleware.APIKeyAuth())

	// Sensor (user ใช้ได้)
	auth.POST("/sensor", controllers.CreateSensorData)
	auth.GET("/sensor", controllers.GetAllSensorData)
	auth.GET("/sensor/:id", controllers.GetSensorDataByID)

	// Admin only
	admin := auth.Group("/users")
	admin.Use(middleware.AdminOnly())

	admin.POST("/bulk", controllers.CreateUsersBulk)
	admin.GET("/", controllers.GetUsers)
	admin.GET("/:id", controllers.GetUserByID)
	admin.PUT("/:id", controllers.UpdateUser)
	admin.DELETE("/:id", controllers.DeleteUser)
}
