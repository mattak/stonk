package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "List up symbols",
		Long:  `List up symbols of NASDAQ`,
		Example: `  stonk symbol
  stonk symbol eodata_nasdaq
  stonk symbol yahoo_tosho
  stonk symbol datahub_nasdaq`,
		Run: runCommandSymbol,
	}
	RetryLimit = 10
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	symbolMapChannel := make(chan map[string]SymbolInfo)
	marketType := "eodata_nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	switch marketType {
	case "eodata_nasdaq":
		go FetchEodataNasdaqSymbols(symbolMapChannel)
		break;
	case "yahoo_tosho":
		go FetchYahooToshoSymbols(symbolMapChannel)
		break;
	case "datahub_nasdaq":
		go FetchDatahubNasdaqListing(symbolMapChannel)
		break;
	default:
		log.Fatalln("undefined marketType")
	}

	tickerMap := <-symbolMapChannel
	PrintSymbols(tickerMap)
}
