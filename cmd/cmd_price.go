package cmd

import (
	"fmt"
	"github.com/mattak/stonk/pkg/price"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"time"
)

var (
	PriceCmd = &cobra.Command{
		Use:   "price",
		Short: "Fetch historical price data",
		Long:  `Fetch historical price data from yahoo finance`,
		Example: `  stonk price AAPL
  stonk price -s 2020-12-01 -e 2021-02-01 -r day U`,
		Run: runCommandPrice,
	}
	argumentStartDate string
	argumentEndDate   string
	argumentRangeType string
)

func init() {
	PriceCmd.Flags().StringVarP(&argumentStartDate, "start", "s", "2000-01-01", "Start date to fetch symbol data")
	PriceCmd.Flags().StringVarP(&argumentEndDate, "end", "e", "", "End date to fetch symbol data")
	PriceCmd.Flags().StringVarP(&argumentRangeType, "range", "r", "month", "Interval type, month or day")
}

func parseDate(dateString string, defaultValue datetime.Datetime) datetime.Datetime {
	dateLayout := "2006-01-02"
	if len(dateString) < 1 {
		return defaultValue
	}
	if t, err := time.Parse(dateLayout, dateString); err != nil {
		panic(err)
	} else {
		return datetime.Datetime{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}
	}
}

func runCommandPrice(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("ERROR: please specify ticker symbol")
	}
	tickerSymbol := args[0]

	t := time.Now()
	startDate := parseDate(argumentStartDate, datetime.Datetime{Year: 2000, Month: 01, Day: 01})
	endDate := parseDate(argumentEndDate, datetime.Datetime{Year: t.Year(), Month: int(t.Month()), Day: t.Day()})

	var interval datetime.Interval
	if argumentRangeType == "day" {
		interval = datetime.OneDay
	} else {
		interval = datetime.OneMonth
	}

	params := &chart.Params{
		Symbol:   tickerSymbol,
		Interval: interval,
		Start:    &startDate,
		End:      &endDate,
	}

	candles, err := price.FetchYahooPriceCandlesWithRetry(params, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: fetching yahoo prices: ", tickerSymbol)
		log.Fatalln(err)
	}
	fmt.Println(strings.Join(candles.ToTsv(), "\n"))
}
