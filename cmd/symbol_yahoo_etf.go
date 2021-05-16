package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"regexp"
)

func FetchYahooEtfSymbols(symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}

	page := 1
	re := regexp.MustCompile(`^(\d{4})$`)
	c := NewColly()
	retryCount := 0

	c.OnHTML("body", func(e *colly.HTMLElement) {
		//list = Array.from(document.querySelectorAll(''))

		e.ForEach("div#etf_list table tbody tr", func(y int, e *colly.HTMLElement) {
			code := ""
			info := SymbolInfo{}

			// parse symbol
			e.ForEachWithBreak("td a", func(x int, e *colly.HTMLElement) bool {
				if x == 0 {
					matches := re.FindStringSubmatch(e.Text)
					if matches != nil && len(matches) > 0 {
						code = matches[1]
						info.Symbol = code
					}
					return true
				}

				return false
			})

			// parse name
			e.ForEach("td", func(x int, e *colly.HTMLElement) {
				if x == 2 {
					// string normalize
					info.Name = NormalizeName(e.Text)
				}
			})

			// set result
			if len(code) > 0 {
				symbolMap[code] = info
			}
		})

		pageLink := ""
		e.ForEachWithBreak("div#etf_list div.ymuiPagingTop a.ymuiNext", func(y int, e *colly.HTMLElement) bool {
			pageLink = e.Text
			return true
		})

		if pageLink == "次のページ" {
			page++
			url := fmt.Sprintf("https://stocks.finance.yahoo.co.jp/etf/list/?p=%d", page)
			err := e.Request.Visit(url)
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

	c.OnResponse(func(r *colly.Response) {
		fmt.Fprintln(os.Stderr, "Result", r.StatusCode)

		if r.StatusCode != 200 {
			if retryCount < retryLimit {
				retryCount++
				fmt.Fprintln(os.Stderr, "Retry", retryCount)
				r.Request.Retry()
			}
		}
	})

	err := c.Visit(fmt.Sprintf("https://stocks.finance.yahoo.co.jp/etf/list/?p=%d", page))
	if err != nil {
		panic(err)
	}
}
