package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

var (
	SymbolCmd = &cobra.Command{
		Use:     "symbol",
		Short:   "List up symbols",
		Long:    `List up symbols of NASDAQ`,
		Example: `  stonk symbol`,
		Run:     runCommandSymbol,
	}
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	tickerMapChannel := make(chan map[string]bool)
	go fetchTickers(tickerMapChannel)

	tickerMap := <-tickerMapChannel
	tickers := []string{}
	for key := range tickerMap {
		tickers = append(tickers, key)
	}
	printSymbols(tickers)
}

func printSymbols(list []string) {
	sort.Strings(list)
	for _, line := range list {
		fmt.Printf("%s\n", line)
	}
}

// fetch nasdaq symbols
func fetchTickers(tickerMapChannel chan map[string]bool) {
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
