package github

type ContributionDay struct {
	ContributionCount int    `json:"contributionCount"`
	Date              string `json:"date"`
	Color             string `json:"color"`
}

type ContributionWeek struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
	FirstDay         string            `json:"firstDay"`
}

type ContributionCalendar struct {
	TotalContributions int                `json:"totalContributions"`
	Weeks              []ContributionWeek `json:"weeks"`
}

type ContributionGraph struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}
