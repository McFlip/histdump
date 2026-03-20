/*
Copyright © 2026 Grady Denton
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/McFlip/histdump/histdump/internal/firefox"
)

// firefoxCmd represents the firefox command
var firefoxCmd = &cobra.Command{
	Use:   "firefox",
	Short: "Extract Firefox browser history to a CSV file",
	Long: `The input file is a sqlite database file called "places.sqlite".
The places file is typically located at:
Windows: %APPDATA%\\Mozilla\\Firefox\\Profiles\\<profile>\\places.sqlite`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Extracting Firefox browser history...")
		if historyFile == "" || outputFile == "" {
			fmt.Println("Error: Both --file and --output flags are required.")
			fmt.Println("Usage: histdump firefox -f [path to file] -o [output file]")
			return
		}
		fmt.Printf("History file: %s\n", historyFile)
		fmt.Printf("Output file: %s\n", outputFile)
		if afterDate != "" {
			fmt.Printf("Filtering for visits after: %s\n", afterDate)
		}
		if beforeDate != "" {
			fmt.Printf("Filtering for visits before: %s\n", beforeDate)
		}
		firefox.ExtractFirefoxHistory(historyFile, outputFile, afterDate, beforeDate)
	},
}

func init() {
	rootCmd.AddCommand(firefoxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// firefoxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// firefoxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
