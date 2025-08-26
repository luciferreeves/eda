package github

import (
	"context"
	"eda/config"

	"github.com/google/go-github/v74/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var GithubClient = github.NewClient(nil).WithAuthToken(config.GitHub.GitHubToken)

var GraphQLClient = githubv4.NewClient(oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: config.GitHub.GitHubToken},
)))
