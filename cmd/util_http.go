package cmd

import (
	"github.com/gocolly/colly"
	"net/http"
	"time"
)

func NewColly() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"),
	)
	c.SetRequestTimeout(time.Duration(time.Second * 60))
	return c
}

func DoHttpGetRequest(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("UserAgent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")


	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return res
}
