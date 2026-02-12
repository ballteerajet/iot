package middleware

import (
	"net/http"

	"iot/config"
	"iot/models"
	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key required"})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.Where("api_key = ?", apiKey).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
