package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron"
	"github.com/sanjaysans/gocryptobot/config"
	"github.com/sanjaysans/gocryptobot/handler"
	"github.com/sanjaysans/gocryptobot/utils"

	tb "gopkg.in/tucnak/telebot.v2"
)

var one bool = true

func main() {
	log.Println("port: " + os.Getenv("PORT"))
	config.LoadConfig()
	http.HandleFunc("/", homePage)

	// To keep heroku deployment alive more than 30 minutes
	log.Println("Create new cron")
	c := cron.New()
	c.AddFunc("15 * * * *", func() {
		log.Println("[Job 1]Every minute job")
		utils.GetHerokuApp()
	})
	log.Println("Start cron")
	c.Start()

	log.Println("port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	if one {
		b, err := tb.NewBot(tb.Settings{
			Token:  config.LoadConfig().Token,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal(err)
			return
		}

		utils.GetCoins()

		for k, v := range handler.LoadHandler(b) {
			b.Handle(k, v)
			log.Println(k + "✅ Loaded!")
		}
		one = false
		b.Start()
	}
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
