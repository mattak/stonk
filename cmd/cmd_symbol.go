package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol [DATASOURCE]",
		Short: "List up symbols",
		Long:  `List up symbols`,
		Example: `  stonk symbol
  stonk symbol jpx
  stonk symbol finhub_us
  stonk symbol finhub_t
  stonk symbol eodata_nasdaq
  stonk symbol yahoo_kabu
  stonk symbol yahoo_etf
  stonk symbol datahub_nasdaq

Note:
  finhub requires environment variable: FINHUB_API_KEY
`,
		Run: runCommandSymbol,
	}
	RetryLimit   = 10
	outputFormat = "tsv"
)

func init() {
	SymbolCmd.Flags().StringVarP(&outputFormat, "format", "f", "tsv", "Output format: tsv(default), json")
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
	case "yahoo_kabu":
		go FetchYahooKabuSymbols(symbolMapChannel)
		break
	case "yahoo_etf":
		go FetchYahooEtfSymbols(symbolMapChannel)
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
	case "jpx":
		go FetchJpxSymbols(symbolMapChannel)
		break
	default:
		log.Fatalln("undefined marketType")
	}

	symbolMap := <-symbolMapChannel
	PrintSymbols(symbolMap, outputFormat)
}
