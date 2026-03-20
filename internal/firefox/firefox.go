package firefox

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/McFlip/histdump/histdump/internal/firefox/sqlc_ff"
	"github.com/McFlip/histdump/histdump/internal/util"
	_ "modernc.org/sqlite"
)

func parseSqlite(dbFile string) []sqlc_ff.GetVisitsRow {
	// Open the SQLite database file
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create a new Queries instance
	queries := sqlc_ff.New(db)

	// Get visits from the database
	rows, err := queries.GetVisits(context.Background())
	if err != nil {
		log.Fatalf("Failed to get visits: %v", err)
	}

	return rows
}

func formatAndFilter(rows []sqlc_ff.GetVisitsRow, afterDate, beforeDate time.Time) [][]string {
	var results [][]string

	for _, row := range rows {
		urlStr, _ := row.Url.(string)
		titleStr, _ := row.Title.(string)
		visitTime := firefoxDateToTime(row.VisitDate.Int64)
		if !util.FilterDates(visitTime, afterDate, beforeDate) {
			continue // Skip visits outside the date range
		}
		visitCount := ""
		if row.VisitCount.Valid {
			visitCount = fmt.Sprintf("%d", row.VisitCount.Int64)
		}
		lastVisitTime := firefoxDateToTime(row.LastVisitDate.Int64)
		record := []string{urlStr, titleStr, visitTime.Format("2006-01-02 15:04:05"), visitCount, lastVisitTime.Format("2006-01-02 15:04:05")}
		results = append(results, record)
	}
	return results
}

func firefoxDateToTime(timestamp int64) time.Time {
	// Convert microseconds to seconds and nanoseconds
	return time.Unix(
		timestamp/1000000,        // seconds
		(timestamp%1000000)*1000, // nanoseconds
	).UTC()
}

func ExtractFirefoxHistory(dbFile, outputFile, afterDate, beforeDate string) {
	afterDateTime, beforeDateTime := util.ParseDateRange(afterDate, beforeDate)

	// Parse the SQLite database file
	rows := parseSqlite(dbFile)
	// Export the rows to a CSV file
	results := formatAndFilter(rows, afterDateTime, beforeDateTime)
	util.ExportToCSV(results, outputFile)
}
