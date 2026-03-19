package util

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

// ExportToCSV writes the provided rows to a CSV file at the specified output path.
// Each row is a slice of strings, and the first row is treated as the header.
func ExportToCSV(rows [][]string, outputFile string) {
	// Open the output CSV file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"URL", "Title", "Visit Time", "Visit Count", "Last Visit Time"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("failed to write header: %v", err)
	}

	// Write each row
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			log.Fatalf("failed to write record: %v", err)
		}
	}
}

// ParseDateRange parses the afterDate and beforeDate strings into time.Time objects.
// If the strings are empty, it returns zero values for those times.
func ParseDateRange(afterDate, beforeDate string) (afterTime, beforeTime time.Time) {
	var afterDateTime, beforeDateTime time.Time

	if afterDate != "" {
		parsedTime, err := time.Parse("2006-01-02", afterDate)
		if err != nil {
			log.Fatalf("Invalid after date format: %v", err)
		}
		afterDateTime = parsedTime
	}
	if beforeDate != "" {
		parsedTime, err := time.Parse("2006-01-02", beforeDate)
		if err != nil {
			log.Fatalf("Invalid before date format: %v", err)
		}
		beforeDateTime = parsedTime
	}

	return afterDateTime, beforeDateTime
}

// FilterDates checks if the visitTime is within the specified date range.
// Returns true for valid dates, false otherwise.
func FilterDates(visitTime, afterDate, beforeDate time.Time) bool {
	if !afterDate.IsZero() {
		if visitTime.Before(afterDate) {
			return false
		}
	}
	if !beforeDate.IsZero() {
		if visitTime.After(beforeDate) {
			return false
		}
	}
	return true
}
