package middleware

import (
	"github.com/gin-gonic/gin"
	// "github.com/rs/cors"
	"net/http"
)

// Setting up CORS middleware
func SetUpCORS() gin.HandlerFunc{
	return func (c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
            return
		}
		c.Next()
	}
}