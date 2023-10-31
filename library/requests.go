package library

import (
	"crypto/tls"
	"log"
	"net/http"
)

type StressRequests struct {
	URL  string
	Mode string
}

func (s *StressRequests) Execute(threadNumber int, packageNumber int) {
	switch s.Mode {
	case "wordpress":
		log.Println("WP")
	case "normal":
		s.defaultGETRequest(s.URL, threadNumber, packageNumber)
	default:
		log.Println("Not supported")
	}
}

func (s *StressRequests) defaultGETRequest(url string, threadNumber int, packageNumber int) {
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
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("package with thread %d and seq %d: %v", threadNumber, packageNumber, err)
	} else {
		log.Printf("package with thread %d and seq %d: %d\n", threadNumber, packageNumber, resp.StatusCode)
	}
}
