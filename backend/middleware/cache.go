package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() >= http.StatusBadRequest {
			c.Writer.Header().Set("Cache-Control", "no-store")
		} else {
			c.Writer.Header().Set("Cache-Control", "public, max-age=600")
		}
	}
}
