/*
Copyright © 2026 Grady Denton
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	historyFile string // Path to the browser history file
	outputFile  string // Path to the output CSV file
	afterDate   string // Filter for visits after this date (YYYY-MM-DD)
	beforeDate  string // Filter for visits before this date (YYYY-MM-DD)
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "browser-history [ chrome | firefox ] -f [path to file] -o [output file] -a [after YYYY-MM-DD] -b [before YYYY-MM-DD]",
		Short: "A CLI tool to extract browser history to a CSV file",
		Long:  `Import browser history sqlite database files from Chrome or Firefox and export them to a CSV file.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.browser-history.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&historyFile, "file", "f", "", "Path to the browser history file (required)")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Path to the output CSV file (required)")
	rootCmd.PersistentFlags().StringVarP(&afterDate, "after", "a", "", "Filter visits after this date (YYYY-MM-DD) (optional)")
	rootCmd.PersistentFlags().StringVarP(&beforeDate, "before", "b", "", "Filter visits before this date (YYYY-MM-DD) (optional)")
}
