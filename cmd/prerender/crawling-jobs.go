package prerender

func GenerateUrls(baseUrls []string, subPaths []string, pdps []string) []string {
	var urls []string

	for _, baseUrl := range baseUrls {
		for _, subPath := range subPaths {
			for _, pdp := range pdps {

				// Add trailing slash if not present base url
				if baseUrl[len(baseUrl)-1] != '/' {
					baseUrl += "/"
				}

				// remove starting slash if present
				if subPath[0] == '/' {
					subPath = subPath[1:]
				}

				// Add trailing slash if not present
				if subPath[len(subPath)-1] != '/' {
					subPath += "/"
				}

				urls = append(urls, baseUrl+subPath+pdp)
			}
		}
	}
	return urls
}
