package price

import (
	"encoding/json"
	"fmt"
	"github.com/mattak/stonk/pkg/util"
	"github.com/piquette/finance-go/datetime"
	"math/big"
	"sort"
	"time"
)

type PriceCandle struct {
	Date   datetime.Datetime `json:"date"`
	Open   *big.Float        `json:"open"`
	Close  *big.Float        `json:"close"`
	High   *big.Float        `json:"high"`
	Low    *big.Float        `json:"low"`
	Volume int64             `json:"volume"`
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

func (pcs PriceCandles) SortByDate() {
	sort.SliceStable(pcs, func(i, j int) bool { return pcs[i].Date.Unix() < pcs[i].Date.Unix() })
}

func (pcs PriceCandles) SummarizeRange(date datetime.Datetime, startIndex, endIndex int) PriceCandle {
	_open := pcs[startIndex].Open
	_close := pcs[endIndex].Close
	_high := pcs[startIndex].High
	_low := pcs[startIndex].Low
	_volume := pcs[startIndex].Volume

	for i := startIndex + 1; i <= endIndex; i++ {
		if pcs[i].High.Cmp(_high) > 0 {
			_high = pcs[i].High
		}
		if pcs[i].Low.Cmp(_low) < 0 {
			_low = pcs[i].Low
		}
		_volume += pcs[i].Volume
	}

	return PriceCandle{
		Date:   date,
		Open:   _open,
		Close:  _close,
		High:   _high,
		Low:    _low,
		Volume: _volume,
	}
}

func (pcs PriceCandles) ReduceSampleByNextDate(nextDate func(curr time.Time) time.Time) PriceCandles {
	if len(pcs) < 1 {
		return PriceCandles{}
	}

	minDate := pcs[0].Date.Time().UTC()
	maxDate := pcs[len(pcs)-1].Date.Time().UTC()

	startIndex := 0
	current := minDate
	next := nextDate(current)

	candles := []PriceCandle{}

	for current.Unix() <= maxDate.Unix() && startIndex < len(pcs) {
		endIndex := -1
		for i := startIndex; i < len(pcs); i++ {
			date := pcs[i].Date.Time().UTC()
			if current.Unix() <= date.Unix() && date.Unix() < next.Unix() {
				endIndex = i
			} else {
				break
			}
		}
		if endIndex == -1 {
			current = next
			next = nextDate(current)
			continue
		}

		startDatetime := datetime.New(&current)
		candle := pcs.SummarizeRange(*startDatetime, startIndex, endIndex)
		candles = append(candles, candle)

		current = next
		next = nextDate(current)
		startIndex = endIndex + 1
	}

	return candles
}

func (pcs PriceCandles) ReduceSample(unit string, length int) PriceCandles {
	f := util.GetNextRangeDateFunction(unit, length)
	return pcs.ReduceSampleByNextDate(f)
}

func (pcs PriceCandles) ReduceRange(fromDatetime, toDatetime datetime.Datetime) PriceCandles {
	candles := PriceCandles{}
	for _, pc := range pcs {
		if fromDatetime.Unix() <= pc.Date.Unix() && pc.Date.Unix() <= toDatetime.Unix() {
			candles = append(candles, pc)
		}
	}
	return candles
}
