package middlewares

import (
	"apiMetrics/utils"
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "token not provided"})
			c.Abort()
			return
		}

		// Parse the token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Set role in the context for downstream handlers
		c.Set("role", claims.Role)
		c.Next()
	}
}
