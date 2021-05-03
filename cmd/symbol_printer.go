package cmd

import (
	"fmt"
	"sort"
)

func PrintSymbols(list []string) {
	sort.Strings(list)
	for _, line := range list {
		fmt.Printf("%s\n", line)
	}
}

