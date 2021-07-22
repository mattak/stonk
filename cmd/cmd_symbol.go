package cmd

import (
	"github.com/mattak/stonk/pkg/symbol"
	"github.com/mattak/stonk/pkg/util"
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
	symbolMapChannel := make(chan map[string]symbol.SymbolInfo)
	marketType := "eodata_nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	switch marketType {
	case "eodata_nasdaq":
		go symbol.FetchEodataNasdaqSymbols(symbolMapChannel)
		break
	case "yahoo_kabu":
		go symbol.FetchYahooKabuSymbols(symbolMapChannel)
		break
	case "yahoo_etf":
		go symbol.FetchYahooEtfSymbols(symbolMapChannel)
		break
	case "datahub_nasdaq":
		go symbol.FetchDatahubNasdaqListing(symbolMapChannel)
		break
	case "finhub_us":
		apiKey := util.LoadEnv("FINHUB_API_KEY")
		go symbol.FetchFinhubSymbols(apiKey, "US", symbolMapChannel)
		break
	case "finhub_t":
		apiKey := util.LoadEnv("FINHUB_API_KEY")
		go symbol.FetchFinhubSymbols(apiKey, "T", symbolMapChannel)
		break
	case "jpx":
		go symbol.FetchJpxSymbols(symbolMapChannel)
		break
	default:
		log.Fatalln("undefined marketType")
	}

	symbolMap := <-symbolMapChannel
	symbol.PrintSymbols(symbolMap, outputFormat)
}
