package services

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/saul178/personal-website/middleware"
)

const (
	GithubUser string = "saul178"
)

func getPinnedRepos() []string {
	return []string{"personal-website", "manga-library-proj", "Personal-SampleShare"}
}

type GithubService struct {
	Client *github.Client
}

func NewGithubService(token string) *GithubService {
	if token == "" {
		middleware.Logger.Info("Proceeding without an unauthenticated Github client. Certain features might not work and requests are limited to 60 requests.")
		unauthClient := github.NewClient(nil)
		return &GithubService{Client: unauthClient}
	}

	middleware.Logger.Info("proceeding with an authenticated github user")
	client := github.NewClient(nil).WithAuthToken(token)
	return &GithubService{Client: client}
}

type RepoMetadata struct {
	Title     string         `json:"title,omitempty"`
	Desc      string         `json:"desc,omitempty"`
	Languages map[string]int `json:"languages,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
	URL       string         `json:"url,omitempty"`
}

// TODO: check for rate limit abuse and also give a context.Withvalue to these functions so that requests sleep till reset
// look here https://github.com/google/go-github for when doing this.
func getRepoLanguages(ctx context.Context, s *GithubService, owner string, repos string) (map[string]int, error) {
	languages := make(map[string]int)
	repository, _, err := s.Client.Repositories.ListLanguages(ctx, owner, repos)
	if err != nil {
		middleware.Logger.Error("Error fetching repo languages", "repo", repos, "err", err)
		return nil, err
	}
	languages = repository

	return languages, nil
}

func (s *GithubService) GetPinnedRepos(ctx context.Context) ([]RepoMetadata, error) {
	owner := GithubUser
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
				middleware.Logger.Error("Error fetching repo", "repo data", repoData, "err", err)
				return
			}
			if resp.Response.StatusCode != http.StatusOK {
				middleware.Logger.Error("GET request for repo", "repo data", repoData, "err", resp.Response.StatusCode)
				return
			}

			repoLanguages, err := getRepoLanguages(ctx, s, owner, repoData)
			if err != nil {
				middleware.Logger.Error("Error fetching repo languages", "repo data", repoData, "err", err)
			}

			metadata := RepoMetadata{
				Title:     repository.GetName(),
				Desc:      repository.GetDescription(),
				Languages: repoLanguages,
				CreatedAt: repository.GetCreatedAt().Format(time.DateOnly),
				UpdatedAt: repository.GetUpdatedAt().Format(time.DateOnly),
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
	CommitMsg []string `json:"commit_msg,omitempty"`
	Time      []string `json:"time,omitempty"`
}

func (s *GithubService) GetCommitsForRepo(ctx context.Context, repo string, limit int) (*RepoCommitMetadata, error) {
	owner := GithubUser
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: limit},
	}

	commits, resp, err := s.Client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		middleware.Logger.Error("GET request for repo commits failed", "repo", repo, "err", resp.Response.StatusCode)
		return nil, err
	}

	var commitMsg, commitAuth, commitTime []string
	for _, c := range commits {
		if c.GetCommit() != nil {
			commitMsg = append(commitMsg, c.GetCommit().GetMessage())
			commitAuth = append(commitAuth, c.GetCommit().GetAuthor().GetName())
			commitTime = append(commitTime, c.GetCommit().GetAuthor().GetDate().Format(time.DateOnly))
		}
	}

	metadata := &RepoCommitMetadata{
		Author:    commitAuth,
		CommitMsg: commitMsg,
		Time:      commitTime,
	}

	return metadata, nil
}
