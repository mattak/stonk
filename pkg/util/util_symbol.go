package util

import "regexp"

func IsTokyoNoiseSymbol(symbol string) bool {
	re := regexp.MustCompile("^\\d{4}\\.\\w$")
	return !re.MatchString(symbol)
}
