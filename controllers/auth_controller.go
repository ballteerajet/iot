package controllers

import (
	"net/http"

	"iot/config"
	"iot/models"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin123"`
}

type LoginResponse struct {
	APIKey string `json:"api_key" example:"b3f1c9c1c4..."`
	Role   string `json:"role" example:"admin"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

// Login godoc
// @Summary Login
// @Description Login with username and password to get API Key
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login payload"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		return
	}

	var user models.User
	if err := config.DB.
		Where("username = ? AND password = ?", req.Username, req.Password).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		APIKey: user.APIKey,
		Role:   user.Role,
	})
}
