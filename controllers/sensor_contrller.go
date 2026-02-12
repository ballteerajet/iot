package controllers

import (
	"net/http"

	"iot/config"
	"iot/models"

	"github.com/gin-gonic/gin"
)

// CreateSensorData godoc
// @Summary Create sensor data
// @Description Create new sensor data (user & admin)
// @Tags Sensor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.SensorData true "Sensor data"
// @Success 200 {object} models.SensorData
// @Failure 401 {object} ErrorResponse
// @Router /sensor [post]
func CreateSensorData(c *gin.Context) {
	var data models.SensorData
	c.BindJSON(&data)

	config.DB.Create(&data)
	c.JSON(http.StatusOK, data)
}

// GetAllSensorData godoc
// @Summary Get all sensor data
// @Description Get all sensor data
// @Tags Sensor
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.SensorData
// @Failure 401 {object} ErrorResponse
// @Router /sensor [get]
func GetAllSensorData(c *gin.Context) {
	var data []models.SensorData
	config.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

// GetSensorDataByID godoc
// @Summary Get sensor data by ID
// @Description Get sensor data by ID
// @Tags Sensor
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Sensor ID"
// @Success 200 {object} models.SensorData
// @Failure 401 {object} ErrorResponse
// @Router /sensor/{id} [get]
func GetSensorDataByID(c *gin.Context) {
	var data models.SensorData
	config.DB.First(&data, c.Param("id"))
	c.JSON(http.StatusOK, data)
}
