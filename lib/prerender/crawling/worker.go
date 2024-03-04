package crawling

import (
	"context"
	"fmt"
	"sync"
)

func Crawl(urls []string, userAgent string, timeOut int) ([]*CrawlResult, error) {

	ch := make(chan *CrawlResult)

	go func() {
		wg := sync.WaitGroup{}
		for _, url := range urls {
			wg.Add(1)
			fmt.Println("Crawling: ", url)
			go func(u string) {
				res := CrawlingTest(context.Background(), u, userAgent, timeOut)
				ch <- res
			}(url)
		}
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}

	return nil, nil
}
