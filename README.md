# Go telegram bot

The purpose of this bot is to become a full-featured group management bot thats simple to use. Join [the support group](https://t.me/GotgbotChat) if you need support or have feature requests.

# Setup
To setup this bot you have to fork this repo. After that use sample_config.go as a template to make `config.go`. After that [install go](https://golang.org/dl/) and execute `go get github.com/PaulSonOfLars/gotgbot`

This bot makes use of [Redis](https://redis.io/) database, find a guide that works for your OS and have it installed. After, get the [radix.v2](https://github.com/mediocregopher/radix.v2/) library by executing `go get github.com/mediocregopher/radix.v2/redis` in the terminal

# Starting the bot

Once you've forked, installed Go and setup the `config.go`, open the terminal and navigate to $HOME/go/src/projectfolder/ and execute `go run tgbot.go` 
