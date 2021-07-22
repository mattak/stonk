package test

import (
	"github.com/mattak/stonk/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNextRangeDate_Day(t *testing.T) {
	t.Run("add1", func(t *testing.T) {
		f := util.GetNextRangeDateFunction("D", 1)
		next := f(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 2)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("add2", func(t *testing.T) {
		f := util.GetNextRangeDateFunction("D", 2)
		next := f(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 3)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
}

func TestGetNextRangeDate_Week(t *testing.T) {
	f := util.GetNextRangeDateFunction("W", 1)
	t.Run("wed", func(t *testing.T) {
		next := f(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 5)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})

	t.Run("thu", func(t *testing.T) {
		next := f(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 5)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("fri", func(t *testing.T) {
		next := f(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 5)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("sat", func(t *testing.T) {
		next := f(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 5)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("sun", func(t *testing.T) {
		next := f(time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 12)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("mon", func(t *testing.T) {
		next := f(time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 12)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("tue", func(t *testing.T) {
		next := f(time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 12)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
}

func TestGetNextRangeDate_Month(t *testing.T) {
	f := util.GetNextRangeDateFunction("M", 1)
	t.Run("month", func(t *testing.T) {
		next := f(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 2)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})

	t.Run("year", func(t *testing.T) {
		next := f(time.Date(2020, 12, 2, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2021)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
}

func TestGetNextRangeDate_Quater(t *testing.T) {
	f := util.GetNextRangeDateFunction("Q", 1)
	t.Run("1Q", func(t *testing.T) {
		next := f(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 4)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})

	t.Run("2Q", func(t *testing.T) {
		next := f(time.Date(2020, 4, 2, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 7)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})

	t.Run("3Q", func(t *testing.T) {
		next := f(time.Date(2020, 7, 2, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2020)
		assert.Equal(t, int(next.Month()), 10)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
	t.Run("4Q", func(t *testing.T) {
		next := f(time.Date(2020, 10, 2, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2021)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
}

func TestGetNextRangeDate_Year(t *testing.T) {
	f := util.GetNextRangeDateFunction("Y", 1)
	t.Run("Y", func(t *testing.T) {
		next := f(time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, next.Year(), 2021)
		assert.Equal(t, int(next.Month()), 1)
		assert.Equal(t, next.Day(), 1)
		assert.Equal(t, next.Hour(), 0)
		assert.Equal(t, next.Minute(), 0)
		assert.Equal(t, next.Second(), 0)
		assert.Equal(t, next.Nanosecond(), 0)
		assert.Equal(t, next.Location(), time.UTC)
	})
}
