package common

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestRetryDelayStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Setenv("GITHUB_TOKEN1", "test-token")
	t.Setenv("GITHUB_TOKEN2", "test-token-2")
	calls := 0
	_, err := RetryerContext(ctx, func(interface{}, string) (*http.Response, error) {
		calls++
		cancel()
		return nil, errors.New("network failed")
	}, nil, 0, 2, time.Hour)
	if !errors.Is(err, context.Canceled) || calls != 1 {
		t.Fatalf("err=%v calls=%d", err, calls)
	}
}
