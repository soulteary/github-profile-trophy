package fetchers

import "time"

// Language represents a programming language
type Language struct {
	Name string `json:"name"`
}

// Stargazers represents stargazer count
type Stargazers struct {
	TotalCount int `json:"totalCount"`
}

// Repository represents a GitHub repository
type Repository struct {
	Languages struct {
		Nodes []Language `json:"nodes"`
	} `json:"languages"`
	Stargazers Stargazers `json:"stargazers"`
	CreatedAt  string     `json:"createdAt"`
}

// GitHubUserRepository represents repository data from GitHub API
type GitHubUserRepository struct {
	Repositories struct {
		TotalCount int          `json:"totalCount"`
		Nodes      []Repository `json:"nodes"`
	} `json:"repositories"`
}

// GitHubUserIssue represents issue data from GitHub API
type GitHubUserIssue struct {
	OpenIssues struct {
		TotalCount int `json:"totalCount"`
	} `json:"openIssues"`
	ClosedIssues struct {
		TotalCount int `json:"totalCount"`
	} `json:"closedIssues"`
}

// GitHubUserPullRequest represents pull request data from GitHub API
type GitHubUserPullRequest struct {
	PullRequests struct {
		TotalCount int `json:"totalCount"`
	} `json:"pullRequests"`
}

// GitHubUserActivity represents user activity data from GitHub API
type GitHubUserActivity struct {
	CreatedAt               string `json:"createdAt"`
	ContributionsCollection struct {
		TotalCommitContributions            int `json:"totalCommitContributions"`
		RestrictedContributionsCount        int `json:"restrictedContributionsCount"`
		TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
	} `json:"contributionsCollection"`
	Organizations struct {
		TotalCount int `json:"totalCount"`
	} `json:"organizations"`
	Followers struct {
		TotalCount int `json:"totalCount"`
	} `json:"followers"`
}

// UserInfo aggregates all user statistics
type UserInfo struct {
	TotalCommits       int
	TotalFollowers     int
	TotalIssues        int
	TotalOrganizations int
	TotalPullRequests  int
	TotalReviews       int
	TotalStargazers    int
	TotalRepositories  int
	LanguageCount      int
	DurationYear       int
	DurationDays       int
	AncientAccount     int
	Joined2020         int
	OGAccount          int
}

// NewUserInfo creates a UserInfo from GitHub API responses
func NewUserInfo(
	userActivity GitHubUserActivity,
	userIssue GitHubUserIssue,
	userPullRequest GitHubUserPullRequest,
	userRepository GitHubUserRepository,
) UserInfo {
	// Calculate total commits
	totalCommits := userActivity.ContributionsCollection.RestrictedContributionsCount +
		userActivity.ContributionsCollection.TotalCommitContributions

	// Calculate total stargazers
	totalStargazers := 0
	for _, node := range userRepository.Repositories.Nodes {
		totalStargazers += node.Stargazers.TotalCount
	}

	// Calculate unique languages
	languages := make(map[string]bool)
	for _, node := range userRepository.Repositories.Nodes {
		for _, lang := range node.Languages.Nodes {
			if lang.Name != "" {
				languages[lang.Name] = true
			}
		}
	}
	languageCount := len(languages)

	// Find earliest repository creation date
	earliestRepoDate := userActivity.CreatedAt
	for _, node := range userRepository.Repositories.Nodes {
		if node.CreatedAt != "" {
			if earliestRepoDate == "" || node.CreatedAt < earliestRepoDate {
				earliestRepoDate = node.CreatedAt
			}
		}
	}

	// Calculate duration
	var durationYear, durationDays int
	var ancientAccount, joined2020, ogAccount int

	if earliestRepoDate != "" {
		earliestTime, err := time.Parse(time.RFC3339, earliestRepoDate)
		if err == nil {
			now := time.Now()
			duration := now.Sub(earliestTime)

			// Duration in years (approximate)
			durationYear = int(duration.Hours() / 24 / 365.25)

			// Duration in days / 100
			durationDays = int(duration.Hours() / 24 / 100)

			// Check special dates
			year := earliestTime.Year()
			if year <= 2010 {
				ancientAccount = 1
			}
			if year == 2020 {
				joined2020 = 1
			}
			if year <= 2008 {
				ogAccount = 1
			}
		}
	}

	return UserInfo{
		TotalCommits:       totalCommits,
		TotalFollowers:     userActivity.Followers.TotalCount,
		TotalIssues:        userIssue.OpenIssues.TotalCount + userIssue.ClosedIssues.TotalCount,
		TotalOrganizations: userActivity.Organizations.TotalCount,
		TotalPullRequests:  userPullRequest.PullRequests.TotalCount,
		TotalReviews:       userActivity.ContributionsCollection.TotalPullRequestReviewContributions,
		TotalStargazers:    totalStargazers,
		TotalRepositories:  userRepository.Repositories.TotalCount,
		LanguageCount:      languageCount,
		DurationYear:       durationYear,
		DurationDays:       durationDays,
		AncientAccount:     ancientAccount,
		Joined2020:         joined2020,
		OGAccount:          ogAccount,
	}
}
