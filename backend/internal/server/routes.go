package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/config"
	"github.com/saul178/personal-website/handlers"
	"github.com/saul178/personal-website/services"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	githubServices := services.NewGithubService(cfg.GithubToken.Token)
	redisServices := services.NewRedisCacheService(cfg.Redis.Addr, cfg.Redis.Passwd)

	// github metadata
	r.GET("/api/github/repos", handlers.GetOwnerReposHandler(githubServices, redisServices))
	r.GET("/api/github/commits", handlers.GetReposCommitsHandler(githubServices, redisServices))

	// home page metadata
	r.GET("/api/resume", handlers.GetHomeDataHandler(redisServices))

	// serve the resume file to be downloaded
	r.GET("/api/download-resume", handlers.GetDownloadResumeHandler())

	// TODO: serve music files to be streamed on personal page through a json pointing to the files.
	// r.GET("api/music", handlers.GetMusicDataHandler())
}
