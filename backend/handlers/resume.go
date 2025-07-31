package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/middleware"
)

func GetDownloadResumeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPDFVersion := 2
		filePath := path.Join("internal", "assets", "files", fmt.Sprintf("resume-v%d.pdf", currentPDFVersion))

		_, err := os.Stat(filePath)
		if err != nil {
			middleware.Logger.Error("file not found", "error", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "pdf file not found"})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Type", "application/pdf")

		c.FileAttachment(filePath, "resume.pdf")
	}
}
