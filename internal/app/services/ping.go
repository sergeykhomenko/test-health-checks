package services

import (
	"context"
	"errors"
	"github.com/sergeykhomenko/test-health-checks/internal/types"
	"log/slog"
	"net/http"
	"time"
)

var requestTimeout int

func SetPingbackTimeout(timeout int) {
	requestTimeout = timeout
}

func PingUrl(ctx context.Context, url string, resultCh chan types.PingReport) {
	status := getStatusForUrl(url)

	for {
		select {
		case <-ctx.Done():
			slog.Debug("Skipped ping for " + url + " due to context cancel")
			return
		default:
			resultCh <- types.PingReport{
				Url:    url,
				Status: status,
			}
			return
		}
	}
}

func getStatusForUrl(url string) types.PingStatus {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(requestTimeout))

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.Warn("Cant create a request with timeout for " + url)
	}

	res, err := http.DefaultClient.Do(request)
	if errors.Is(err, context.DeadlineExceeded) {
		return types.PingStatusTimeout
	}

	if err != nil {
		slog.Info("Error when making request for "+url, slog.String("error", err.Error()))
		return types.PingStatusError
	}

	if res.StatusCode < 200 {
		return types.PingStatusError
	}

	if res.StatusCode < 300 {
		return types.PingStatusActive
	}

	if res.StatusCode < 400 {
		return types.PingStatusRedirect
	}

	return types.PingStatusInactive
}
