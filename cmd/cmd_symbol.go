package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol [DATASOURCE]",
		Short: "List up symbols",
		Long:  `List up symbols of NASDAQ`,
		Example: `  stonk symbol
  stonk symbol finhub_us
  stonk symbol finhub_t
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
		break
	case "yahoo_tosho":
		go FetchYahooToshoSymbols(symbolMapChannel)
		break
	case "datahub_nasdaq":
		go FetchDatahubNasdaqListing(symbolMapChannel)
		break
	case "finhub_us":
		apiKey := LoadEnv("FINHUB_API_KEY")
		go FetchFinhubSymbols(apiKey, "US", symbolMapChannel)
		break
	case "finhub_t":
		apiKey := LoadEnv("FINHUB_API_KEY")
		go FetchFinhubSymbols(apiKey, "T", symbolMapChannel)
		break
	default:
		log.Fatalln("undefined marketType")
	}

	tickerMap := <-symbolMapChannel
	PrintSymbols(tickerMap)
}
