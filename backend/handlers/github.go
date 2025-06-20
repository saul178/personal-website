package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v72/github"
	"github.com/joho/godotenv"
)

type RepoMetadata struct {
	Title     string   `json:"title,omitempty"`
	Desc      string   `json:"desc,omitempty"`
	Languages string   `json:"languages,omitempty"`
	Commits   []string `json:"commits,omitempty"`
	URL       string   `json:"url,omitempty"`
}

const (
	githubUser    string = "saul178"
	githubTimeout        = 5 * time.Second
)

func getPinnedRepos() []string {
	return []string{"personal-website", "manga-library-proj", "Personal-SampleShare"}
}

func getGithubToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("failed to load .env file: Github api will not function correctly without it.")
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		return "", errors.New("GITHUB_TOKEN is empty: Github api will not function correctly without it.")
	}
	return ghToken, nil
}

func initGithubClient() *github.Client {
	ghToken, err := getGithubToken()
	if err != nil {
		log.Printf("Warning: %v.\n Proceeding without an authenticated Github client.\nCertain features might not work and requests are limited.", err)

		unauthClient := github.NewClient(nil)
		return unauthClient
	}

	client := github.NewClient(nil).WithAuthToken(ghToken)
	return client
}

func GetOwnerReposHandler(c *gin.Context) {
	client := initGithubClient()

	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	owner := githubUser
	repos := getPinnedRepos()

	var pinnedReposData []RepoMetadata
	for i := range repos {
		repository, resp, err := client.Repositories.Get(ctx, owner, repos[i])
		if err != nil {
			log.Printf("Error fetching repo %s: %v", repos[i], err)
			continue
		}
		if resp.Response.StatusCode != http.StatusOK {
			log.Printf("GET request for repo %s responded with %v", repos[i], resp.Response.StatusCode)
			continue
		}

		metadata := RepoMetadata{
			Title:     repository.GetName(),
			Desc:      repository.GetDescription(),
			Languages: repository.GetLanguage(),
			URL:       repository.GetHTMLURL(),
		}
		pinnedReposData = append(pinnedReposData, metadata)
	}

	if len(pinnedReposData) == 0 {
		log.Println("No repo data was fetched successfully")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all repository data"})
		return
	}
	c.JSON(http.StatusOK, pinnedReposData)
}

func GetReposCommitsHandler(c *gin.Context) {
}
