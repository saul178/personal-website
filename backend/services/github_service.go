package services

import (
	"context"
	"errors"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/joho/godotenv"
	m "github.com/saul178/personal-website/middleware"
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
		m.WarningLog.Printf("Warning: %v.\nProceeding without an authenticated Github client.\nCertain features might not work and requests are limited.", err)
		unauthClient := github.NewClient(nil)
		return unauthClient
	}

	client := github.NewClient(nil).WithAuthToken(ghToken)
	return client
}

// TODO: clean up the errors so that it's more readable for all the functions that use errors.new
// TODO: look into caching so that requests arent being made every refresh: redis vs manually doing it
func getGithubToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("failed to load .env file: Github api will not function correctly without it.")
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	// TODO: this check feels useless, decide to keep or get rid
	if ghToken == "" {
		return "", errors.New("GITHUB_TOKEN is empty: Github api will not function correctly without it.")
	}
	return ghToken, nil
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
		m.ErrorLog.Printf("Error fetching repo (%s) languages: %v", repos, err)
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
				m.ErrorLog.Printf("Error fetching repo (%s): %v", repoData, err)
				return
			}
			if resp.Response.StatusCode != http.StatusOK {
				m.ErrorLog.Printf("GET request for repo (%s) responded with %v", repoData, resp.Response.StatusCode)
				return
			}

			repoLanguages, err := getRepoLanguages(ctx, s, owner, repoData)
			if err != nil {
				m.ErrorLog.Printf("Error fetching repo (%s) languages: %v", repoData, err)
			}

			// TODO: look into better time formats
			metadata := RepoMetadata{
				Title:     repository.GetName(),
				Desc:      repository.GetDescription(),
				Languages: repoLanguages,
				CreatedAt: repository.GetCreatedAt().Format(time.DateTime),
				UpdatedAt: repository.GetUpdatedAt().Format(time.DateTime),
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

func (s *GithubService) GetRepoCommits(ctx context.Context, limit int) (map[string]RepoCommitMetadata, error) {
	owner := githubUser
	repos := getPinnedRepos()
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: limit},
	}

	commitMetadata := make(map[string]RepoCommitMetadata)
	for _, repoName := range repos {
		commits, resp, err := s.Client.Repositories.ListCommits(ctx, owner, repoName, opts)
		if err != nil {
			m.ErrorLog.Printf("Error fetching repo commits (%s): %v", repoName, err)
			continue
		}
		if resp.Response.StatusCode != http.StatusOK {
			m.ErrorLog.Printf("GET request for repo commits (%s) responded with %v", repoName, resp.Response.StatusCode)
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

			} else {
				m.ErrorLog.Println("GET commits returned nil, failed to append results.")
				continue
			}
		}
		commitMetadata[repoName] = RepoCommitMetadata{
			Author:    commitAuth,
			CommitMsg: commitMsg,
			Time:      commitTime,
		}
	}

	if len(commitMetadata) == 0 {
		return nil, errors.New("failed to fetch commits for all repos")
	}
	return commitMetadata, nil
}

func (s *GithubService) GetCommitsForRepo(ctx context.Context, repo string, limit int) (*RepoCommitMetadata, error) {
	owner := githubUser
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: limit},
	}

	commits, resp, err := s.Client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		m.ErrorLog.Printf("GET request for repo commits (%s) responded with %v", repo, resp.Response.StatusCode)
		return nil, err
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

	metadata := &RepoCommitMetadata{
		Author:    commitAuth,
		CommitMsg: commitMsg,
		Time:      commitTime,
	}

	return metadata, nil
}
