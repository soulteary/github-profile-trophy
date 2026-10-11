package fetchers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type repositoryTransport func(*http.Request) (*http.Response, error)

func (f repositoryTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func repositoryClient(t *testing.T, responses []string, cursors []interface{}) *GitHubClient {
	t.Helper()
	t.Setenv("GITHUB_TOKEN1", "test-token")
	calls := 0
	t.Cleanup(func() {
		if calls != len(responses) {
			t.Errorf("requests = %d, want %d", calls, len(responses))
		}
	})
	return &GitHubClient{client: &http.Client{Transport: repositoryTransport(func(r *http.Request) (*http.Response, error) {
		if calls >= len(responses) {
			t.Fatal("unexpected extra page request")
		}
		var req GraphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}
		if req.Variables["username"] != "example" || req.Variables["after"] != cursors[calls] {
			t.Fatalf("unexpected variables: %#v", req.Variables)
		}
		if !strings.Contains(req.Query, "after: $after") || !strings.Contains(req.Query, "hasNextPage") {
			t.Fatal("query does not request cursor pagination")
		}
		if !strings.Contains(req.Query, "direction: ASC, field: CREATED_AT") || strings.Contains(req.Query, "field: STARGAZERS") {
			t.Fatal("repositories must use immutable creation ordering")
		}
		body := responses[calls]
		calls++
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})}}
}

func TestRequestUserRepositoryPagination(t *testing.T) {
	c := repositoryClient(t, []string{
		`{"data":{"user":{"repositories":{"totalCount":2,"pageInfo":{"hasNextPage":true,"endCursor":"page1"},"nodes":[{"stargazers":{"totalCount":10},"languages":{"nodes":[{"name":"Go"}]}}]}}}}`,
		`{"data":{"user":{"repositories":{"totalCount":2,"pageInfo":{"hasNextPage":false,"endCursor":"page2"},"nodes":[{"stargazers":{"totalCount":7},"languages":{"nodes":[{"name":"Rust"}]}}]}}}}`,
	}, []interface{}{nil, "page1"})
	repos, err := c.RequestUserRepository("example")
	if err != nil {
		t.Fatal(err)
	}
	info := NewUserInfo(GitHubUserActivity{}, GitHubUserIssue{}, GitHubUserPullRequest{}, repos)
	if len(repos.Repositories.Nodes) != 2 || info.TotalRepositories != 2 || info.TotalStargazers != 17 || info.LanguageCount != 2 {
		t.Fatalf("incorrect aggregate: %+v", info)
	}
}

func TestRequestUserRepositoryEmpty(t *testing.T) {
	c := repositoryClient(t, []string{`{"data":{"user":{"repositories":{"totalCount":0,"pageInfo":{"hasNextPage":false},"nodes":[]}}}}`}, []interface{}{nil})
	repos, err := c.RequestUserRepository("example")
	if err != nil || repos.Repositories.TotalCount != 0 || len(repos.Repositories.Nodes) != 0 {
		t.Fatalf("repos=%+v err=%v", repos, err)
	}
}

func TestRequestUserRepositoryRejectsIncompleteResults(t *testing.T) {
	for _, tc := range []struct {
		name      string
		responses []string
		cursors   []interface{}
	}{
		{"later page error", []string{`{"data":{"user":{"repositories":{"totalCount":2,"pageInfo":{"hasNextPage":true,"endCursor":"page1"},"nodes":[{"stargazers":{"totalCount":10}}]}}}}`, `{"errors":[{"message":"query failed"}]}`}, []interface{}{nil, "page1"}},
		{"missing cursor", []string{`{"data":{"user":{"repositories":{"pageInfo":{"hasNextPage":true},"nodes":[]}}}}`}, []interface{}{nil}},
		{"repeated cursor", []string{`{"data":{"user":{"repositories":{"pageInfo":{"hasNextPage":true,"endCursor":"same"},"nodes":[]}}}}`, `{"data":{"user":{"repositories":{"pageInfo":{"hasNextPage":true,"endCursor":"same"},"nodes":[]}}}}`}, []interface{}{nil, "same"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := repositoryClient(t, tc.responses, tc.cursors)
			repos, err := c.RequestUserRepository("example")
			if err == nil || len(repos.Repositories.Nodes) != 0 {
				t.Fatalf("expected error without partial results, got %+v, %v", repos, err)
			}
		})
	}
}

func TestRepositoryPaginationHasOneOverallDeadline(t *testing.T) {
	t.Setenv("GITHUB_TOKEN1", "test-token")
	calls := 0
	var deadline time.Time
	client := &GitHubClient{fetchTimeout: 25 * time.Millisecond, client: &http.Client{Transport: repositoryTransport(func(r *http.Request) (*http.Response, error) {
		got, ok := r.Context().Deadline()
		if !ok {
			t.Error("missing overall deadline")
		}
		if calls == 0 {
			deadline = got
		} else if !got.Equal(deadline) {
			t.Error("deadline restarted for next page")
		}
		calls++
		timer := time.NewTimer(10 * time.Millisecond)
		defer timer.Stop()
		select {
		case <-r.Context().Done():
			return nil, r.Context().Err()
		case <-timer.C:
		}
		body := fmt.Sprintf(`{"data":{"user":{"repositories":{"pageInfo":{"hasNextPage":true,"endCursor":"page%d"},"nodes":[{}]}}}}`, calls)
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})}}
	repos, err := client.RequestUserRepositoryContext(context.Background(), "example")
	if !errors.Is(err, context.DeadlineExceeded) || len(repos.Repositories.Nodes) != 0 {
		t.Fatalf("expected deadline without partial totals, got %+v, %v", repos, err)
	}
}

func TestUserInfoPropagatesCallerCancellation(t *testing.T) {
	t.Setenv("GITHUB_TOKEN1", "test-token")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{}, 4)
	stopped := make(chan struct{}, 4)
	client := &GitHubClient{client: &http.Client{Transport: repositoryTransport(func(r *http.Request) (*http.Response, error) {
		started <- struct{}{}
		<-r.Context().Done()
		stopped <- struct{}{}
		return nil, r.Context().Err()
	})}}
	result := make(chan error, 1)
	go func() { _, err := client.RequestUserInfoContext(ctx, "example"); result <- err }()
	for i := 0; i < 4; i++ {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("fetch did not start")
		}
	}
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected cancellation, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("caller cancellation did not stop the fetch")
	}
	if len(stopped) != 4 {
		t.Fatalf("only %d of 4 requests stopped", len(stopped))
	}
}
