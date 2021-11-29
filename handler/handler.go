package handler

import (
	"log"
	"strings"

	"github.com/sanjaysans/gocryptobot/commands"
	"github.com/sanjaysans/gocryptobot/utils"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/start"] = func(m *tb.Message) {
		b.Send(m.Chat, `Welcome to the crypto bot. Find the below commands that help you get benifits from the bot. 
		1. /start - Reply's with the welcome text.
		2. /help - Reply's with list of commands that can be executed.
		3. /price COIN - Returns the current price of the coin (/price btc)
		4. /historic COIN - Returns the current price & historic value of the coin (/historic btc)`)
	}

	commandMap["/price"] = func(m *tb.Message) {
		log.Println(">>>>>>>>>>>>>>>>>>>> Start /price " + m.Payload)
		coin_id, symbol := utils.GetCoinId(strings.ToUpper(m.Payload))
		res, _ := commands.GetPrice(coin_id)
		log.Println("response " + res)
		log.Println(">>>>>>>>>>>>>>>>>>>> End /price " + m.Payload)
		b.Send(m.Chat, symbol+"'s Current price is: $ "+res+" USD.")
	}

	commandMap["/historic"] = func(m *tb.Message) {
		log.Println(">>>>>>>>>>>>>>>>>>>> Start /historic " + m.Payload)
		coin_id, symbol := utils.GetCoinId(strings.ToUpper(m.Payload))
		res, _ := commands.GetHistoric(coin_id)
		log.Println("response " + res)
		log.Println(">>>>>>>>>>>>>>>>>>>> End /historic " + m.Payload)
		b.Send(m.Chat, symbol+"'s Historic price values are: \n\n"+res)
	}

	commandMap["/help"] = func(m *tb.Message) {
		b.Send(m.Chat, `Welcome to the crypto bot. Find the below commands that help you get benifits from the bot. 
		1. /start - Reply's with the welcome text.
		2. /help - Reply's with list of commands that can be executed.
		3. /price COIN - Returns the current price of the coin (/price btc)
		4. /historic COIN - Returns the current price & historic value of the coin (/historic btc)`)
	}

	return commandMap
}
