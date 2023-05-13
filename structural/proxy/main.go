package main

import "fmt"

func main() {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	displayString := "\nUrl: %s\nHttpCode: %d\nBody: %s\n"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf(displayString, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf(displayString, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf(displayString, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf(displayString, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf(displayString, appStatusURL, httpCode, body)
}
