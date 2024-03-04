package crawling

import (
	"context"
	"net/http"
	"time"
)

func CrawlingTest(ctx context.Context, url string, userAgent string, timeOut int) *CrawlResult {

	ttlCtx, _ := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)

	c, cancel := context.WithCancel(ttlCtx)

	defer cancel()

	req, err := http.NewRequestWithContext(c, "GET", url, nil)

	if err != nil {
		return &CrawlResult{
			Url:        url,
			StatusCode: 0,
			Success:    false,
			Error:      err,
		}
	}

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return &CrawlResult{
			Url:        url,
			StatusCode: resp.StatusCode,
			Success:    false,
			Error:      err,
		}
	}

	return &CrawlResult{
		Url:        url,
		StatusCode: resp.StatusCode,
		Success:    true,
		Error:      nil,
	}
}
