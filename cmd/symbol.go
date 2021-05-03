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
  stonk symbol eodata-nasdaq
  stonk symbol yahoo-tosho
  stonk symbol datahub-nasdaq`,
		Run: runCommandSymbol,
	}
	RetryLimit = 10
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	symbolMapChannel := make(chan map[string]SymbolInfo)
	marketType := "eodata-nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	switch marketType {
	case "eodata-nasdaq":
		go FetchEodataNasdaqSymbols(symbolMapChannel)
		break;
	case "yahoo-tosho":
		go FetchYahooToshoSymbols(symbolMapChannel)
		break;
	case "datahub-nasdaq":
		go FetchDatahubNasdaqListing(symbolMapChannel)
		break;
	default:
		log.Fatalln("undefined marketType")
	}

	tickerMap := <-symbolMapChannel
	PrintSymbols(tickerMap)
}
