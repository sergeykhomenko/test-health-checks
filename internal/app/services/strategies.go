package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sergeykhomenko/test-health-checks/internal/types"
	"net/http"
	"sync"
)

func FirstToFall(urls []string) types.PingResponse {
	var (
		resultsCh   = make(chan types.PingReport)
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
		firstReport types.PingReport
	)

	defer cancel()

	for i := range urls {
		url := urls[i]
		wg.Add(1)
		go func() {
			PingUrl(ctx, url, resultsCh)
			wg.Done()
		}()
	}

	go func() {
		firstReport = <-resultsCh
		cancel()
	}()

	wg.Wait()

	results := types.PingResponse{
		Content: gin.H{firstReport.Url: firstReport.Status},
		Status:  http.StatusOK,
	}

	for _, url := range urls {
		if _, ok := results.Content[url]; ok {
			continue
		}

		results.Content[url] = types.PingStatusTerminated
	}

	if firstReport.Status != types.PingStatusActive {
		results.Status = http.StatusNoContent
	}

	return results
}

func AtLeastOne(urls []string) types.PingResponse {
	var (
		resultsCh = make(chan types.PingReport, len(urls))
		results   = types.PingResponse{
			Content: gin.H{},
			Status:  http.StatusNoContent,
		}
		wg sync.WaitGroup
	)

	for i := range urls {
		url := urls[i]
		wg.Add(1)
		go func() {
			PingUrl(context.Background(), url, resultsCh)
			wg.Done()
		}()
	}

	wg.Wait()

	for range urls {
		resp := <-resultsCh
		results.Content[resp.Url] = resp.Status
		if resp.Status == types.PingStatusActive {
			results.Status = http.StatusOK
		}
	}

	return results
}
