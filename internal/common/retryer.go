package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// GraphQLResponse represents a GraphQL API response
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// FetcherFunction represents a function that fetches data
type FetcherFunction func(variables interface{}, token string) (*http.Response, error)

// GetPATCount returns the number of GitHub API tokens available
func GetPATCount() int {
	count := 0
	for i := 1; i <= 10; i++ {
		token := os.Getenv(fmt.Sprintf("GITHUB_TOKEN%d", i))
		if token != "" {
			count++
		}
	}
	return count
}

// GetPAT returns a PAT token by index (1-based)
func GetPAT(index int) string {
	return os.Getenv(fmt.Sprintf("GITHUB_TOKEN%d", index))
}

// Retryer tries to execute the fetcher function until it succeeds or max retries is reached
func Retryer(fetcher FetcherFunction, variables interface{}, retries int, maxRetries int, delay time.Duration) (*http.Response, error) {
	if maxRetries == 0 {
		maxRetries = GetPATCount()
	}
	if maxRetries == 0 {
		return nil, NewCustomError("No GitHub API tokens found", ErrorTypeNoTokens)
	}

	if retries >= maxRetries {
		return nil, NewCustomError(
			"Downtime due to GitHub API rate limiting",
			ErrorTypeMaxRetry,
		)
	}

	token := GetPAT(retries + 1)
	if token == "" {
		return nil, NewCustomError("No GitHub API tokens found", ErrorTypeNoTokens)
	}

	response, err := fetcher(variables, token)
	if err != nil {
		// Network/unexpected error → retry with delay
		if retries < maxRetries-1 {
			time.Sleep(delay)
			return Retryer(fetcher, variables, retries+1, maxRetries, delay)
		}
		return nil, err
	}

	// Check for rate limiting in GraphQL errors
	body, err := io.ReadAll(response.Body)
	if err != nil {
		response.Body.Close()
		return nil, err
	}
	response.Body.Close()

	var graphqlResp GraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err == nil {
		if len(graphqlResp.Errors) > 0 {
			errorType := graphqlResp.Errors[0].Type
			errorMsg := graphqlResp.Errors[0].Message
			isRateLimited := errorType == "RATE_LIMITED" || regexp.MustCompile(`(?i)rate limit`).MatchString(errorMsg)
			isNotFound := strings.Contains(strings.ToLower(errorMsg), "could not resolve to a user") ||
				strings.Contains(strings.ToLower(errorMsg), "not found")

			if isNotFound {
				// User not found, don't retry
				return nil, NewCustomError("User not found", ErrorTypeNotFound)
			}

			if isRateLimited {
				// Retry with next token
				time.Sleep(delay)
				return Retryer(fetcher, variables, retries+1, maxRetries, delay)
			}
		}
	}

	// Check for bad credentials or account suspended in REST API errors
	if response.StatusCode == 401 || response.StatusCode == 403 {
		bodyStr := string(body)
		isBadCredential := strings.Contains(bodyStr, "Bad credentials")
		isAccountSuspended := strings.Contains(bodyStr, "Sorry. Your account was suspended.")

		if isBadCredential || isAccountSuspended {
			// Retry with next token
			time.Sleep(delay)
			return Retryer(fetcher, variables, retries+1, maxRetries, delay)
		}
	}

	// Check for 404 (user not found)
	if response.StatusCode == 404 {
		return nil, NewCustomError("User not found", ErrorTypeNotFound)
	}

	// Create a new response reader for the body
	response.Body = io.NopCloser(strings.NewReader(string(body)))
	return response, nil
}
