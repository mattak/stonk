package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"sort"
)

var (
	SymbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "List up symbols",
		Long:  `List up symbols of NASDAQ`,
		Example: `  stonk symbol
  stonk symbol nasdaq
  stonk symbol tokyo`,
		Run: runCommandSymbol,
	}
)

func init() {
}

func runCommandSymbol(cmd *cobra.Command, args []string) {
	tickerMapChannel := make(chan map[string]bool)
	marketType := "nasdaq"
	if len(args) >= 1 {
		marketType = args[0]
	}

	if marketType == "tokyo" {
		go fetchTokyoTickers(tickerMapChannel)
	} else {
		go fetchNasdaqTickers(tickerMapChannel)
	}

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

func fetchTokyoTickers(tickerMapChannel chan map[string]bool) {
	codes := map[string]bool{}
	page := 1
	re := regexp.MustCompile(`code=([\d\.\w]+)`)
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("div#contents-body-bottom table.rankingTable tbody tr", func(y int, e *colly.HTMLElement) {
			line := []string{}
			e.ForEach("td a", func(x int, e *colly.HTMLElement) {
				link := e.Attr("href")
				if len(link) > 0 {
					result := re.FindAllStringSubmatch(link, -1)
					if len(result) > 0 && len(result[0]) > 1 {
						codes[result[0][1]] = true
					}
				}
				line = append(line, e.Text)
			})
		})

		pageLinks := []string{}
		e.ForEach("div#contents-body-bottom ul.ymuiPagingBottom a", func(y int, e *colly.HTMLElement) {
			pageLinks = append(pageLinks, e.Text)
		})

		if pageLinks[len(pageLinks)-1] == "次へ" {
			page++
			url := fmt.Sprintf("https://info.finance.yahoo.co.jp/ranking/?kd=4&tm=d&vl=a&mk=1&p=%d", page)
			err := e.Request.Visit(url)
			if err != nil {
				panic(err)
			}
		} else {
			tickerMapChannel <- codes
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stderr, "Visiting", r.URL.String())
	})

	err := c.Visit(fmt.Sprintf("https://info.finance.yahoo.co.jp/ranking/?kd=4&tm=d&vl=a&mk=1&p=%d", page))
	if err != nil {
		panic(err)
	}
}

// fetch nasdaq symbols
func fetchNasdaqTickers(tickerMapChannel chan map[string]bool) {
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
