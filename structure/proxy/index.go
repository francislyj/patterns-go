package proxy

import "fmt"

func TestProxy() {
	nginx := newNginx(3)

	statusUrl := "/app/status"
	createUrl := "/user/create"


	// handle status url
	code, message := nginx.handleRequest(statusUrl, "GET")
	fmt.Printf("\n handle request url %s \n result code \n %d message %s \n", statusUrl, code, message)

	code, message = nginx.handleRequest(statusUrl, "GET")
	fmt.Printf("\n handle request url %s \n result code \n %d message %s \n", statusUrl, code, message)

	code, message = nginx.handleRequest(statusUrl, "GET")
	fmt.Printf("\n handle request url %s \n result code \n %d message %s \n", statusUrl, code, message)

	code, message = nginx.handleRequest(statusUrl, "GET")
	fmt.Printf("\n handle request url %s \n result code \n %d message %s \n", statusUrl, code, message)


	// handle create url
	code, message = nginx.handleRequest(createUrl, "GET")
	fmt.Printf("\n handle request url %s \n result code \n %d message %s \n", createUrl, code, message)


}
