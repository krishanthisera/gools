package crawling

type CrawlResult struct {
	Url        string
	StatusCode int
	Success    bool
	Error      error
}
