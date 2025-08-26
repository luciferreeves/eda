package github

import (
	"context"
	"eda/config"

	"github.com/google/go-github/v74/github"
	"github.com/shurcooL/githubv4"
)

func GetProfile() *github.User {
	user, _, err := GithubClient.Users.Get(context.TODO(), config.GitHub.GitHubUsername)
	if err != nil {
		return nil
	}
	return user
}

func GetContributionGraph() (*ContributionGraph, error) {
	var query struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `graphql:"totalContributions"`
					Weeks              []struct {
						ContributionDays []struct {
							ContributionCount int    `graphql:"contributionCount"`
							Date              string `graphql:"date"`
							Color             string `graphql:"color"`
						} `graphql:"contributionDays"`
						FirstDay string `graphql:"firstDay"`
					} `graphql:"weeks"`
				} `graphql:"contributionCalendar"`
			} `graphql:"contributionsCollection"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String(config.GitHub.GitHubUsername),
	}

	err := GraphQLClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}

	contributionGraph := &ContributionGraph{
		ContributionCalendar: ContributionCalendar{
			TotalContributions: query.User.ContributionsCollection.ContributionCalendar.TotalContributions,
			Weeks:              make([]ContributionWeek, len(query.User.ContributionsCollection.ContributionCalendar.Weeks)),
		},
	}

	for i, week := range query.User.ContributionsCollection.ContributionCalendar.Weeks {
		contributionWeek := ContributionWeek{
			FirstDay:         week.FirstDay,
			ContributionDays: make([]ContributionDay, len(week.ContributionDays)),
		}

		for j, day := range week.ContributionDays {
			contributionWeek.ContributionDays[j] = ContributionDay{
				ContributionCount: day.ContributionCount,
				Date:              day.Date,
				Color:             day.Color,
			}
		}

		contributionGraph.ContributionCalendar.Weeks[i] = contributionWeek
	}

	return contributionGraph, nil
}
