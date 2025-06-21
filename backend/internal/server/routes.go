package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/handlers"
	"github.com/saul178/personal-website/services"
)

func SetupRoutes(r *gin.Engine) {
	// routes will be setup here such as {"/", "projects", "personal", "github"}
	githubServices := services.NewGithubService()

	r.GET("/github-repos", func(ctx *gin.Context) {
		handlers.GetOwnerReposHandler(githubServices)
	})
}
