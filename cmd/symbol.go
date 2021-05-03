package cmd

import (
	"github.com/spf13/cobra"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "List up symbols",
		Long:  `List up symbols of NASDAQ`,
		Example: `  stonk symbol
  stonk symbol nasdaq
  stonk symbol tokyo`,
		Run: runCommandSymbol,
	}
	RetryLimit = 10
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	tickerMapChannel := make(chan map[string]bool)
	marketType := "nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	if marketType == "tokyo" {
		go FetchYahooTokyoSymbols(tickerMapChannel)
	} else {
		go FetchEodataNasdaqSymbols(tickerMapChannel)
	}

	tickerMap := <-tickerMapChannel
	tickers := []string{}
	for key := range tickerMap {
		tickers = append(tickers, key)
	}
	PrintSymbols(tickers)
}
