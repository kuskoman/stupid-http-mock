package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := getEnvOrDefault("PORT", ":8080")
	http.HandleFunc("/", handleRequest)
	log.Printf("Mock server listening on %s", port)
	http.ListenAndServe(port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var requestProperties []string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error when reading request body: %+v", err) // not using panic because app should not exit after wrong request
	}

	headers := stringifyHeaders(r.Header)

	requestProperties = append(requestProperties, "Request received")
	requestProperties = append(requestProperties, fmt.Sprintf("URL: %s", r.URL))
	requestProperties = append(requestProperties, fmt.Sprintf("Method: %s", r.Method))
	requestProperties = append(requestProperties, fmt.Sprintf("Host: %s", r.Host))
	requestProperties = append(requestProperties, fmt.Sprintf("Body: %s", string(body)))
	requestProperties = append(requestProperties, fmt.Sprintf("Headers:\n%s", headers))

	if r.Method == "POST" {
		r.ParseForm()

		requestProperties = append(requestProperties, fmt.Sprintf("Form data:\n%s", r.Form.Encode()))
	}

	stringifiedRequestProperties := strings.Join(requestProperties, "\n")
	log.Println(stringifiedRequestProperties)

	w.Write([]byte(stringifiedRequestProperties))
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func stringifyHeaders(headers http.Header) string {
	var stringifiedHeaders string

	for name, value := range headers {
		for _, header := range value {
			stringifiedHeaders += fmt.Sprintf("  %s: %s\n", name, header)
		}
	}

	return stringifiedHeaders
}
