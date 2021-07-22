package price

import (
	"encoding/json"
	"fmt"
	"github.com/piquette/finance-go/datetime"
	"math/big"
)

type PriceCandle struct {
	Date   datetime.Datetime `json:"date"`
	Open   *big.Float        `json:"open"`
	Close  *big.Float        `json:"close"`
	High   *big.Float        `json:"high"`
	Low    *big.Float        `json:"low"`
	Volume int               `json:"volume"`
}

type PriceCandles []PriceCandle

func (pc PriceCandle) ToTsvHeader() string {
	return "date\topen\tclose\thigh\tlow\tvolume"
}

func (pc PriceCandle) ToTsv() string {
	return fmt.Sprintf(
		"%04d-%02d-%02d\t%f\t%f\t%f\t%f\t%d",
		pc.Date.Year, pc.Date.Month, pc.Date.Day,
		pc.Open,
		pc.Close,
		pc.High,
		pc.Low,
		pc.Volume,
	)
}

func (pcs PriceCandles) ToTsv() []string {
	lines := []string{PriceCandle{}.ToTsvHeader()}
	for _, line := range pcs {
		lines = append(lines, line.ToTsv())
	}
	return lines
}

func (pc PriceCandle) ToJson() string {
	jsonValue, _ := json.Marshal(pc)
	return string(jsonValue)
}
