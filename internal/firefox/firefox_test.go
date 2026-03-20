package firefox

import (
	"database/sql"
	"encoding/csv"
	"os"
	"testing"
	"time"

	"github.com/McFlip/histdump/histdump/internal/firefox/sqlc_ff"
)

func TestExtractFirefoxHistory(t *testing.T) {
	// Test data
	dbFile := "testdata/places.sqlite"
	outputFile := "testoutput/firefox_history.csv"
	afterDate, beforeDate := "", "" // No date filtering for this test

	// Remove the output file if it exists
	_ = os.Remove(outputFile)

	ExtractFirefoxHistory(dbFile, outputFile, afterDate, beforeDate)

	// Check if the output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("Output file %s was not created", outputFile)
	}

	// Read the output csv file
	file, err := os.Open(outputFile)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()
	// Read the CSV file
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	if len(records) == 0 {
		t.Fatalf("No records found in the CSV file")
	}

	header := records[0]

	if header[0] != "URL" || header[1] != "Title" || header[2] != "Visit Time" || header[3] != "Visit Count" || header[4] != "Last Visit Time" {
		t.Fatalf("CSV header does not match expected values")
	}

	const expectedURL = "https://time.gov/"
	if records[1][0] != expectedURL {
		t.Errorf("Expected URL %s, but got %s", expectedURL, records[1][0])
	}
	const expectedTitle = "National Institute of Standards and Technology | NIST"
	if records[1][1] != expectedTitle {
		t.Errorf("Expected Title %s, but got %s", expectedTitle, records[1][1])
	}
	const expectedVisitTime = "2025-07-01 16:50:06"
	if records[1][2] != expectedVisitTime {
		t.Errorf("Expected Visit Time %s, but got %s", expectedVisitTime, records[1][2])
	}
	const expectedVisitCount = "1"
	if records[1][3] != expectedVisitCount {
		t.Errorf("Expected Visit Count %s, but got %s", expectedVisitCount, records[1][3])
	}
	const expectedLastVisitTime = "2025-07-01 16:50:06"
	if records[1][4] != expectedLastVisitTime {
		t.Errorf("Expected Last Visit Time %s, but got %s", expectedLastVisitTime, records[1][4])
	}
}

var rows = []sqlc_ff.GetVisitsRow{
	{
		Url:           "https://time.gov/",
		Title:         "National Institute of Standards and Technology | NIST",
		VisitDate:     sql.NullInt64{Int64: 1740992400000000, Valid: true}, // "2025-03-03 09:00:00 UTC"
		VisitCount:    sql.NullInt64{Int64: 1, Valid: true},
		LastVisitDate: sql.NullInt64{Int64: 1740992400000000, Valid: true}, // "2025-03-03 09:00:00 UTC"
	},
	{
		Url:           "https://example.com/",
		Title:         "Example Domain",
		VisitDate:     sql.NullInt64{Int64: 1738486800000000, Valid: true}, // "2025-02-02 09:00:00 UTC"
		VisitCount:    sql.NullInt64{Int64: 1, Valid: true},
		LastVisitDate: sql.NullInt64{Int64: 1738486800000000, Valid: true}, // "2025-02-02 09:00:00 UTC"
	},
	{
		Url:           "https://xkcd.com/",
		Title:         "A webcomic of romance, sarcasm, math, and language.",
		VisitDate:     sql.NullInt64{Int64: 1737622800000000, Valid: true}, // "2025-01-23 09:00:00 UTC"
		VisitCount:    sql.NullInt64{Int64: 1, Valid: true},
		LastVisitDate: sql.NullInt64{Int64: 1737622800000000, Valid: true}, // "2025-01-23 09:00:00 UTC"
	},
}

