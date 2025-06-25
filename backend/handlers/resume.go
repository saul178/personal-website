package handlers

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func GetDownloadResumeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := path.Join("internal", "assets", "files", "resume.pdf")

		_, err := os.Stat(filePath)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "pdf file not found"})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filePath)
		c.Header("Content-Type", "application/pdf")

		c.FileAttachment(filePath, "resume")
	}
}
