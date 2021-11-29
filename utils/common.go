package utils

import (
	"log"

	"github.com/sanjaysans/gocryptobot/model"
)

var Coins []model.Coin

func GetCoinId(symbol1 string) (string, string) {
	coin_id := ""
	symbol := ""

	for _, v := range Coins {
		if v.Symbol == symbol1 {
			coin_id = v.Coin_id
			symbol = v.Symbol
			break
		}
	}

	if coin_id == "" {
		coin_id = "btc-bitcoin"
		symbol = "BTC"
	}
	log.Println(coin_id, symbol)
	return coin_id, symbol
}

func GetCoins() {
	var err error

	Coins, err = GetCoinListApiCall()
	if err != nil {
		log.Fatal(err)
		return
	}
}
