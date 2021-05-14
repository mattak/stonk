package cmd

import (
	"fmt"
	"sort"
)

func PrintSymbols(symbolMap map[string]SymbolInfo) {
	symbols := []string{}
	for k, _ := range symbolMap {
		symbols = append(symbols, k)
	}
	sort.Strings(symbols)

	for _, symbol := range symbols {
		fmt.Printf("%s\t%s\n", symbol, symbolMap[symbol].Name)
	}
}