func TestExtractFirefoxeHistoryWithAfter3Results(t *testing.T) {
	var beforeDate time.Time // Not used in this test
	afterDate, err := time.Parse("2006-01-02", "2025-01-01")
	if err != nil {
		t.Fatalf("Failed to parse after date: %v", err)
	}
	var expectedRows = [][]string{
		{"https://time.gov/", "National Institute of Standards and Technology | NIST", "2025-03-03 09:00:00", "1", "2025-03-03 09:00:00"},
		{"https://example.com/", "Example Domain", "2025-02-02 09:00:00", "1", "2025-02-02 09:00:00"},
		{"https://xkcd.com/", "A webcomic of romance, sarcasm, math, and language.", "2025-01-23 09:00:00", "1", "2025-01-23 09:00:00"},
	}

	actualRows := formatAndFilter(rows, afterDate, beforeDate)
	if len(actualRows) != len(expectedRows) {
		t.Errorf("Actual rows %v", actualRows)
		t.Fatalf("Expected %d rows, but got %d", len(expectedRows), len(actualRows))
	}

	for i, row := range actualRows {
		if row[0] != expectedRows[i][0] {
			t.Errorf("Expected URL %s, but got %s", expectedRows[i][0], row[0])
		}
		if row[1] != expectedRows[i][1] {
			t.Errorf("Expected Title %s, but got %s", expectedRows[i][1], row[1])
		}
		if row[2] != expectedRows[i][2] {
			t.Errorf("Expected Visit Time %s, but got %s", expectedRows[i][2], row[2])
		}
		if row[3] != expectedRows[i][3] {
			t.Errorf("Expected Visit Count %s, but got %s", expectedRows[i][3], row[3])
		}
		if row[4] != expectedRows[i][4] {
			t.Errorf("Expected Last Visit Time %s, but got %s", expectedRows[i][4], row[4])
		}
	}
}

func TestExtractFirefoxeHistoryWithAfter2Results(t *testing.T) {
	var beforeDate time.Time // Not used in this test
	afterDate, err := time.Parse("2006-01-02", "2025-02-01")
	if err != nil {
		t.Fatalf("Failed to parse after date: %v", err)
	}
	var expectedRows = [][]string{
		{"https://time.gov/", "National Institute of Standards and Technology | NIST", "2025-03-03 09:00:00", "1", "2025-03-03 09:00:00"},
		{"https://example.com/", "Example Domain", "2025-02-02 09:00:00", "1", "2025-02-02 09:00:00"},
	}

	actualRows := formatAndFilter(rows, afterDate, beforeDate)
	if len(actualRows) != len(expectedRows) {
		t.Errorf("Actual rows %v", actualRows)
		t.Fatalf("Expected %d rows, but got %d", len(expectedRows), len(actualRows))
	}

	for i, row := range actualRows {
		if row[0] != expectedRows[i][0] {
			t.Errorf("Expected URL %s, but got %s", expectedRows[i][0], row[0])
		}
		if row[1] != expectedRows[i][1] {
			t.Errorf("Expected Title %s, but got %s", expectedRows[i][1], row[1])
		}
		if row[2] != expectedRows[i][2] {
			t.Errorf("Expected Visit Time %s, but got %s", expectedRows[i][2], row[2])
		}
		if row[3] != expectedRows[i][3] {
			t.Errorf("Expected Visit Count %s, but got %s", expectedRows[i][3], row[3])
		}
		if row[4] != expectedRows[i][4] {
			t.Errorf("Expected Last Visit Time %s, but got %s", expectedRows[i][4], row[4])
		}
	}
}

func TestExtractFirefoxeHistoryWithBefore3Results(t *testing.T) {
	var afterDate time.Time // Not used in this test
	beforeDate, err := time.Parse("2006-01-02", "2026-01-01")
	if err != nil {
		t.Fatalf("Failed to parse before date: %v", err)
	}
	var expectedRows = [][]string{
		{"https://time.gov/", "National Institute of Standards and Technology | NIST", "2025-03-03 09:00:00", "1", "2025-03-03 09:00:00"},
		{"https://example.com/", "Example Domain", "2025-02-02 09:00:00", "1", "2025-02-02 09:00:00"},
		{"https://xkcd.com/", "A webcomic of romance, sarcasm, math, and language.", "2025-01-23 09:00:00", "1", "2025-01-23 09:00:00"},
	}

	actualRows := formatAndFilter(rows, afterDate, beforeDate)
	if len(actualRows) != len(expectedRows) {
		t.Errorf("Actual rows %v", actualRows)
		t.Fatalf("Expected %d rows, but got %d", len(expectedRows), len(actualRows))
	}

	for i, row := range actualRows {
		if row[0] != expectedRows[i][0] {
			t.Errorf("Expected URL %s, but got %s", expectedRows[i][0], row[0])
		}
		if row[1] != expectedRows[i][1] {
			t.Errorf("Expected Title %s, but got %s", expectedRows[i][1], row[1])
		}
		if row[2] != expectedRows[i][2] {
			t.Errorf("Expected Visit Time %s, but got %s", expectedRows[i][2], row[2])
		}
		if row[3] != expectedRows[i][3] {
			t.Errorf("Expected Visit Count %s, but got %s", expectedRows[i][3], row[3])
		}
		if row[4] != expectedRows[i][4] {
			t.Errorf("Expected Last Visit Time %s, but got %s", expectedRows[i][4], row[4])
		}
	}
}

