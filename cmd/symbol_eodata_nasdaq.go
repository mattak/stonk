package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

// fetch nasdaq symbols
func FetchEodataNasdaqSymbols(tickerMapChannel chan map[string]bool) {
	tickerMap := map[string]bool{}
	c := colly.NewCollector()
	pages := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pageIndex := 0

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("table.quotes", func(p int, e *colly.HTMLElement) bool {
			e.ForEach("tr", func(p int, e *colly.HTMLElement) {
				e.ForEachWithBreak("td", func(p int, e *colly.HTMLElement) bool {
					tickerMap[e.Text] = true
					return false
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
			tickerMapChannel <- tickerMap
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
