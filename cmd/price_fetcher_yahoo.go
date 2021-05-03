package cmd

import (
	"fmt"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"os"
	"time"
)

func FetchYahooPriceCandles(params *chart.Params) PriceCandles {
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
