package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/sanjaysans/gocryptobot/utils"
)

func GetHistoric(coin_id string) (string, error) {
	p, err := utils.GetApiCall(coin_id)
	printfstr := "Current price: $ %.6f \n\n"
	printfstr += "Volume 24 hrs: %.2f \n\n"
	printfstr += "Volume change 24 hrs: %.2f %%  \n\n"
	printfstr += "Market cap: %.2f \n\n"
	printfstr += "Market cap change 24 hrs: %.2f %% \n\n"
	printfstr += "Percent change 15 mins: %.2f %% \n\n"
	printfstr += "Percent change 30 mins: %.2f %% \n\n"
	printfstr += "Percent change 1 hr: %.2f %% \n\n"
	printfstr += "Percent change 6 hrs: %.2f %% \n\n"
	printfstr += "Percent change 12 hrs: %.2f %% \n\n"
	printfstr += "Percent change 24 hrs: %.2f %% \n\n"
	printfstr += "Percent change 7 days: %.2f %% \n\n"
	printfstr += "Percent change 30 days: %.2f %% \n\n"
	printfstr += "Percent change 1 year: %.2f %% \n\n"
	printfstr += "All time high price: $ %.6f \n\n"
	printfstr += "All time high date: %s \n\n"
	printfstr += "Percentage for all time high and current price: %.2f %%"

	if err != nil {
		log.Fatalf(err.Error())
	}

	myDate, err := time.Parse("2006-01-02T15:04:15Z", p.Quotes.USD.Ath_date)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return fmt.Sprintf(printfstr, p.Quotes.USD.Price, p.Quotes.USD.Volume_24h, p.Quotes.USD.Volume_24h_change_24h, p.Quotes.USD.Market_cap, p.Quotes.USD.Market_cap_change_24h, p.Quotes.USD.Percent_change_15m, p.Quotes.USD.Percent_change_30m, p.Quotes.USD.Percent_change_1h, p.Quotes.USD.Percent_change_6h, p.Quotes.USD.Percent_change_12h, p.Quotes.USD.Percent_change_24h, p.Quotes.USD.Percent_change_7d, p.Quotes.USD.Percent_change_30d, p.Quotes.USD.Percent_change_1y, p.Quotes.USD.Ath_price, myDate.Format(time.RFC822), p.Quotes.USD.Percent_from_price_ath), err
}
