package crawling

import "net/url"

func GenerateUrls(baseUrls []string, subPaths []string, pdps []string) []string {
	var urls []string

	for _, baseUrl := range baseUrls {
		for _, subPath := range subPaths {
			for _, pdp := range pdps {

				u, err := url.JoinPath(baseUrl, subPath, pdp)

				if err != nil {
					panic(err)
				}

				urls = append(urls, u)
			}
		}
	}
	return urls
}
