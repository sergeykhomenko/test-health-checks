package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/sergeykhomenko/test-health-checks/internal/types"
	"net/http"
	"time"
)

func PingUrl(ctx context.Context, url string, resultCh chan types.PingReport) {
	status := getStatusForUrl(url)

	for {
		select {
		case <-ctx.Done():
			// todo: here can be debug information about cancelling
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*600)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Cant create a request with timeout for", url)
	}

	res, err := http.DefaultClient.Do(request)
	if errors.Is(err, context.DeadlineExceeded) {
		return types.PingStatusTimeout
	}

	if err != nil {
		fmt.Println("Error when making request for", url, err.Error())
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
