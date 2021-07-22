package symbol

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func PrintSymbols(symbolMap map[string]SymbolInfo, outputFormat string) {
	symbols := []string{}
	for k, _ := range symbolMap {
		symbols = append(symbols, k)
	}
	sort.Strings(symbols)

	switch outputFormat {
	case "tsv":
		printSymbolsByTsv(symbols, symbolMap)
		break
	case "json":
		printSymbolsByJson(symbols, symbolMap)
		break
	}
}

func printSymbolsByTsv(sortedSymbols []string, symbolMap map[string]SymbolInfo) {
	for _, symbol := range sortedSymbols {
		fmt.Printf("%s\t%s\n", symbol, symbolMap[symbol].Name)
	}
}

func printSymbolsByJson(sortedSymbols []string, symbolMap map[string]SymbolInfo) {
	symbolList := make([]SymbolInfo, len(sortedSymbols))
	for i := 0; i < len(sortedSymbols); i++ {
		symbolList[i] = symbolMap[sortedSymbols[i]]
	}

	bytes, err := json.Marshal(symbolList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
