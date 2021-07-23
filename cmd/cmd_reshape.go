package cmd

import (
	"fmt"
	"github.com/mattak/stonk/pkg/price"
	"github.com/mattak/stonk/pkg/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	ReshapeCmd = &cobra.Command{
		Use:   "reshape",
		Short: "Reshape price data as specific sampling",
		Long:  `Reshape price data as specified sampling`,
		Example: `  stonk reshape 1M1Y < /tmp/AAPL.tsv
`,
		Run: runCommandReshape,
	}
)

func init() {
	ReshapeCmd.Flags().StringVarP(&argumentStartDate, "start", "s", "", "Start date to fetch symbol data. e.g. 2000-01-01")
	ReshapeCmd.Flags().StringVarP(&argumentEndDate, "end", "e", "", "End date to fetch symbol data. e.g. 2021-01-01")
}

func runCommandReshape(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("ERROR: please specify ticker symbol")
	}
	rangeTypeText := args[0]
	rangeType, err := util.ParseRangeType(rangeTypeText)
	if err != nil {
		log.Fatalln("ERROR: parse RangeType: ", argumentRangeType, err)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	candles := price.ParsePriceCandlesTSV(string(bytes))
	startDatetime, endDatetime := rangeType.GetRangeDatetime(time.Now())
	if argumentStartDate != "" {
		startDatetime = util.ParseDatetimeOrDefault(argumentStartDate, startDatetime)
	}
	if argumentEndDate != "" {
		endDatetime = util.ParseDatetimeOrDefault(argumentEndDate, endDatetime)
	}
	candles = candles.ReduceRange(startDatetime, endDatetime)
	candles = candles.ReduceSample(rangeType.SampleUnit, rangeType.SampleLength)
	fmt.Println(strings.Join(candles.ToTsv(), "\n"))
}
