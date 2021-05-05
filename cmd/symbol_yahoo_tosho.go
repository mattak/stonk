package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"regexp"
)

var (
	retryLimit = 10
)

func FetchYahooToshoSymbols(symbolMapChannel chan map[string]SymbolInfo) {
	symbolMap := map[string]SymbolInfo{}

	page := 1
	re := regexp.MustCompile(`code=([\d\.\w]+)`)
	c := NewColly()
	retryCount := 0

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("div#contents-body-bottom table.rankingTable tbody tr", func(y int, e *colly.HTMLElement) {
			line := []string{}
			code := ""
			info := SymbolInfo{}

			// parse symbol
			e.ForEach("td a", func(x int, e *colly.HTMLElement) {
				link := e.Attr("href")
				if len(link) > 0 {
					result := re.FindAllStringSubmatch(link, -1)
					if len(result) > 0 && len(result[0]) > 1 {
						code = result[0][1]
						info.Symbol = code
					}
				}
				line = append(line, e.Text)
			})

			// parse name
			e.ForEach("td", func(x int, e *colly.HTMLElement) {
				if x == 3 {
					// string normalize
					info.Name = NormalizeName(e.Text)
				}
			})

			// set result
			if len(code) > 0 {
				symbolMap[code] = info
			}
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

	err := c.Visit(fmt.Sprintf("https://info.finance.yahoo.co.jp/ranking/?kd=4&tm=d&vl=a&mk=1&p=%d", page))
	if err != nil {
		panic(err)
	}
}
