package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/handlers"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/resume", func(c *gin.Context) {
		data, err := os.ReadFile("../metadata/data.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read json data file"})
			return
		}

		var metadata handlers.Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse json data file"})
			return
		}

		c.JSON(http.StatusOK, metadata)
	})

	r.GET("/api/download-resume", func(c *gin.Context) {
		c.File("../assets/files/resume-v2.pdf")
	})

	return r
}

func TestMetadataEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/resume", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var response handlers.Metadata
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Education)
	assert.NotEmpty(t, response.Skills)
	assert.NotEmpty(t, response.Skills.Databases)
	assert.NotEmpty(t, response.Skills.Frameworks)
	assert.NotEmpty(t, response.Skills.Languages)
	assert.NotEmpty(t, response.Skills.OS)
	assert.NotEmpty(t, response.Skills.Tools)
}

func TestServingResumePDF(t *testing.T) {
	router := setupRouter()
	req, err := http.NewRequest("GET", "/api/download-resume", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/pdf")
	assert.Greater(t, w.Body.Len(), 0, "If pdf file exist, there should be more then 0 bytes")
}
