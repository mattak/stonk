package symbol

import (
	"encoding/csv"
	"github.com/mattak/stonk/pkg/util"
)

func FetchDatahubNasdaqListing(symbolMapChannel chan map[string]SymbolInfo) {
	url := "https://datahub.io/core/nasdaq-listings/r/nasdaq-listed.csv"
	res := util.DoHttpGetRequest(url)
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
