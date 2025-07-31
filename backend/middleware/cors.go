package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if domain == "" {
			log.Fatal("DOMAIN env not set")
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", domain) // this will have to be restricted when in prod replace w/ my domain soon
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
