package cmd

import (
	"encoding/csv"
	"net/http"
)

func FetchDatahubNasdaqListing(symbolMapChannel chan map[string]SymbolInfo) {
	url := "https://datahub.io/core/nasdaq-listings/r/nasdaq-listed.csv"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("UserAgent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	csvreader := csv.NewReader(res.Body)
	csvreader.LazyQuotes = true
	csvreader.TrimLeadingSpace = true
	data, _ := csvreader.ReadAll()

	symbolMap := map[string]SymbolInfo{}
	for i := 0; i < len(data); i++ {
		symbol := data[i][0]
		name := data[i][1]
		symbolMap[symbol] = SymbolInfo{Symbol: symbol, Name: name}
	}

	symbolMapChannel <- symbolMap
}
