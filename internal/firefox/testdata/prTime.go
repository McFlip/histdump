package main

// Create firefox test data for testing Firefox history extraction
// Reference: https://firefox-source-docs.mozilla.org/nspr/reference/prtime.html

import (
	"fmt"
	"time"
)

const (
	marchThird         = "2025-03-03 09:00:00"
	febrarySecond      = "2025-02-02 09:00:00"
	januaryTwentyThird = "2025-01-23 09:00:00"
)

func main() {
	var dates = []string{marchThird, febrarySecond, januaryTwentyThird}
	for _, date := range dates {
		// Parse the date string into a time.Time object
		t, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		// Convert to Chrome timestamp
		firefoxTimestamp := t.Unix() * 1_000_000

		fmt.Printf("Firefox timestamp for %s is %d\n", date, firefoxTimestamp)
	}
}
