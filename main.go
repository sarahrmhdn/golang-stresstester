package main

import (
	"flag"
	"github.com/sarahrmhdn/golang-stresstester/library"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var (
		url     string
		threads int
		mode    string
		wg      sync.WaitGroup
		howlong int
	)

	// Define flags
	flag.StringVar(&url, "url", "https://example.com", "Specify target url (with http/https)")
	flag.IntVar(&threads, "threads", 10, "Specify number of threads to attack")
	flag.IntVar(&howlong, "howlong", 3, "How long attack will be running (in seconds). -1 means forever")
	flag.StringVar(&mode, "mode", "normal", "Specify mode ('normal', 'wordpress')")

	// Parse the flags
	flag.Parse()
	T := &library.StressRequests{
		URL:  url,
		Mode: mode,
	}
	// Use the values from flags

	// Simulate a main task that could take a long time
	if howlong != -1 {
		go func() {
			log.Printf("This will run maximum %d seconds\n", howlong)
			time.Sleep(time.Duration(howlong) * time.Second)
			os.Exit(0)
		}()
	} else {
		log.Println("This job will run forever")
	}

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
