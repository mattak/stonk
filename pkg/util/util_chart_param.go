package util

import (
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
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

func CreateChartParam(tickerSymbol, argumentStartDate, argumentEndDate, argumentInterval string) chart.Params {
	t := time.Now()
	startDate := ParseDatetime(argumentStartDate, datetime.Datetime{Year: 2000, Month: 01, Day: 01})
	endDate := ParseDatetime(argumentEndDate, datetime.Datetime{Year: t.Year(), Month: int(t.Month()), Day: t.Day()})

	var interval datetime.Interval
	if argumentInterval == "day" {
		interval = datetime.OneDay
	} else {
		interval = datetime.OneMonth
	}

	return chart.Params{
		Symbol:   tickerSymbol,
		Interval: interval,
		Start:    &startDate,
		End:      &endDate,
	}
}
