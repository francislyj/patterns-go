package proxy

type nginx struct {
	application application
	maxAllowedRequest int
	rateLimiter map[string]int
}

func newNginx(maxAllowedRequest int) *nginx {
	return &nginx{
		application: application{},
		maxAllowedRequest: maxAllowedRequest,
		rateLimiter: make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {

	if !n.checkRateLimit(url) {
		return 403, "not allowed"
	}

	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimit(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}