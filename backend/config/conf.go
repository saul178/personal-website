package config

import "os"

type Config struct {
	Domain      DomainConfig
	GithubToken GithubConfig
	Redis       RedisConfig
}

type DomainConfig struct {
	DomainName string
}

type GithubConfig struct {
	Token string
}

type RedisConfig struct {
	Addr   string
	Passwd string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Domain: DomainConfig{
			os.Getenv("DOMAIN"),
		},
		GithubToken: GithubConfig{
			Token: os.Getenv("GITHUB_TOKEN"),
		},
		Redis: RedisConfig{
			Addr:   os.Getenv("REDIS_ADDR"),
			Passwd: os.Getenv("REDIS_PASSWD"),
		},
	}
	return cfg, nil
}
