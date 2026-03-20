/*
Copyright © 2026 Grady Denton
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/McFlip/histdump/histdump/internal/chrome"
)

// chromeCmd represents the chrome command
var chromeCmd = &cobra.Command{
	Use:   "chrome",
	Short: "Extract Chrome browser history to a CSV file",
	Long: `The input file is a sqlite database file called "History" with no extension.
	The history file is typically located at:
	Windows: %LOCALAPPDATA%\\Google\\Chrome\\User Data\\Default\\History`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Extracting Chrome or Edge browser history...")
		if historyFile == "" || outputFile == "" {
			fmt.Println("Error: Both --file and --output flags are required.")
			fmt.Println("Usage: histdump chrome -f [path to file] -o [output file]")
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
		chrome.ExtractChromeHistory(historyFile, outputFile, afterDate, beforeDate)
	},
}

func init() {
	rootCmd.AddCommand(chromeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chromeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chromeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
