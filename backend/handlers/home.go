package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	m "github.com/saul178/personal-website/middleware"
)

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
	DateCompleted string `json:"dateCompleted"`
}

type Metadata struct {
	Education []*Education `json:"education"`
	Skills    *Skills      `json:"skills"`
}

func fetchHomeMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		filepath := path.Join("internal", "metadata", "data.json")
		data, err := os.ReadFile(filepath)
		if err != nil {
			m.ErrorLog.Printf("Failed to read data from json file: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data from metadata.json"})
		}

		var metadata *Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse json file"})
		}

		c.JSON(http.StatusOK, &metadata)
	}
}
