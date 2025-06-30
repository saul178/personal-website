package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/saul178/personal-website/middleware"
	"github.com/saul178/personal-website/services"
)

const (
	githubTimeout  = 5 * time.Second
	commitMsgLimit = 3
	cacheKey       = "repos:" + services.GithubUser
)

// TODO: cache these requests so that it doesnt request all the time maybe redis?
func GetOwnerReposHandler(g *services.GithubService, r *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		// check redis cache
		var cachedRepos []services.RepoMetadata
		err := r.Get(ctx, cacheKey, &cachedRepos)
		if err == nil {
			log.Println("CACHE HIT: fetching repos from redis cache..")
			c.JSON(http.StatusOK, cachedRepos)
			return
		}

		if err != redis.Nil {
			// Log real Redis error
			log.Printf("Redis error for key %s: %v\n", cacheKey, err)
		} else {
			log.Println("Cache MISS: key not found")
		}

		repoData, err := g.GetPinnedRepos(ctx)
		if err != nil || len(repoData) == 0 {
			middleware.ErrorLog.Printf("No repos fetched or error occurred: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch repository data"})
			return
		}

		// set the results to the redis cache
		log.Println("setting repo data to redis cache..")
		_ = r.Set(ctx, cacheKey, repoData)

		c.JSON(http.StatusOK, repoData)
	}
}

func GetReposCommitsHandler(g *services.GithubService, r *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		repoName := c.Query("repo")
		cacheKey := fmt.Sprintf("commits: %s", repoName)

		// check redis cache
		var cachedRepoCommits services.RepoCommitMetadata
		err := r.Get(ctx, cacheKey, &cachedRepoCommits)
		if err == nil {
			log.Println("Cache HIT: fetched commits from redis")
			c.JSON(http.StatusOK, cachedRepoCommits)
			return
		}

		if err != redis.Nil {
			// Log real Redis error
			log.Printf("Redis error for key %s: %v\n", cacheKey, err)
		} else {
			log.Println("Cache MISS: key not found")
		}

		// NOTE: we can assume repoName wont be empty sinces we are hard coding the pinned repos anyways
		commitData, err := g.GetCommitsForRepo(ctx, repoName, commitMsgLimit)
		if err != nil {
			middleware.ErrorLog.Printf("Error fetching commits for %s: %v", repoName, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch commit data for repo"})
			return
		}

		log.Println("setting commits to redis cache..")
		_ = r.Set(ctx, cacheKey, commitData)

		c.JSON(http.StatusOK, commitData)
	}
}
