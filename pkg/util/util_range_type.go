package util

import (
	"errors"
	"fmt"
	"github.com/piquette/finance-go/datetime"
	"regexp"
	"strconv"
	"time"
)

const (
	unitSecondsDay     = int64(60 * 60 * 24)
	unitSecondsWeek    = unitSecondsDay * 7
	unitSecondsMonth   = unitSecondsDay * 30
	unitSecondsQuarter = unitSecondsMonth * 3
	unitSecondsYear    = unitSecondsDay * 365
)

// e.g. 1D30D, 1D3M, 5D1Y, 1W200D, 1M12M, 1Y20Y
type RangeType struct {
	SampleLength int    // 1,2,3,...
	SampleUnit   string // D, W, M, Q, Y
	RangeLength  int    // 1,2,3,...
	RangeUnit    string // D, W, M, Q, Y
}

func ParseRangeType(text string) (*RangeType, error) {
	regex := regexp.MustCompile("^(\\d+)(D|W|M|Q|Y)(\\d+)(D|W|M|Q|Y)$")
	match := regex.FindStringSubmatch(text)
	if len(match) < 1 {
		return nil, errors.New("Not found matched RangeType format")
	}

	sampleLength, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return nil, err
	}
	if sampleLength < 1 {
		return nil, errors.New("RangeType Sampling Length should be greater than 1")
	}
	rangeLength, err := strconv.ParseInt(match[3], 10, 64)
	if err != nil {
		return nil, err
	}
	if rangeLength < 1 {
		return nil, errors.New("RangeType Range Length should be greater than 1")
	}

	return &RangeType{
		SampleLength: int(sampleLength),
		SampleUnit:   match[2],
		RangeLength:  int(rangeLength),
		RangeUnit:    match[4],
	}, nil
}

func (r RangeType) GetRangeSeconds() int64 {
	switch r.RangeUnit {
	case "D":
		return unitSecondsDay * int64(r.RangeLength)
	case "W":
		return unitSecondsWeek * int64(r.RangeLength)
	case "M":
		return unitSecondsMonth * int64(r.RangeLength)
	case "Q":
		return unitSecondsQuarter * int64(r.RangeLength)
	case "Y":
		return unitSecondsYear * int64(r.RangeLength)
	default:
		panic(fmt.Sprintf("Unknown range unit: %s", r.RangeUnit))
	}
}

func (r RangeType) GetRangeDatetime(t time.Time) (datetime.Datetime, datetime.Datetime) {
	fromSeconds, toSeconds := r.GetUnixTimeRangeTo(t)
	fromTime := time.Unix(fromSeconds, 0).UTC()
	toTime := time.Unix(toSeconds, 0).UTC()

	fromDatetime := *datetime.New(&fromTime)
	toDatetime := *datetime.New(&toTime)
	return fromDatetime, toDatetime
}

func (r RangeType) GetUnixTimeRangeTo(toTime time.Time) (int64, int64) {
	toEpoch := toTime.Unix()
	fromEpoch := toEpoch - r.GetRangeSeconds()
	return fromEpoch, toEpoch
}

func (r RangeType) ToString() string {
	return fmt.Sprintf("%d%s%d%s", r.SampleLength, r.SampleUnit, r.RangeLength, r.RangeUnit)
}

func (r RangeType) GetSampleInterval() datetime.Interval {
	switch r.SampleUnit {
	case "D":
		return datetime.OneDay
	case "W":
		return datetime.OneDay
	case "M":
		return datetime.OneMonth
	case "Q":
		return datetime.OneMonth
	case "Y":
		return datetime.OneMonth
	}
	panic("Unknown sample unit: " + r.SampleUnit)
}
