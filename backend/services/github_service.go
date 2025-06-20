package services

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v72/github"
	"github.com/joho/godotenv"
)

const (
	githubUser string = "saul178"
)

func getPinnedRepos() []string {
	return []string{"personal-website", "manga-library-proj", "Personal-SampleShare"}
}

type RepoMetadata struct {
	Title     string   `json:"title,omitempty"`
	Desc      string   `json:"desc,omitempty"`
	Languages string   `json:"languages,omitempty"`
	Commits   []string `json:"commits,omitempty"`
	URL       string   `json:"url,omitempty"`
}

type GithubService struct {
	Client *github.Client
}

func newGithubService() *GithubService {
	client := initNewGithubClient()
	return &GithubService{Client: client}
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

func initNewGithubClient() *github.Client {
	ghToken, err := getGithubToken()
	if err != nil {
		log.Printf("Warning: %v.\n Proceeding without an authenticated Github client.\nCertain features might not work and requests are limited.", err)

		unauthClient := github.NewClient(nil)
		return unauthClient
	}

	client := github.NewClient(nil).WithAuthToken(ghToken)
	return client
}

func (s *GithubService) GetPinnedRepos(ctx context.Context) ([]RepoMetadata, error) {
	owner := githubUser
	repos := getPinnedRepos()

	var pinnedReposData []RepoMetadata
	for i := range repos {
		repository, resp, err := s.Client.Repositories.Get(ctx, owner, repos[i])
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
		return nil, errors.New("failed to fetch repositories successfully.")
	}
	return pinnedReposData, nil
}

func (s *GithubService) GetRepoCommits(ctx context.Context) ([]RepoMetadata, error) {
}
