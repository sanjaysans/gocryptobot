<p align="center">
  <h1 align="center">CRYPTO-TELEGRAM-BOT</h1>
  <p align="center">
    ·
    <a href="https://github.com/sanjaysans/gocryptobot/issues">Report Bug</a>
    ·
    <a href="https://github.com/sanjaysans/gocryptobot/issues">Request Feature</a>
    ·
  </p>
</p>

<p align="center">
  <a href="https://github.com/sanjaysans/gocryptobot/graphs/contributors"><img src="	https://img.shields.io/github/issues/sanjaysans/gocryptobot"></a>
  <a href="https://github.com/sanjaysans/gocryptobot/blob/master/LICENSE"><img src="https://img.shields.io/github/license/sanjaysans/gocryptobot"></a>
  <a href="https://github.com/sanjaysans/gocryptobot/network/members"><img src="https://img.shields.io/github/forks/sanjaysans/gocryptobot"></a>
  <a href="https://img.shields.io/github/stars/sanjaysans/gocryptobot"><img src="https://img.shields.io/github/stars/sanjaysans/gocryptobot"></a>
</p>

Gocryptobot is a simple Telegram bot to get historical prices for any crypto currency at your finger tips. It's written in Golang with Telebot's Framework

## :star: Getting started

### What you will need:

- You are going to need a computer or server where to host the bot.
- Git
- Golang v1.13
- A device with Telegram

### :computer: Installation

Open a Terminal and copy these commands (Linux & Mac devices):

```bash
cd ~
git clone https://github.com/sanjaysans/gocryptobot.git
cd ./gocryptobot
mv .env.example .env
go get github.com/sanjaysans/gocryptobot
go run main.go
```

##### Warning: 
This won't work unless you replace the **REPLACE_WITH_TOKEN** on the .env file with the Token granted by @BotFather

### :white_check_mark: Create new bot to your Telegram Channel

Open [@BotFather](https://telegram.me/botfather) on telegram and create a new bot with it's __/newbot__ command.

Assign it a name. This name won't be the one that is shown on each message, so you can name it whatever you want.

@BotFather will grant you a Token. This token is the one that will replace the **REPLACE_WITH_TOKEN** on the .env.example file on this repository. (Don't forget to rename that file to .env)

You can also play a little bit more with @BotFather. For example you can use the __/setcommands__ to define the uses your bot has on the '/' icon:

```
start - Gets welcome text and commands list
price - Gets current price for any crypto currency (/price eth)
historic - Gets historic price details for any crypto currency (/historic eth)
help - Gets list of available commands
```

## :battery: Usage

Once the bot is running and added to your Telegram Group, you can use any of the following commands:

```
    * \start - Gets welcome text and commands list
    * \price - Gets current price for any crypto currency (/price eth)
    * \historic - Gets historic price details for any crypto currency (/historic eth)
    * \help - Gets list of available commands
```

## :building_construction: Contribution Guidelines:

-   **_fork_** and **_clone_** this repository
-   Make a new branch using `git checkout -b change/username`
-   Commit the desired changes to that branch
-   Sign off your commits using `git commit -s -m w/signoff`
-   Push your changes to the branch and open a pull request
