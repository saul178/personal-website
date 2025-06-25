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
	githubTimeout  = 5 * time.Second
	commitMsgLimit = 3
)

// TODO: cache these requests so that it doesnt request all the time maybe redis?
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

		repoName := c.Query("repo")
		// NOTE: we can assume repoName wont be empty sinces we are hard coding the pinned repos anyways
		commitData, err := s.GetCommitsForRepo(ctx, repoName, commitMsgLimit)
		if err != nil {
			log.Printf("Error fetching commits for %s: %v", repoName, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch commit data for repo"})
			return
		}

		c.JSON(http.StatusOK, commitData)
	}
}
