package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ollybritton/cal8/pkg/cal8"
	"github.com/spf13/cobra"
	"github.com/tj/go-naturaldate"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cal8",
	Short: "cal8 is an alertnative calendar utility.",
	Long: `cal8 is an experiment in a different calendar system.
	
Under the cal8 system, weeks are 8 days long and consist of the typical Monday
to Friday, plus an extra "Valday".`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cal := getCal(cmd)

		compare, err := cmd.PersistentFlags().GetBool("compare")
		if err != nil {
			fmt.Println("error parsing compare flag:", err)
			os.Exit(1)
		}

		var date time.Time

		if len(args) == 0 {
			date = time.Now()
		} else {
			date, err = naturaldate.Parse(
				strings.Join(args, " "),
				time.Now(),
				naturaldate.WithDirection(naturaldate.Past),
			)

			if err != nil {
				fmt.Println("error parsing date specified:", err)
				os.Exit(1)
			}
		}

		// fmt.Println(date)
		monthView(cal, date, date.AddDate(0, 2, 0), compare)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Flags which apply to every command.
	rootCmd.PersistentFlags().StringP("start-date", "s", "2020-07-06", "day to start the calendar, YYYY-MM-DD format")
	rootCmd.PersistentFlags().StringArrayP("days", "d", cal8.Days, "days to use for the calendar")
	rootCmd.PersistentFlags().BoolP("compare", "c", false, "toggles compare mode. In compare mode, the new calendar along with its variant are shown side by side")
}

// getCal parses cmd's flags to return a cal8.Calendar struct. If it encounters an error, it will exit instead of
// returning it.
func getCal(cmd *cobra.Command) cal8.Calendar {
	startDateRaw, err := cmd.Flags().GetString("start-date")
	if err != nil {
		fmt.Println("error parsing start date flag:", err)
		os.Exit(1)
	}

	startDate, err := time.Parse("2006-01-02", startDateRaw)
	if err != nil {
		fmt.Println("error parsing start date flag:", err)
		os.Exit(1)
	}

	days, err := cmd.Flags().GetStringArray("days")
	if err != nil {
		fmt.Println("error parsing days flag:", err)
		os.Exit(1)
	}

	return cal8.NewCalendar(days, startDate.Year(), int(startDate.Month()), startDate.Day())
}

// monthView shows the calendar for the months specified.
func monthView(cal cal8.Calendar, start, end time.Time, compare bool) {
	if compare {
		// show side by side comparison on modified calendar and normal.
		for curr := start; !(curr.Year() == end.Year() && curr.Month() > end.Month()); curr = curr.Local().AddDate(0, 1, 0) {
			modified := cal.StringMonth(curr.Year(), int(curr.Month()))
			normal := cal8.GetNormalDay.StringMonth(curr.Year(), int(curr.Month()))

			fmt.Println(cal8.HorizontalAppend(modified, normal, " |  "))
			fmt.Println("")
		}

		return
	}

	// print modified calendar in groups of two
	for curr := start; !(curr.Year() == end.Year() && curr.Month() > end.Month()); curr = curr.Local().AddDate(0, 2, 0) {
		currNext := curr.AddDate(0, 1, 0)
		m1 := cal.StringMonth(curr.Year(), int(curr.Month()))

		if currNext.Year() == end.Year() && currNext.Month() > end.Month() {
			fmt.Println(m1)
			fmt.Println("")
			break
		}

		m2 := cal.StringMonth(currNext.Year(), int(currNext.Month()))

		fmt.Println(cal8.HorizontalAppend(m1, m2, "   "))
		fmt.Println("")
	}
}
