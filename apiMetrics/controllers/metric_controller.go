package controllers

import (
	"apiMetrics/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SaveMetrics(c *gin.Context, db *gorm.DB) {
	deviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	var metric models.Monitor_detail
	if err := c.BindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set device ID
	metric.DeviceID = uint(deviceID)

	if err := db.Create(&metric).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, metric)
}

func GetLatestMetrics(c *gin.Context, db *gorm.DB) {
	deviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	var metric models.Monitor_detail
	if err := db.Where("device_id = ?", deviceID).Order("timestamp DESC").First(&metric).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metric)
}

func GetAllMetricsHistory(c *gin.Context, db *gorm.DB) {
	deviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	var metrics []models.Monitor_detail
	if err := db.Where("device_id = ?", deviceID).Order("timestamp DESC").Find(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}
