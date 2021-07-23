package util

import (
	"github.com/piquette/finance-go/chart"
	"time"
)

func CreateChartParamByRangeType(tickerSymbol string, rangeType RangeType) chart.Params {
	t := time.Now()

	startDate, endDate := rangeType.GetRangeDatetime(t)
	interval := rangeType.GetSampleInterval()

	return chart.Params{
		Symbol:   tickerSymbol,
		Interval: interval,
		Start:    &startDate,
		End:      &endDate,
	}
}
