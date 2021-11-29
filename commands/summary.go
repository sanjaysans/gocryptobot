package commands

import (
	"fmt"
	"math"

	"github.com/sanjaysans/gocryptobot/utils"
)

func GetSummary(symbol string) (string, string, error) {
	p, err := utils.GetApiCall(symbol)
	l := p.Quotes.USD.Price
	o := p.Quotes.USD.Ath_price
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		return fmt.Sprintf("%.2f", p.Quotes.USD.Price), "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), err
	} else {
		return fmt.Sprintf("%.2f", p.Quotes.USD.Price), "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), err
	}
}
