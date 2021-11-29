package commands

import (
	"fmt"

	"github.com/sanjaysans/gocryptobot/utils"
)

func GetPrice(coin_id string) (string, error) {
	p, err := utils.GetApiCall(coin_id)
	return fmt.Sprintf("%.6f", p.Quotes.USD.Price), err
}
