package http

import "github.com/gin-gonic/gin"

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-KEY")

		if key != "supersecret" {
			c.JSON(401, gin.H{"error": "UNAUTHORIZED"})
			c.Abort()
			return
		}

		c.Next()
	}
}
