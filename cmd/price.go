package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/spf13/cobra"
	"log"
	"math/big"
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

type PriceCandle struct {
	Date   datetime.Datetime `json:"date"`
	Open   *big.Float        `json:"open"`
	Close  *big.Float        `json:"close"`
	High   *big.Float        `json:"high"`
	Low    *big.Float        `json:"low"`
	Volume int               `json:"volume"`
}

type PriceCandles []PriceCandle

func (pc PriceCandle) ToTsvHeader() string {
	return "date\topen\tclose\thigh\tlow\tvolume"
}

func (pc PriceCandle) ToTsv() string {
	return fmt.Sprintf(
		"%04d-%02d-%02d\t%f\t%f\t%f\t%f\t%d",
		pc.Date.Year, pc.Date.Month, pc.Date.Day,
		pc.Open,
		pc.Close,
		pc.High,
		pc.Low,
		pc.Volume,
	)
}

func (pcs PriceCandles) ToTsv() []string {
	lines := []string{PriceCandle{}.ToTsvHeader()}
	for _, line := range pcs {
		lines = append(lines, line.ToTsv())
	}
	return lines
}

func (pc PriceCandle) ToJson() string {
	jsonValue, _ := json.Marshal(pc)
	return string(jsonValue)
}

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

	candles := fetchPriceCandles(params)
	fmt.Println(strings.Join(candles.ToTsv(), "\n"))
}

func fetchPriceCandles(params *chart.Params) PriceCandles {
	iter := chart.Get(params)
	candles := []PriceCandle{}

	for iter.Next() {
		date := time.Unix(int64(iter.Bar().Timestamp), 0)
		datetime := datetime.Datetime{Year: date.Year(), Month: int(date.Month()), Day: date.Day()}
		bar := iter.Bar()
		candle := PriceCandle{
			Date:   datetime,
			Open:   bar.Open.BigFloat(),
			Close:  bar.Close.BigFloat(),
			High:   bar.High.BigFloat(),
			Low:    bar.Low.BigFloat(),
			Volume: bar.Volume,
		}
		candles = append(candles, candle)
	}

	if err := iter.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return candles
}
