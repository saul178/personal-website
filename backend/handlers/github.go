package handlers

import (
	"context"
	"fmt"
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
	repoCacheKey   = "repos:" + services.GithubUser
)

func GetOwnerReposHandler(g *services.GithubService, r *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		// check redis cache
		var cachedRepos []services.RepoMetadata
		err := r.Get(ctx, repoCacheKey, &cachedRepos)
		if err == nil {
			middleware.InfoLog.Println("CACHE HIT: fetching repos from redis cache..")
			c.JSON(http.StatusOK, cachedRepos)
			return
		}

		if err != redis.Nil {
			middleware.ErrorLog.Printf("Redis error for key %s: %v\n", repoCacheKey, err)
		}

		repoData, err := g.GetPinnedRepos(ctx)
		if err != nil || len(repoData) == 0 {
			middleware.ErrorLog.Printf("No repos fetched or error occurred: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch repository data"})
			return
		}

		// set the results to the redis cache
		middleware.InfoLog.Println("setting repo data to redis cache..")
		if err := r.Set(ctx, repoCacheKey, repoData); err != nil {
			middleware.ErrorLog.Printf("failed to set repo data to redis cache. Is redis server running? %v", err)
		}

		c.JSON(http.StatusOK, repoData)
	}
}

func GetReposCommitsHandler(g *services.GithubService, r *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), githubTimeout)
		defer cancel()

		repoName := c.Query("repo")
		cacheKey := fmt.Sprintf("commits:%s", repoName)

		// check redis cache
		var cachedRepoCommits services.RepoCommitMetadata
		err := r.Get(ctx, cacheKey, &cachedRepoCommits)
		if err == nil {
			middleware.InfoLog.Println("Cache HIT: fetched commits from redis")
			c.JSON(http.StatusOK, cachedRepoCommits)
			return
		}

		if err != redis.Nil {
			middleware.ErrorLog.Printf("Redis error for key %s: %v\n", cacheKey, err)
		}

		// NOTE: we can assume repoName wont be empty sinces we are hard coding the pinned repos anyways
		commitData, err := g.GetCommitsForRepo(ctx, repoName, commitMsgLimit)
		if err != nil {
			middleware.ErrorLog.Printf("Error fetching commits for %s: %v", repoName, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch commit data for repo"})
			return
		}

		middleware.InfoLog.Println("CACHE MISS: setting commits to redis cache..")
		if err := r.Set(ctx, cacheKey, commitData); err != nil {
			middleware.ErrorLog.Printf("failed to set commits to redis cache. Is redis server running? %v", err)
		}

		c.JSON(http.StatusOK, commitData)
	}
}
