package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"time"
)

var (
	retryLimit = 10
)

func FetchYahooTokyoSymbols(tickerMapChannel chan map[string]bool) {
	codes := map[string]bool{}

	page := 1
	re := regexp.MustCompile(`code=([\d\.\w]+)`)
	c := colly.NewCollector()
	c.SetRequestTimeout(time.Duration(time.Second * 60))
	retryCount := 0

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