package cmd

import (
	"context"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func FilterFinhubSymbols(exchange string, dic *map[string]SymbolInfo) {
	if exchange == "T" {
		delete_keys := make([]string, 0, len(*dic))

		for key, _ := range *dic {
			if IsTokyoNoiseSymbol(key) {
				delete_keys = append(delete_keys, key)
			}
		}

		for _, key := range delete_keys {
			delete(*dic, key)
		}
	}
}

func FetchFinhubSymbols(apiKey string, exchange string, symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}

	client := finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: apiKey,
	})

	stockSymbols, _, err := client.StockSymbols(auth, exchange)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(stockSymbols); i++ {
		stock := stockSymbols[i]

		info := SymbolInfo{}
		info.Symbol = stock.Symbol
		info.Name = stock.Description
		symbolMap[info.Symbol] = info
	}

	FilterFinhubSymbols(exchange, &symbolMap)
	symbolMapChannel <- symbolMap
}
