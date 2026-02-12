package controllers

import (
	"net/http"
	"strconv"

	"iot/config"
	"iot/models"

	"github.com/gin-gonic/gin"
)

// CreateUsersBulk godoc
// @Summary Create users (bulk)
// @Description Admin create users in bulk
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body []models.User true "Users"
// @Success 200 {array} models.User
// @Failure 403 {object} ErrorResponse
// @Router /users/bulk [post]
func CreateUsersBulk(c *gin.Context) {
	var users []models.User
	c.BindJSON(&users)

	for i := range users {
		users[i].APIKey = models.GenerateAPIKey()
	}

	config.DB.Create(&users)
	c.JSON(http.StatusOK, users)
}

// GetUsers godoc
// @Summary Get all users
// @Description Admin get all users
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} ErrorResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Admin get user by ID
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 403 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	config.DB.First(&user, id)
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Admin update user
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body models.User true "User"
// @Success 200 {object} models.User
// @Failure 403 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	config.DB.First(&user, id)

	c.BindJSON(&user)
	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Admin delete user
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 403 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.User{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
