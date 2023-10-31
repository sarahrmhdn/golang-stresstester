package main

import (
	"flag"
	"github.com/sarahrmhdn/golang-stresstester/library"
	"log"
	"sync"
)

func main() {
	var (
		url     string
		threads int
		mode    string
		wg      sync.WaitGroup
	)

	// Define flags
	flag.StringVar(&url, "url", "https://example.com", "Specify target url (with http/https)")
	flag.IntVar(&threads, "threads", 10, "Specify number of threads to attack")
	flag.StringVar(&mode, "mode", "normal", "Specify mode ('normal', 'wordpress')")

	// Parse the flags
	flag.Parse()
	T := &library.StressRequests{
		URL:  url,
		Mode: mode,
	}
	// Use the values from flags
	log.Printf("Request received with URL %s and threads %d and mode %s\n", url, threads, mode)
	wg.Add(threads)
	for i := 1; i <= threads; i++ {
		go func(id int) {
			defer wg.Done()
			x := 0
			for {
				x++
				T.Execute(id, x)
			}
		}(i)
	}
	wg.Wait()
}
