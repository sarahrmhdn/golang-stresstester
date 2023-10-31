package library

import (
	"net/http"
	"crypto/tls"
	"log"
)

func DefaultGETRequest(url string, threadNumber int, packageNumber int) {
	// Create an HTTP client
	client := &http.Client{}

	// If the protocol is https, skip SSL verification
	if len(url) > 5 && url[0:5] == "https" {
		// Warning: This skips SSL/TLS verification for the https request
		// It's insecure and should not be used in production environments
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}

	// Create and execute the GET request
	log.Printf("Sending package with thread %d and seq %d", threadNumber, packageNumber)
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Status code: %d\n", resp.StatusCode)
}