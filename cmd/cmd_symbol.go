package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "List up symbols",
		Long:  `List up symbols of NASDAQ`,
		Example: `  stonk symbol
  stonk symbol finhub
  stonk symbol eodata_nasdaq
  stonk symbol yahoo_tosho
  stonk symbol datahub_nasdaq

Note:
  finhub requires environment variable: FINHUB_API_KEY
`,
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
	case "finhub":
		apiKey := os.Getenv("FINHUB_API_KEY")
		if len(apiKey) <= 0 {
			log.Fatal("FINHUB_API_KEY is blank")
		}
		go FetchFinhubSymbols(apiKey, symbolMapChannel)
		break;
	default:
		log.Fatalln("undefined marketType")
	}

	tickerMap := <-symbolMapChannel
	PrintSymbols(tickerMap)
}
