package handlers

// TODO: will do this when ready, look into bandcamp api to make this easier
// import (
// 	"net/http"
// 	"os"
// 	"path"
//
// 	"github.com/gin-gonic/gin"
// )
//
// func GetMusicDataHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		filePath := path.Join("internal", "metadata", "music.json")
//
// 		_, err := os.Stat(filePath)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "music.json file not found"})
// 			return
// 		}
//
// 		c.Header("Content-Transfer-Encoding", "binary")
// 		c.Header("Content-Disposition", "attachment; filename="+filePath)
// 		c.Header("Content-Type", "audio/mpeg")
//
// 		c.File(filePath)
// 	}
// }
