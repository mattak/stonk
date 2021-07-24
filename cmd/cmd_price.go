package cmd

import (
	"fmt"
	"github.com/mattak/stonk/pkg/price"
	"github.com/mattak/stonk/pkg/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	PriceCmd = &cobra.Command{
		Use:   "price",
		Short: "Fetch historical price data",
		Long:  `Fetch historical price data from yahoo finance`,
		Example: `  stonk price AAPL
  stonk price -r 1D30D AAPL
  stonk price -r 1M1Y -s 2020-12-01 -e 2021-02-01 U
`,
		Run: runCommandPrice,
	}
	argumentStartDate string
	argumentEndDate   string
	argumentRangeType string
)

func init() {
	PriceCmd.Flags().StringVarP(&argumentStartDate, "start", "s", "", "Start date to fetch symbol data. e.g. 2000-01-01")
	PriceCmd.Flags().StringVarP(&argumentEndDate, "end", "e", "", "End date to fetch symbol data. e.g. 2021-01-01")
	PriceCmd.Flags().StringVarP(&argumentRangeType, "range", "r", "1D30D", "Range type, format: '[0-9]+(D|W|M|Q|Y)[0-9]+(D|W|M|Q|Y)' e.g. 1D30D, 1W3Q, 1M1Y, ...")
}

func runCommandPrice(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("ERROR: please specify ticker symbol")
	}
	tickerSymbol := args[0]
	rangeType, err := util.ParseRangeType(argumentRangeType)
	if err != nil {
		log.Fatalln("ERROR: parse RangeType: ", argumentRangeType, err)
	}

	params := util.CreateChartParamByRangeType(tickerSymbol, *rangeType)
	if argumentStartDate != "" {
		startDatetime := util.ParseDatetimeOrDefault(argumentStartDate, util.NewDatetimeUTC(2000, 1, 1))
		params.Start = &startDatetime
	}
	if argumentEndDate != "" {
		endDatetime := util.ParseDatetimeOrDefault(argumentEndDate, util.NowDatetimeUTC())
		params.End = &endDatetime
	}

	candles, err := price.FetchYahooPriceCandles(&params, rangeType, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: fetching yahoo prices: ", tickerSymbol)
		log.Fatalln(err)
	}
	fmt.Println(strings.Join(candles.ToTsv(), "\n"))
}
