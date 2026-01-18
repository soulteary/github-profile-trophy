package fetchers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/soulteary/github-profile-trophy/internal/common"
)

const (
	githubAPIURL = "https://api.github.com/graphql"
)

// GraphQLRequest represents a GraphQL request
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// queryUserActivity is the GraphQL query for user activity
const queryUserActivity = `
query userInfo($username: String!) {
  user(login: $username) {
    createdAt
    contributionsCollection {
      totalCommitContributions
      restrictedContributionsCount
      totalPullRequestReviewContributions
    }
    organizations(first: 1) {
      totalCount
    }
    followers(first: 1) {
      totalCount
    }
  }
}
`

// queryUserIssue is the GraphQL query for user issues
const queryUserIssue = `
query userInfo($username: String!) {
  user(login: $username) {
    openIssues: issues(states: OPEN) {
      totalCount
    }
    closedIssues: issues(states: CLOSED) {
      totalCount
    }
  }
}
`

// queryUserPullRequest is the GraphQL query for user pull requests
const queryUserPullRequest = `
query userInfo($username: String!) {
  user(login: $username) {
    pullRequests(first: 1) {
      totalCount
    }
  }
}
`

// queryUserRepository is the GraphQL query for user repositories
const queryUserRepository = `
query userInfo($username: String!) {
  user(login: $username) {
    repositories(first: 50, ownerAffiliations: OWNER, orderBy: {direction: DESC, field: STARGAZERS}) {
      totalCount
      nodes {
        languages(first: 3, orderBy: {direction:DESC, field: SIZE}) {
          nodes {
            name
          }
        }
        stargazers {
          totalCount
        }
        createdAt
      }
    }
  }
}
`

// GitHubClient handles GitHub API requests
type GitHubClient struct {
	client *http.Client
}

