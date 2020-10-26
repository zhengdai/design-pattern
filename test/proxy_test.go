package test

import (
	"design-pattern/proxy"
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := proxy.NewNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("url: %s\n HttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("url: %s\n HttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("url: %s\n HttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "POST")
	fmt.Printf("url: %s\n HttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "GET")
	fmt.Printf("url: %s\n HttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
}
