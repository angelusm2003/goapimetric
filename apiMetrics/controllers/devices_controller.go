package controllers

import (
	"apiMetrics/models"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//var db *gorm.DB

func RegisterDevice(c *gin.Context, db *gorm.DB) {
	fmt.Println("RegisterDevice function called") // Add logging statement
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set creation date
	device.DateCreation = time.Now()

	if err := db.Create(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, device)
}
