package test

import (
	"github.com/mattak/stonk/pkg/price"
	"github.com/piquette/finance-go/datetime"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestReduceDay(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		pcs := price.PriceCandles{}
		result := pcs.Reduce("D", 1)
		assert.Equal(t, len(result), 0)
	})
	t.Run("1 point", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   datetime.Datetime{Year: 2020, Month: 1, Day: 1},
				Open:   big.NewFloat(2.0),
				Close:  big.NewFloat(4.0),
				High:   big.NewFloat(8.0),
				Low:    big.NewFloat(1.0),
				Volume: int64(1),
			},
		}
		result := pcs.Reduce("D", 1)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0].Open, big.NewFloat(2))
		assert.Equal(t, result[0].Close, big.NewFloat(4))
		assert.Equal(t, result[0].High, big.NewFloat(8))
		assert.Equal(t, result[0].Low, big.NewFloat(1))
		assert.Equal(t, result[0].Volume, int64(1))
	})
	t.Run("2 point: simple add", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   datetime.Datetime{Year: 2020, Month: 1, Day: 1},
				Open:   big.NewFloat(2.0),
				Close:  big.NewFloat(4.0),
				High:   big.NewFloat(4.0),
				Low:    big.NewFloat(2.0),
				Volume: int64(1),
			},
			price.PriceCandle{
				Date:   datetime.Datetime{Year: 2020, Month: 1, Day: 2},
				Open:   big.NewFloat(4.0),
				Close:  big.NewFloat(8.0),
				High:   big.NewFloat(8.0),
				Low:    big.NewFloat(1.0),
				Volume: int64(2),
			},
		}
		result := pcs.Reduce("D", 1)
		assert.Equal(t, len(result), 2)
		assert.Equal(t, result[0].Open, big.NewFloat(2))
		assert.Equal(t, result[0].Close, big.NewFloat(4))
		assert.Equal(t, result[0].High, big.NewFloat(4))
		assert.Equal(t, result[0].Low, big.NewFloat(2))
		assert.Equal(t, result[0].Volume, int64(1))

		assert.Equal(t, result[1].Open, big.NewFloat(4))
		assert.Equal(t, result[1].Close, big.NewFloat(8))
		assert.Equal(t, result[1].High, big.NewFloat(8))
		assert.Equal(t, result[1].Low, big.NewFloat(1))
		assert.Equal(t, result[1].Volume, int64(2))
	})
	t.Run("2 point: sparse day", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   datetime.Datetime{Year: 2020, Month: 1, Day: 1},
				Open:   big.NewFloat(2.0),
				Close:  big.NewFloat(4.0),
				High:   big.NewFloat(4.0),
				Low:    big.NewFloat(2.0),
				Volume: int64(1),
			},
			price.PriceCandle{
				Date:   datetime.Datetime{Year: 2020, Month: 1, Day: 3},
				Open:   big.NewFloat(4.0),
				Close:  big.NewFloat(8.0),
				High:   big.NewFloat(8.0),
				Low:    big.NewFloat(1.0),
				Volume: int64(2),
			},
		}
		result := pcs.Reduce("D", 1)
		assert.Equal(t, len(result), 2)
		assert.Equal(t, result[0].Date.Year, 2020)
		assert.Equal(t, result[0].Date.Month, 1)
		assert.Equal(t, result[0].Date.Day, 1)
		assert.Equal(t, result[0].Open, big.NewFloat(2))
		assert.Equal(t, result[0].Close, big.NewFloat(4))
		assert.Equal(t, result[0].High, big.NewFloat(4))
		assert.Equal(t, result[0].Low, big.NewFloat(2))
		assert.Equal(t, result[0].Volume, int64(1))

		assert.Equal(t, result[1].Date.Year, 2020)
		assert.Equal(t, result[1].Date.Month, 1)
		assert.Equal(t, result[1].Date.Day, 3)
		assert.Equal(t, result[1].Open, big.NewFloat(4))
		assert.Equal(t, result[1].Close, big.NewFloat(8))
		assert.Equal(t, result[1].High, big.NewFloat(8))
		assert.Equal(t, result[1].Low, big.NewFloat(1))
		assert.Equal(t, result[1].Volume, int64(2))
	})
}
