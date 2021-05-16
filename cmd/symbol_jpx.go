package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type MarketType string

const (
	MarketType_Jpx_ETF_ETN                = MarketType("ETF・ETN")
	MarketType_Jpx_JasdaqGrowth           = MarketType("JASDAQ(グロース・内国株)")
	MarketType_Jpx_JasdaqStandardDomestic = MarketType("JASDAQ(スタンダード・内国株)")
	MarketType_Jpx_JasdaqStandardForeign  = MarketType("JASDAQ(スタンダード・外国株)")
	MarketType_Jpx_ProMarket              = MarketType("PRO Market")
	MarketType_Jpx_REIT                   = MarketType("REIT・ベンチャーファンド・カントリーファンド・インフラファンド")
	MarketType_Jpx_Investment             = MarketType("出資証券")
	MarketType_Jpx_MothorsDomesitic       = MarketType("マザーズ(内国株)")
	MarketType_Jpx_MothersForeign         = MarketType("マザーズ(外国株)")
	MarketType_Jpx_Tosho1Domestic         = MarketType("市場第一部(内国株)")
	MarketType_Jpx_Tosho1Foreign          = MarketType("市場第一部(外国株)")
	MarketType_Jpx_Tosho2Domestic         = MarketType("市場第二部(内国株)")
	MarketType_Jpx_Tosho2Foreign          = MarketType("市場第二部(外国株)")
)

func FetchJpxSymbols(symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}
	urlChannel := make(chan string)
	bodyChannel := make(chan []byte)

	go fetchJpxSymbolExcelLinkUrl(urlChannel)
	url := <-urlChannel

	xlsPath := "/tmp/jpx.xls"
	csvPath := "/tmp/jpx.csv"
	go DownloadFile(url, xlsPath, bodyChannel)
	<-bodyChannel
	fmt.Fprintln(os.Stderr, "File: ", csvPath)

	ConvertXlsFileToCsvFile(xlsPath, csvPath)
	data := ReadCsv(csvPath)
	symbolMap = createSymbolInfo(data)

	symbolMapChannel <- symbolMap
}

func createSymbolInfo(data [][]string) map[string]SymbolInfo {
	symbolMap := map[string]SymbolInfo{}

	for i := 0; i < len(data)-1; i++ {
		//date := data[i+1][0]
		code := data[i+1][1]
		name := NormalizeName(data[i+1][2])
		market := MarketType(NormalizeName(data[i+1][3]))
		if market == MarketType_Jpx_Investment || market == MarketType_Jpx_ProMarket {
			continue
		}
		if len(code) != 4 {
			continue
		}

		info := SymbolInfo{
			Name:   name,
			Symbol: code,
		}
		symbolMap[info.Symbol] = info
	}

	return symbolMap
}

func fetchJpxSymbolExcelLinkUrl(url chan string) {
	href := ""
	c := NewColly()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("div#readArea table tbody tr td a", func(p int, e *colly.HTMLElement) bool {
			href = e.Attr("href")
			return true
		})

		url <- fmt.Sprint("https://www.jpx.co.jp", href)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stderr, "Visiting", r.URL.String())
	})

	err := c.Visit("https://www.jpx.co.jp/markets/statistics-equities/misc/01.html")
	if err != nil {
		panic(err)
	}
}
