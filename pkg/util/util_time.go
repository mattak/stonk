package util

import "time"

func GetNextRangeDateFunction(unit string, length int) func(curr time.Time) time.Time {
	f := GetNextRangeDateUnitFunction(unit)
	return func(curr time.Time) time.Time {
		t := curr
		for i := 0; i < length; i++ {
			t = f(t)
		}
		return t
	}
}

func GetNextRangeDateUnitFunction(unit string) func(curr time.Time) time.Time {
	switch unit {
	case "D":
		return func(curr time.Time) time.Time {
			return curr.AddDate(0, 0, 1)
		}
	case "W":
		return func(curr time.Time) time.Time {
			return curr.AddDate(0, 0, 7-int(curr.Weekday()))
		}
	case "M":
		return func(curr time.Time) time.Time {
			if int(curr.Month()) == 12 {
				return time.Date(curr.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
			}
			return time.Date(curr.Year(), curr.Month()+1, 1, 0, 0, 0, 0, time.UTC)
		}
	case "Q":
		return func(curr time.Time) time.Time {
			month := int(curr.Month())
			if month < 4 {
				return time.Date(curr.Year(), 4, 1, 0, 0, 0, 0, time.UTC)
			} else if month < 7 {
				return time.Date(curr.Year(), 7, 1, 0, 0, 0, 0, time.UTC)
			} else if month < 10 {
				return time.Date(curr.Year(), 10, 1, 0, 0, 0, 0, time.UTC)
			}
			return time.Date(curr.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
		}
	case "Y":
		return func(curr time.Time) time.Time {
			return time.Date(curr.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
		}
	}
	panic("ERROR: unknown unit: " + unit)
}
