package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/middleware"
	"github.com/saul178/personal-website/services"
)

const (
	cacheResumeKey   = "resume:saul"
	resumeReqTimeout = time.Second * 5
)

// TODO: these models should be separated from the handlers
type Skills struct {
	OS         map[string]string `json:"os"`
	Languages  map[string]string `json:"languages"`
	Tools      map[string]string `json:"tools"`
	Frameworks map[string]string `json:"frameworks"`
	Databases  map[string]string `json:"databases"`
}

type Education struct {
	School        string `json:"school"`
	Degree        string `json:"degree"`
	DateCompleted string `json:"date_completed"`
}

type Metadata struct {
	Education []Education `json:"education"`
	Skills    Skills      `json:"skills"`
}

func GetHomeDataHandler(r *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), resumeReqTimeout)
		defer cancel()

		var cachedMetadata Metadata
		if err := r.Get(ctx, cacheResumeKey, &cachedMetadata); err == nil {
			c.JSON(http.StatusOK, cachedMetadata)
			return
		}

		filepath := path.Join("internal", "metadata", "data.json")
		data, err := os.ReadFile(filepath)
		if err != nil {
			middleware.ErrorLog.Printf("Error reading metadata: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read json file data"})
			return
		}

		var metadata Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			middleware.ErrorLog.Printf("Error parsing metadata: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse metadata"})
			return
		}

		if err := r.Set(ctx, cacheResumeKey, &metadata); err != nil {
			middleware.ErrorLog.Printf("failed to set resume data to redis cache. Is redis server running? %v", err)
		}

		c.JSON(http.StatusOK, metadata)
	}
}
