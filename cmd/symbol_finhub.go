package cmd

import (
	"context"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func FetchFinhubSymbols(apiKey string, symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}

	client := finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: apiKey,
	})

	stockSymbols, _, err := client.StockSymbols(auth, "US")
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

	symbolMapChannel <- symbolMap
}
