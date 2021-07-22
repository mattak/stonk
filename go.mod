module github.com/mattak/stonk

go 1.14

require (
	github.com/Finnhub-Stock-API/finnhub-go v1.2.1
	github.com/PuerkitoBio/goquery v1.7.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.6 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/ktnyt/go-moji v1.0.0
	github.com/piquette/finance-go v1.0.0
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/temoto/robotstxt v1.1.2 // indirect
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	github.com/stretchr/testify v1.7.0
)

replace github.com/mattak/stonk/cmd => ./cmd

replace github.com/mattak/stonk/pkg/price => ./pkg/price

replace github.com/mattak/stonk/pkg/symbol => ./pkg/symbol

replace github.com/mattak/stonk/pkg/util => ./pkg/util
