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
  stonk symbol tosho`,
		Run: runCommandSymbol,
	}
	RetryLimit = 10
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	symbolMapChannel := make(chan map[string]SymbolInfo)
	marketType := "nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	if marketType == "tosho" {
		go FetchYahooToshoSymbols(symbolMapChannel)
	} else {
		go FetchEodataNasdaqSymbols(symbolMapChannel)
	}

	tickerMap := <-symbolMapChannel
	PrintSymbols(tickerMap)
}
