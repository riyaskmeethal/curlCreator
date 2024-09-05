package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	accessToken := ""

	importURL := ""

	bytJson := []byte{}

	req, err := http.NewRequest("POST", importURL, bytes.NewBuffer(bytJson))
	if err != nil {
		log.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", " application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	curlCommand := generateCurlCommand(req)
	fmt.Println("Curl Command:")
	fmt.Println(curlCommand)

}

func generateCurlCommand(req *http.Request) string {
	var curlCommand string

	// Start with the curl command and URL
	curlCommand = fmt.Sprintf("curl -X %s '%s'", req.Method, req.URL.String())

	// Add headers
	for name, values := range req.Header {
		for _, value := range values {
			curlCommand += fmt.Sprintf(" -H '%s: %s'", name, value)
		}
	}

	// Add data if it's a POST or PUT request
	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		body, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the body for the actual request
		curlCommand += fmt.Sprintf(" -d '%s'", string(body))
	}

	return curlCommand
}
