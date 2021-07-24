package test

import (
	"github.com/mattak/stonk/pkg/price"
	"github.com/mattak/stonk/pkg/util"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSummerizeRange(t *testing.T) {
	t.Run("1point", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 1),
				Open:   big.NewFloat(2),
				Close:  big.NewFloat(4),
				High:   big.NewFloat(8),
				Low:    big.NewFloat(1),
				Volume: 1,
			},
		}
		dt := util.NewDatetimeUTC(2021, 1, 1)
		result := pcs.SummarizeRange(dt, 0, 0)
		assert.Equal(t, result.Open, big.NewFloat(2))
		assert.Equal(t, result.Close, big.NewFloat(4))
		assert.Equal(t, result.High, big.NewFloat(8))
		assert.Equal(t, result.Low, big.NewFloat(1))
		assert.Equal(t, result.Volume, int64(1))
	})
	t.Run("2point 1range", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 1),
				Open:   big.NewFloat(2),
				Close:  big.NewFloat(4),
				High:   big.NewFloat(4),
				Low:    big.NewFloat(2),
				Volume: 1,
			},
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 2),
				Open:   big.NewFloat(4),
				Close:  big.NewFloat(8),
				High:   big.NewFloat(8),
				Low:    big.NewFloat(1),
				Volume: 1,
			},
		}
		dt := util.NewDatetimeUTC(2021, 1, 1)
		result := pcs.SummarizeRange(dt, 0, 0)
		assert.Equal(t, result.Date.Year, 2021)
		assert.Equal(t, result.Date.Month, 1)
		assert.Equal(t, result.Date.Day, 1)
		assert.Equal(t, result.Open, big.NewFloat(2))
		assert.Equal(t, result.Close, big.NewFloat(4))
		assert.Equal(t, result.High, big.NewFloat(4))
		assert.Equal(t, result.Low, big.NewFloat(2))
		assert.Equal(t, result.Volume, int64(1))
	})
	t.Run("2point 2range", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 1),
				Open:   big.NewFloat(2),
				Close:  big.NewFloat(4),
				High:   big.NewFloat(4),
				Low:    big.NewFloat(2),
				Volume: 1,
			},
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 2),
				Open:   big.NewFloat(4),
				Close:  big.NewFloat(8),
				High:   big.NewFloat(8),
				Low:    big.NewFloat(1),
				Volume: 1,
			},
		}
		dt := util.NewDatetimeUTC(2021, 1, 1)
		result := pcs.SummarizeRange(dt, 0, 1)
		assert.Equal(t, result.Date.Year, 2021)
		assert.Equal(t, result.Date.Month, 1)
		assert.Equal(t, result.Date.Day, 1)
		assert.Equal(t, result.Open, big.NewFloat(2))
		assert.Equal(t, result.Close, big.NewFloat(8))
		assert.Equal(t, result.High, big.NewFloat(8))
		assert.Equal(t, result.Low, big.NewFloat(1))
		assert.Equal(t, result.Volume, int64(2))
	})
}
