package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "format parses a Go time format string and prints the result",
	Long: `format parses a Got time format string with support for a modified calendar.

It uses Go's time.Time.Format with two differences:
- Day names are correct for the calendar being used.
- The actual day can be specified using %A and %a.

For example:
cal8 format "Monday (%a), 2006-01-02 15:04"
-> Wednesday (Fri), 2020-07-24 08:46`,
	Run: func(cmd *cobra.Command, args []string) {
		cal := getCal(cmd)

		strTime, err := cmd.Flags().GetString("time")
		if err != nil {
			fmt.Println("error parsing 'time' flag:", err)
			os.Exit(1)
		}

		t, err := time.Parse("2006-01-02T15:04:05-0700", strTime)
		if err != nil {
			fmt.Println("error parsing 'time' flag:", err)
			os.Exit(1)
		}

		fmt.Println(cal.Format(t, strings.Join(args, "")))
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	formatCmd.Flags().StringP("time", "t", time.Now().Format("2006-01-02T15:04:05-0700"), "time to format for in ISO 8601 format. Defaults to today.")
}