// NewGitHubClient creates a new GitHub API client
func NewGitHubClient() *GitHubClient {
	return &GitHubClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// executeQuery executes a GraphQL query
func (c *GitHubClient) executeQuery(query string, variables map[string]interface{}, token string) (*http.Response, error) {
	reqBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", githubAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RequestUserActivity fetches user activity data
func (c *GitHubClient) RequestUserActivity(username string) (GitHubUserActivity, error) {
	variables := map[string]interface{}{
		"username": username,
	}

	fetcher := func(vars interface{}, token string) (*http.Response, error) {
		v := vars.(map[string]interface{})
		return c.executeQuery(queryUserActivity, v, token)
	}

	resp, err := common.Retryer(fetcher, variables, 0, 0, time.Duration(common.DefaultGitHubRetryDelay)*time.Millisecond)
	if err != nil {
		return GitHubUserActivity{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubUserActivity{}, err
	}

	var graphqlResp struct {
		Data struct {
			User *GitHubUserActivity `json:"user"`
		} `json:"data"`
		Errors []common.GraphQLError `json:"errors"`
	}

	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return GitHubUserActivity{}, err
	}

	if len(graphqlResp.Errors) > 0 {
		return GitHubUserActivity{}, common.NewCustomError(
			graphqlResp.Errors[0].Message,
			common.ErrorTypeNotFound,
		)
	}

	if graphqlResp.Data.User == nil {
		return GitHubUserActivity{}, common.NewCustomError(
			"User not found",
			common.ErrorTypeNotFound,
		)
	}

	return *graphqlResp.Data.User, nil
}

// RequestUserIssue fetches user issue data
func (c *GitHubClient) RequestUserIssue(username string) (GitHubUserIssue, error) {
	variables := map[string]interface{}{
		"username": username,
	}

	fetcher := func(vars interface{}, token string) (*http.Response, error) {
		v := vars.(map[string]interface{})
		return c.executeQuery(queryUserIssue, v, token)
	}

	resp, err := common.Retryer(fetcher, variables, 0, 0, time.Duration(common.DefaultGitHubRetryDelay)*time.Millisecond)
	if err != nil {
		return GitHubUserIssue{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubUserIssue{}, err
	}

	var graphqlResp struct {
		Data struct {
			User *GitHubUserIssue `json:"user"`
		} `json:"data"`
		Errors []common.GraphQLError `json:"errors"`
	}

	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return GitHubUserIssue{}, err
	}

	if len(graphqlResp.Errors) > 0 {
		return GitHubUserIssue{}, common.NewCustomError(
			graphqlResp.Errors[0].Message,
			common.ErrorTypeNotFound,
		)
	}

	if graphqlResp.Data.User == nil {
		return GitHubUserIssue{}, common.NewCustomError(
			"User not found",
			common.ErrorTypeNotFound,
		)
	}

	return *graphqlResp.Data.User, nil
}

// RequestUserPullRequest fetches user pull request data
func (c *GitHubClient) RequestUserPullRequest(username string) (GitHubUserPullRequest, error) {
	variables := map[string]interface{}{
		"username": username,
	}

	fetcher := func(vars interface{}, token string) (*http.Response, error) {
		v := vars.(map[string]interface{})
		return c.executeQuery(queryUserPullRequest, v, token)
	}

	resp, err := common.Retryer(fetcher, variables, 0, 0, time.Duration(common.DefaultGitHubRetryDelay)*time.Millisecond)
	if err != nil {
		return GitHubUserPullRequest{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubUserPullRequest{}, err
	}

	var graphqlResp struct {
		Data struct {
			User *GitHubUserPullRequest `json:"user"`
		} `json:"data"`
		Errors []common.GraphQLError `json:"errors"`
	}

	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return GitHubUserPullRequest{}, err
	}

	if len(graphqlResp.Errors) > 0 {
		return GitHubUserPullRequest{}, common.NewCustomError(
			graphqlResp.Errors[0].Message,
			common.ErrorTypeNotFound,
		)
	}

	if graphqlResp.Data.User == nil {
		return GitHubUserPullRequest{}, common.NewCustomError(
			"User not found",
			common.ErrorTypeNotFound,
		)
	}

	return *graphqlResp.Data.User, nil
}

// RequestUserRepository fetches user repository data
func (c *GitHubClient) RequestUserRepository(username string) (GitHubUserRepository, error) {
	variables := map[string]interface{}{
		"username": username,
	}

	fetcher := func(vars interface{}, token string) (*http.Response, error) {
		v := vars.(map[string]interface{})
		return c.executeQuery(queryUserRepository, v, token)
	}

	resp, err := common.Retryer(fetcher, variables, 0, 0, time.Duration(common.DefaultGitHubRetryDelay)*time.Millisecond)
	if err != nil {
		return GitHubUserRepository{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubUserRepository{}, err
	}

	var graphqlResp struct {
		Data struct {
			User *GitHubUserRepository `json:"user"`
		} `json:"data"`
		Errors []common.GraphQLError `json:"errors"`
	}

	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return GitHubUserRepository{}, err
	}

	if len(graphqlResp.Errors) > 0 {
		return GitHubUserRepository{}, common.NewCustomError(
			graphqlResp.Errors[0].Message,
			common.ErrorTypeNotFound,
		)
	}

	if graphqlResp.Data.User == nil {
		return GitHubUserRepository{}, common.NewCustomError(
			"User not found",
			common.ErrorTypeNotFound,
		)
	}

	return *graphqlResp.Data.User, nil
}

// RequestUserInfo fetches all user information
func (c *GitHubClient) RequestUserInfo(username string) (UserInfo, error) {
	type result struct {
		activity    *GitHubUserActivity
		issue       *GitHubUserIssue
		pullRequest *GitHubUserPullRequest
		repository  *GitHubUserRepository
		err         error
		which       string // "activity", "issue", "pullRequest", "repository"
	}

	results := make(chan result, 4)

	go func() {
		activity, err := c.RequestUserActivity(username)
		if err != nil {
			results <- result{err: err, which: "activity"}
			return
		}
		results <- result{activity: &activity, which: "activity"}
	}()

	go func() {
		issue, err := c.RequestUserIssue(username)
		if err != nil {
			results <- result{err: err, which: "issue"}
			return
		}
		results <- result{issue: &issue, which: "issue"}
	}()

	go func() {
		pullRequest, err := c.RequestUserPullRequest(username)
		if err != nil {
			results <- result{err: err, which: "pullRequest"}
			return
		}
		results <- result{pullRequest: &pullRequest, which: "pullRequest"}
	}()

	go func() {
		repository, err := c.RequestUserRepository(username)
		if err != nil {
			results <- result{err: err, which: "repository"}
			return
		}
		results <- result{repository: &repository, which: "repository"}
	}()

	var activity GitHubUserActivity
	var issue GitHubUserIssue
	var pullRequest GitHubUserPullRequest
	var repository GitHubUserRepository
	var hasActivity bool
	var firstErr error

	for i := 0; i < 4; i++ {
		res := <-results
		if res.err != nil {
			if firstErr == nil {
				firstErr = res.err
			}
			continue
		}
		switch res.which {
		case "activity":
			if res.activity != nil {
				activity = *res.activity
				hasActivity = true
			}
		case "issue":
			if res.issue != nil {
				issue = *res.issue
			}
		case "pullRequest":
			if res.pullRequest != nil {
				pullRequest = *res.pullRequest
			}
		case "repository":
			if res.repository != nil {
				repository = *res.repository
			}
		}
	}

	// If we got an error and no essential data (activity is required), return the error
	if firstErr != nil && !hasActivity {
		return UserInfo{}, firstErr
	}

	// Activity is required, if we don't have it, return error
	if !hasActivity {
		if firstErr != nil {
			return UserInfo{}, firstErr
		}
		return UserInfo{}, common.NewCustomError("User not found", common.ErrorTypeNotFound)
	}

	return NewUserInfo(activity, issue, pullRequest, repository), nil
}
