package model

type Price struct {
	Name   string `json:"name"`
	Quotes struct {
		USD struct {
			Price                  float64 `json:"price"`
			Volume_24h             float64 `json:"volume_24h"`
			Volume_24h_change_24h  float64 `json:"volume_24h_change_24h"`
			Market_cap             float64 `json:"market_cap"`
			Market_cap_change_24h  float64 `json:"market_cap_change_24h"`
			Percent_change_15m     float64 `json:"percent_change_15m"`
			Percent_change_30m     float64 `json:"percent_change_30m"`
			Percent_change_1h      float64 `json:"percent_change_1h"`
			Percent_change_6h      float64 `json:"percent_change_6h"`
			Percent_change_12h     float64 `json:"percent_change_12h"`
			Percent_change_24h     float64 `json:"percent_change_24h"`
			Percent_change_7d      float64 `json:"percent_change_7d"`
			Percent_change_30d     float64 `json:"percent_change_30d"`
			Percent_change_1y      float64 `json:"percent_change_1y"`
			Ath_price              float64 `json:"ath_price"`
			Ath_date               string  `json:"ath_date"`
			Percent_from_price_ath float64 `json:"percent_from_price_ath"`
		} `json:"USD"`
	} `json:"quotes"`
}