func TestExtractFirefoxeHistoryWithBefore2Results(t *testing.T) {
	var afterDate time.Time // Not used in this test
	beforeDate, err := time.Parse("2006-01-02", "2025-03-01")
	if err != nil {
		t.Fatalf("Failed to parse before date: %v", err)
	}
	var expectedRows = [][]string{
		{"https://example.com/", "Example Domain", "2025-02-02 09:00:00", "1", "2025-02-02 09:00:00"},
		{"https://xkcd.com/", "A webcomic of romance, sarcasm, math, and language.", "2025-01-23 09:00:00", "1", "2025-01-23 09:00:00"},
	}

	actualRows := formatAndFilter(rows, afterDate, beforeDate)
	if len(actualRows) != len(expectedRows) {
		t.Errorf("Actual rows %v", actualRows)
		t.Fatalf("Expected %d rows, but got %d", len(expectedRows), len(actualRows))
	}

	for i, row := range actualRows {
		if row[0] != expectedRows[i][0] {
			t.Errorf("Expected URL %s, but got %s", expectedRows[i][0], row[0])
		}
		if row[1] != expectedRows[i][1] {
			t.Errorf("Expected Title %s, but got %s", expectedRows[i][1], row[1])
		}
		if row[2] != expectedRows[i][2] {
			t.Errorf("Expected Visit Time %s, but got %s", expectedRows[i][2], row[2])
		}
		if row[3] != expectedRows[i][3] {
			t.Errorf("Expected Visit Count %s, but got %s", expectedRows[i][3], row[3])
		}
		if row[4] != expectedRows[i][4] {
			t.Errorf("Expected Last Visit Time %s, but got %s", expectedRows[i][4], row[4])
		}
	}
}

func TestExtractFirefoxeHistoryWithBetween(t *testing.T) {
	afterDate, err := time.Parse("2006-01-02", "2025-02-01")
	if err != nil {
		t.Fatalf("Failed to parse after date: %v", err)
	}
	beforeDate, err := time.Parse("2006-01-02", "2025-03-01")
	if err != nil {
		t.Fatalf("Failed to parse before date: %v", err)
	}
	var expectedRows = [][]string{
		{"https://example.com/", "Example Domain", "2025-02-02 09:00:00", "1", "2025-02-02 09:00:00"},
	}

	actualRows := formatAndFilter(rows, afterDate, beforeDate)
	if len(actualRows) != len(expectedRows) {
		t.Errorf("Actual rows %v", actualRows)
		t.Fatalf("Expected %d rows, but got %d", len(expectedRows), len(actualRows))
	}

	for i, row := range actualRows {
		if row[0] != expectedRows[i][0] {
			t.Errorf("Expected URL %s, but got %s", expectedRows[i][0], row[0])
		}
		if row[1] != expectedRows[i][1] {
			t.Errorf("Expected Title %s, but got %s", expectedRows[i][1], row[1])
		}
		if row[2] != expectedRows[i][2] {
			t.Errorf("Expected Visit Time %s, but got %s", expectedRows[i][2], row[2])
		}
		if row[3] != expectedRows[i][3] {
			t.Errorf("Expected Visit Count %s, but got %s", expectedRows[i][3], row[3])
		}
		if row[4] != expectedRows[i][4] {
			t.Errorf("Expected Last Visit Time %s, but got %s", expectedRows[i][4], row[4])
		}
	}
}
