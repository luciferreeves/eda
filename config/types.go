package config

type ServerConfig struct {
	Host      string `env:"SERVER_HOST" default:"0.0.0.0"`
	Port      int    `env:"SERVER_PORT" default:"8080"`
	IsDevMode bool   `env:"SERVER_IS_DEV_MODE" default:"true"`
}

type GitHubConfig struct {
	GitHubUsername string `env:"GITHUB_USERNAME"`
	GitHubToken    string `env:"GITHUB_TOKEN"`
}
