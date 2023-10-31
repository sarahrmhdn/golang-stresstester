package main

import (
	"flag"
	"log"
	"github.com/sarahrmhdn/golang-stresstester/library"
)

func main() {
	var (
		url     string
		threads int
		mode    string
	)

	// Define flags
	flag.StringVar(&url, "url", "https://example.com", "Specify target url (with http/https)")
	flag.IntVar(&threads, "threads", 1000, "Specify number of threads to attack")
	flag.StringVar(&mode, "mode", "default", "Specify mode ('default', 'wordpress')")

	// Parse the flags
	flag.Parse()

	// Use the values from flags
	log.Printf("Request received with URL %s and threads %d and mode %s\n", url, threads, mode)
	library.DefaultGETRequest(url, 0, 0)
}
