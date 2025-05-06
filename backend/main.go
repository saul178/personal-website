package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin backend will be here

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H)
	})
}
