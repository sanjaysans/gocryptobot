package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sanjaysans/gocryptobot/config"
	"github.com/sanjaysans/gocryptobot/handler"
	"github.com/sanjaysans/gocryptobot/utils"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  config.LoadConfig().Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	utils.GetCoins()

	if err != nil {
		log.Fatal(err)
		return
	}

	for k, v := range handler.LoadHandler(b) {
		b.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	b.Start()

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
