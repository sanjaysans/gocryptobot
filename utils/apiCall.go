package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sanjaysans/gocryptobot/model"
)

func GetApiCall(coin_id string) (*model.Price, error) {
	resp, err := http.Get("http://api.coinpaprika.com/v1/tickers/" + coin_id)
	p := &model.Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	log.Println(p.Name, p.Quotes.USD.Price)
	return p, err
}

func GetCoinListApiCall() ([]model.Coin, error) {
	resp, err := http.Get("https://api.coinpaprika.com/v1/coins")
	p := &[]model.Coin{}

	if err != nil {
		return *p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return *p, err
}

func GetHerokuApp() {
	resp, err := http.Get("http://gocryptobot.herokuapp.com/awake")
	log.Println("resp: ", resp, "\nerr: ", err)
}
