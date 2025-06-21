// gin http handlers here
package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/services"
)

const (
	githubTimeout = 5 * time.Second
)

func GetOwnerReposHandler(s *services.GithubService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		repoData, err := s.GetPinnedRepos(ctx)
		if err != nil || len(repoData) == 0 {
			log.Printf("No repos fetched or error occurred: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch repository data"})
			return
		}

		c.JSON(http.StatusOK, repoData)
	}
}

func GetReposCommitsHandler(s *services.GithubService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		commitData, err := s.GetRepoCommits(ctx, 3)
		if err != nil || len(commitData) == 0 {
			log.Printf("No repos fetched or error occurred: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch repository data"})
			return
		}

		c.JSON(http.StatusOK, commitData)
	}
}
