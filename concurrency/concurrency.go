package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	rsCh := make(chan result)
	for _, url := range urls {
		go func(u string) {
			rsCh <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-rsCh
		res[r.string] = r.bool
	}

	return res
}
