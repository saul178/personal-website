package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/handlers"
	"github.com/saul178/personal-website/services"
)

func SetupRoutes(r *gin.Engine) {
	githubServices := services.NewGithubService()
	redisServices := services.NewRedisCacheService()
	r.GET("/api/github/repos", handlers.GetOwnerReposHandler(githubServices, redisServices))
	r.GET("/api/github/commits", handlers.GetReposCommitsHandler(githubServices, redisServices))

	// home page metadata
	r.GET("api/resume", handlers.GetHomeDataHandler(redisServices))

	// serve the resume file to be downloaded
	r.GET("api/download-resume", handlers.GetDownloadResumeHandler())

	// TODO: serve music files to be streamed on personal page
	// r.GET("api/music", handlers.GetMusicDataHandler())
}
