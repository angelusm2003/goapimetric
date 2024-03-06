package routes

import (
	"apiMetrics/controllers"
	"apiMetrics/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Pass db to controller functions

	protected := router.Group("/")
	protected.Use(middlewares.IsAuthorized()) // Apply isauthorized to all routes in this group
	{
		protected.POST("/devices/register", func(c *gin.Context) {
			controllers.RegisterDevice(c, db)
		})
	}

	//router.POST("/devices/register", func(c *gin.Context) {
	//	controllers.RegisterDevice(c, db)
	//})
	router.POST("/devices/:id/metrics", func(c *gin.Context) {
		controllers.SaveMetrics(c, db)
	})
	router.GET("/devices/:id/metrics", func(c *gin.Context) {
		controllers.GetLatestMetrics(c, db)
	})
	router.GET("/devices/:id/metrics/history", func(c *gin.Context) {
		controllers.GetAllMetricsHistory(c, db)
	})

	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	router.POST("/signup", func(c *gin.Context) {
		controllers.Signup(c, db)
	})
	router.GET("/logout", func(c *gin.Context) {
		controllers.Logout(c)
	})
}
