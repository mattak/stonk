package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

// fetch nasdaq symbols
func FetchEodataNasdaqSymbols(symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}
	c := NewColly()
	pages := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pageIndex := 0

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("table.quotes", func(p int, e *colly.HTMLElement) bool {
			e.ForEach("tr", func(p int, e *colly.HTMLElement) {
				info := SymbolInfo{}
				e.ForEachWithBreak("td", func(p int, e *colly.HTMLElement) bool {
					if p == 0 {
						info.Symbol = e.Text
					} else if p == 1 {
						info.Name = e.Text
						symbolMap[info.Symbol] = info
						return false
					}
					return true
				})
			})
			return false
		})

		if pageIndex+1 < len(pages) {
			nextPage := pageIndex + 1
			pageIndex++
			err := e.Request.Visit(fmt.Sprintf("https://eoddata.com/stocklist/NASDAQ/%s.html", pages[nextPage]))
			if err != nil {
				panic(err)
			}
		} else {
			symbolMapChannel <- symbolMap
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stderr, "Visiting", r.URL.String())
	})

	err := c.Visit(fmt.Sprintf("https://eoddata.com/stocklist/NASDAQ/%s.html", pages[pageIndex]))
	if err != nil {
		panic(err)
	}
}
