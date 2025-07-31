package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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
		err := r.Get(ctx, cacheResumeKey, &cachedMetadata)
		if err == nil {
			c.JSON(http.StatusOK, cachedMetadata)
			return
		}

		if err != redis.Nil {
			middleware.Logger.Error("Redis error", "key", repoCacheKey, "err", err)
		}

		filepath := path.Join("internal", "metadata", "data.json")
		data, err := os.ReadFile(filepath)
		if err != nil {
			middleware.Logger.Error("Error reading metadata", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read json file data"})
			return
		}

		var metadata Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			middleware.Logger.Error("Error parsing json metadata", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse metadata"})
			return
		}

		if err := r.Set(ctx, cacheResumeKey, &metadata); err != nil {
			middleware.Logger.Error(
				"Failed to cache resume metadata to Redis Server, is Redis server running?",
				"key", cacheResumeKey,
				"err", err,
			)
		}
		c.JSON(http.StatusOK, metadata)
	}
}
