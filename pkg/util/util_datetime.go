package util

import (
	"github.com/piquette/finance-go/datetime"
	"time"
)

func ParseDatetime(dateString string, defaultValue datetime.Datetime) datetime.Datetime {
	dateLayout := "2006-01-02"
	if len(dateString) < 1 {
		return defaultValue
	}
	if t, err := time.Parse(dateLayout, dateString); err != nil {
		panic(err)
	} else {
		return datetime.Datetime{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}
	}
}
