package services

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/joho/godotenv"
)

const (
	githubUser string = "saul178"
)

func getPinnedRepos() []string {
	return []string{"personal-website", "manga-library-proj", "Personal-SampleShare"}
}

type GithubService struct {
	Client *github.Client
}

func NewGithubService() *GithubService {
	client := initNewGithubClient()
	return &GithubService{Client: client}
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

type RepoMetadata struct {
	Title     string         `json:"title,omitempty"`
	Desc      string         `json:"desc,omitempty"`
	Languages map[string]int `json:"languages,omitempty"`
	URL       string         `json:"url,omitempty"`
}

func getRepoLanguages(ctx context.Context, s *GithubService, owner string, repos string) (map[string]int, error) {
	languages := make(map[string]int)
	repository, _, err := s.Client.Repositories.ListLanguages(ctx, owner, repos)
	if err != nil {
		log.Printf("Error fetching repo (%s) languages: %v", repos, err)
		return nil, err
	}
	languages = repository

	return languages, nil
}

func (s *GithubService) GetPinnedRepos(ctx context.Context) ([]RepoMetadata, error) {
	owner := githubUser
	repos := getPinnedRepos()
	var wg sync.WaitGroup
	var mu sync.Mutex

	pinnedReposData := make([]RepoMetadata, len(repos))
	for i, repoData := range repos {
		wg.Add(1)
		go func(i int, repo string) {
			defer wg.Done()

			repository, resp, err := s.Client.Repositories.Get(ctx, owner, repoData)
			if err != nil {
				log.Printf("Error fetching repo %s: %v", repoData, err)
				return
			}
			if resp.Response.StatusCode != http.StatusOK {
				log.Printf("GET request for repo %s responded with %v", repoData, resp.Response.StatusCode)
				return
			}

			repoLanguages, err := getRepoLanguages(ctx, s, owner, repoData)
			if err != nil {
				log.Printf("Error fetching repo (%s) languages: %v", repoData, err)
			}

			metadata := RepoMetadata{
				Title:     repository.GetName(),
				Desc:      repository.GetDescription(),
				Languages: repoLanguages,
				URL:       repository.GetHTMLURL(),
			}
			mu.Lock()
			pinnedReposData[i] = metadata
			mu.Unlock()
		}(i, repoData)
	}
	wg.Wait()

	if len(pinnedReposData) == 0 {
		return nil, errors.New("failed to fetch repositories successfully. Possible API/Network issues.")
	}
	return pinnedReposData, nil
}

type RepoCommitMetadata struct {
	Author    []string `json:"author,omitempty"`
	CommitMsg []string `json:"commitMsg,omitempty"`
	Time      []string `json:"time,omitempty"`
}

func (s *GithubService) GetRepoCommits(ctx context.Context, limit int) ([]RepoCommitMetadata, error) {
	owner := githubUser
	repos := getPinnedRepos()
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: limit},
	}

	var commitMetadata []RepoCommitMetadata
	for i := range repos {
		commits, resp, err := s.Client.Repositories.ListCommits(ctx, owner, repos[i], opts)
		if err != nil {
			log.Printf("Error fetching repo commits %s: %v", repos[i], err)
			continue
		}
		if resp.Response.StatusCode != http.StatusOK {
			log.Printf("GET request for repo commits %s responded with %v", repos[i], resp.Response.StatusCode)
			continue
		}

		var commitMsg []string
		var commitAuth []string
		var commitTime []string
		for _, c := range commits {
			if c.GetCommit() != nil {
				commitMsg = append(commitMsg, c.GetCommit().GetMessage())
				commitAuth = append(commitAuth, c.GetCommit().GetAuthor().GetName())
				commitTime = append(commitTime, c.GetCommit().GetAuthor().GetDate().Format(time.DateTime))

			}
		}
		metadata := RepoCommitMetadata{
			Author:    commitAuth,
			CommitMsg: commitMsg,
			Time:      commitTime,
		}
		commitMetadata = append(commitMetadata, metadata)
	}

	if len(commitMetadata) == 0 {
		return nil, errors.New("failed to fetch commits for all repos")
	}
	return commitMetadata, nil
}
