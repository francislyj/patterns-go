package proxy

type application struct {
	
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}

	if url == "/user/create" && method == "POST" {
		return 200, "ok"
	}

	return 404, "NOT FOUND"
}
