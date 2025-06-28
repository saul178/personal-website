package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	m "github.com/saul178/personal-website/middleware"
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

// TODO: Look into caching so that it doesnt always read the file and do all these operations again
func GetHomeDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filepath := path.Join("internal", "metadata", "data.json")
		data, err := os.ReadFile(filepath)
		if err != nil {
			m.ErrorLog.Printf("Error reading metadata: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read json file data"})
			return
		}

		var metadata Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			m.ErrorLog.Printf("Error parsing metadata: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse metadata"})
			return
		}

		c.JSON(http.StatusOK, metadata)
	}
}
