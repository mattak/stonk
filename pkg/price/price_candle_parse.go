package price

import (
	"github.com/mattak/stonk/pkg/util"
	"math/big"
	"strconv"
	"strings"
)

func parsePriceCandle(values []string, indexMap map[string]int) *PriceCandle {
	dt := util.ParseDatetime(values[indexMap["date"]])
	if dt == nil {
		return nil
	}
	open, err := strconv.ParseFloat(values[indexMap["open"]], 64)
	if err != nil {
		return nil
	}
	_close, err := strconv.ParseFloat(values[indexMap["close"]], 64)
	if err != nil {
		return nil
	}
	high, err := strconv.ParseFloat(values[indexMap["high"]], 64)
	if err != nil {
		return nil
	}
	low, err := strconv.ParseFloat(values[indexMap["low"]], 64)
	if err != nil {
		return nil
	}
	volume, err := strconv.ParseInt(values[indexMap["volume"]], 10, 64)
	if err != nil {
		return nil
	}

	return &PriceCandle{
		Date:   *dt,
		Open:   big.NewFloat(open),
		Close:  big.NewFloat(_close),
		High:   big.NewFloat(high),
		Low:    big.NewFloat(low),
		Volume: volume,
	}
}

func ParsePriceCandlesTSV(text string) PriceCandles {
	result := PriceCandles{}
	lines := strings.Split(text, "\n")
	if len(lines) < 2 {
		return result
	}

	headers := strings.Split(lines[0], "\t")
	headerIndexMap := map[string]int{}
	for i, header := range headers {
		headerIndexMap[header] = i
	}

	for _, line := range lines[1:] {
		values := strings.Split(line, "\t")
		if len(values) < len(headerIndexMap) {
			continue
		}
		candle := parsePriceCandle(values, headerIndexMap)
		if candle != nil {
			result = append(result, *candle)
		}
	}

	return result
}
