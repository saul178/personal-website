package handlers

import (
	"errors"
	"log"
	"os"

	"github.com/google/go-github/v72/github"
	"github.com/joho/godotenv"
)

func getGithubToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("failed to load .env file: Github api will not function correctly without it.")
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
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
