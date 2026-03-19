package chrome

import (
	"encoding/csv"
	"os"
	"testing"
	"time"

	"gitlab1.usace.army.mil/csd-f/browser-history/internal/chrome/sqlc"
)

func TestExtractChromeHistory(t *testing.T) {
	// Test data
	dbFile := "testdata/History_Chrome"
	outputFile := "testoutput/chrome_history.csv"

	// Remove the output file if it exists
	_ = os.Remove(outputFile)

	ExtractChromeHistory(dbFile, outputFile, "", "")

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
	const expectedVisitTime = "2025-07-01 16:50:00"
	if records[1][2] != expectedVisitTime {
		t.Errorf("Expected Visit Time %s, but got %s", expectedVisitTime, records[1][2])
	}
	const expectedVisitCount = "1"
	if records[1][3] != expectedVisitCount {
		t.Errorf("Expected Visit Count %s, but got %s", expectedVisitCount, records[1][3])
	}
	const expectedLastVisitTime = "2025-07-01 16:50:00"
	if records[1][4] != expectedLastVisitTime {
		t.Errorf("Expected Last Visit Time %s, but got %s", expectedLastVisitTime, records[1][4])
	}
}

func TestExtractEdgeHistory(t *testing.T) {
	// Test data
	dbFile := "testdata/History_Edge"
	outputFile := "testoutput/edge_history.csv"

	// Remove the output file if it exists
	_ = os.Remove(outputFile)

	ExtractChromeHistory(dbFile, outputFile, "", "")

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
	const expectedVisitTime = "2025-07-01 16:50:12"
	if records[1][2] != expectedVisitTime {
		t.Errorf("Expected Visit Time %s, but got %s", expectedVisitTime, records[1][2])
	}
	const expectedVisitCount = "1"
	if records[1][3] != expectedVisitCount {
		t.Errorf("Expected Visit Count %s, but got %s", expectedVisitCount, records[1][3])
	}
	const expectedLastVisitTime = "2025-07-01 16:50:12"
	if records[1][4] != expectedLastVisitTime {
		t.Errorf("Expected Last Visit Time %s, but got %s", expectedLastVisitTime, records[1][4])
	}
}

var rows = []sqlc.GetVisitsRow{
	{
		Url:           "https://time.gov/",
		Title:         "National Institute of Standards and Technology | NIST",
		VisitTime:     13385466000000000, // "2025-03-03 09:00:00 UTC"
		VisitCount:    1,
		LastVisitTime: 13385466000000000, // "2025-03-03 09:00:00 UTC"
	},
	{
		Url:           "https://example.com/",
		Title:         "Example Domain",
		VisitTime:     13382960400000000, // "2025-02-02 09:00:00 UTC"
		VisitCount:    1,
		LastVisitTime: 13382960400000000, // "2025-02-02 09:00:00 UTC"
	},
	{
		Url:           "https://xkcd.com/",
		Title:         "A webcomic of romance, sarcasm, math, and language.",
		VisitTime:     13382096400000000, // "2025-01-23 09:00:00 UTC"
		VisitCount:    1,
		LastVisitTime: 13382096400000000, // "2025-01-23 09:00:00 UTC"
	},
}

func TestExtractChromeHistoryWithAfter3Results(t *testing.T) {
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

func TestExtractChromeHistoryWithAfter2Results(t *testing.T) {
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

func TestExtractChromeHistoryWithBefore3Results(t *testing.T) {
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

func TestExtractChromeHistoryWithBefore2Results(t *testing.T) {
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

func TestExtractChromeHistoryWithBetween(t *testing.T) {
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
