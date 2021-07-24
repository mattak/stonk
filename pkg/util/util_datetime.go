package util

import (
	"github.com/piquette/finance-go/datetime"
	"time"
)

func NewDatetimeUTC(year int, month int, day int) datetime.Datetime {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return *datetime.New(&t)
}

func NowDatetimeUTC() datetime.Datetime {
	t := time.Now().UTC()
	return *datetime.New(&t)
}


func ParseDatetimeOrDefault(dateString string, defaultValue datetime.Datetime) datetime.Datetime {
	dateLayout := "2006-01-02"
	if len(dateString) < 1 {
		return defaultValue
	}
	if t, err := time.ParseInLocation(dateLayout, dateString, time.UTC); err != nil {
		panic(err)
	} else {
		return *datetime.New(&t)
	}
}

func ParseDatetime(dateString string) *datetime.Datetime {
	dateLayout := "2006-01-02"
	if len(dateString) < 1 {
		return nil
	}
	if t, err := time.ParseInLocation(dateLayout, dateString, time.UTC); err != nil {
		return nil
	} else {
		return datetime.New(&t)
	}
}
