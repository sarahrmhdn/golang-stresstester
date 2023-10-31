package library

import "testing"

func TestWP(t *testing.T) {
	s := &StressRequests{
		URL:  "https://fajarsultra.com",
		Mode: "wordpress",
	}
	s.Execute(0, 0)
}
