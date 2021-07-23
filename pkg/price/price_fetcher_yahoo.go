package price

import (
	"errors"
	"fmt"
	"github.com/mattak/stonk/pkg/util"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"os"
	"sync"
	"time"
)

func FetchYahooPriceCandles(params *chart.Params, rangeType *util.RangeType, retryCount int) (PriceCandles, error) {
	candles, err := fetchYahooPriceCandlesWithRetry(params, retryCount)
	if err != nil {
		return nil, err
	}
	return candles.ReduceSample(rangeType.SampleUnit, rangeType.SampleLength), nil
}

func fetchYahooPriceCandlesWithRetry(params *chart.Params, retryCount int) (PriceCandles, error) {
	waitTime := time.Duration(1)

	for retryCount > 0 {
		candles, err := fetchYahooPriceCandlesWithoutRetry(params)
		if err == nil {
			return candles, nil
		}

		fmt.Fprintln(os.Stderr, "retry:", retryCount, "symbol:", params.Symbol)
		retryCount--

		if retryCount <= 0 {
			return nil, err
		}

		// exponential backoff
		wg := sync.WaitGroup{}
		time.Sleep(time.Second * waitTime)
		wg.Wait()
		waitTime = waitTime * 2
	}

	return nil, errors.New(fmt.Sprintf("unexpected error with retryCount %d", retryCount))
}

func fetchYahooPriceCandlesWithoutRetry(params *chart.Params) (PriceCandles, error) {
	iter := chart.Get(params)
	var candles []PriceCandle

	for iter.Next() {
		t := time.Unix(int64(iter.Bar().Timestamp), 0)
		date := datetime.Datetime{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}
		bar := iter.Bar()
		candle := PriceCandle{
			Date:   date,
			Open:   bar.Open.BigFloat(),
			Close:  bar.Close.BigFloat(),
			High:   bar.High.BigFloat(),
			Low:    bar.Low.BigFloat(),
			Volume: int64(bar.Volume),
		}
		candles = append(candles, candle)
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return candles, nil
}
