package test

import (
	"github.com/mattak/stonk/pkg/price"
	"github.com/mattak/stonk/pkg/util"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestReduceRange(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		pcs := price.PriceCandles{}
		fromDatetime := util.NewDatetimeUTC(2021, 1, 1)
		toDatetime := util.NewDatetimeUTC(2021, 6, 1)
		result := pcs.ReduceRange(fromDatetime, toDatetime)
		assert.Equal(t, len(result), 0)
	})
	t.Run("1point: in first", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 1, 1),
				Open:   big.NewFloat(1),
				Close:  big.NewFloat(1),
				High:   big.NewFloat(1),
				Low:    big.NewFloat(1),
				Volume: int64(1),
			},
		}
		fromDatetime := util.NewDatetimeUTC(2021, 1, 1)
		toDatetime := util.NewDatetimeUTC(2021, 6, 1)
		result := pcs.ReduceRange(fromDatetime, toDatetime)
		assert.Equal(t, len(result), 1)
	})
	t.Run("1point: in last", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 6, 1),
				Open:   big.NewFloat(1),
				Close:  big.NewFloat(1),
				High:   big.NewFloat(1),
				Low:    big.NewFloat(1),
				Volume: int64(1),
			},
		}
		fromDatetime := util.NewDatetimeUTC(2021, 1, 1)
		toDatetime := util.NewDatetimeUTC(2021, 6, 1)
		result := pcs.ReduceRange(fromDatetime, toDatetime)
		assert.Equal(t, len(result), 1)
	})
	t.Run("1point: out before", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2020, 12, 31),
				Open:   big.NewFloat(1),
				Close:  big.NewFloat(1),
				High:   big.NewFloat(1),
				Low:    big.NewFloat(1),
				Volume: int64(1),
			},
		}
		fromDatetime := util.NewDatetimeUTC(2021, 1, 1)
		toDatetime := util.NewDatetimeUTC(2021, 6, 1)
		result := pcs.ReduceRange(fromDatetime, toDatetime)
		assert.Equal(t, len(result), 0)
	})
	t.Run("1point: out after", func(t *testing.T) {
		pcs := price.PriceCandles{
			price.PriceCandle{
				Date:   util.NewDatetimeUTC(2021, 6, 2),
				Open:   big.NewFloat(1),
				Close:  big.NewFloat(1),
				High:   big.NewFloat(1),
				Low:    big.NewFloat(1),
				Volume: int64(1),
			},
		}
		fromDatetime := util.NewDatetimeUTC(2021, 1, 1)
		toDatetime := util.NewDatetimeUTC(2021, 6, 1)
		result := pcs.ReduceRange(fromDatetime, toDatetime)
		assert.Equal(t, len(result), 0)
	})
}
