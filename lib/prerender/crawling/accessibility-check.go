package crawling

import (
	"net/http"
)

func CrawlingTest(url string, userAgent string, statusCode int) (bool, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, err
	}

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != statusCode {
		return false, nil
	}

	return true, nil
}
