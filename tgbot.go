package main

import (
	"./config"
	"Go_tgbot/functions"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"log"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
)

var owner = config.OWNER
var token = config.APIKEY

var updater = gotgbot.NewUpdater(config.APIKEY)

func main() {

	log.Println("Starting gotgbot...")
	if me, err := updater.Dispatcher.Bot.GetMe(); err != nil {
		log.Printf("Error: %+v\n", err)
	} else {
		log.Printf("%+v\n", me)
	}

	updater.Dispatcher.AddHandler(handlers.NewCommand("start", functions.Start))
	updater.Dispatcher.AddHandler(handlers.NewCommand("get", functions.Get))
	updater.Dispatcher.AddHandler(handlers.NewCommand("info", functions.Info))
	updater.Dispatcher.AddHandler(handlers.NewCommand("kick", functions.Kick))
	updater.Dispatcher.AddHandler(handlers.NewCommand("ban", functions.Ban))
	updater.Dispatcher.AddHandler(handlers.NewCommand("unban", functions.UnBan))
	updater.Dispatcher.AddHandler(handlers.NewCommand("admincheck", functions.AdminCheck))
	updater.Dispatcher.AddHandler(handlers.NewCommand("hi", functions.Hi))
	updater.Dispatcher.AddHandler(handlers.NewCommand("link", functions.InviteLink))
	updater.Dispatcher.AddHandler(handlers.NewCommand("id", functions.Id))
	updater.Dispatcher.AddHandler(handlers.NewCommand("stop", functions.Stop))
	updater.Dispatcher.AddHandler(handlers.NewCommand("filter", functions.FiltersSet))
	updater.Dispatcher.AddHandler(handlers.NewRegex("(?i)hello", functions.Hi))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Sticker, functions.StickerDetect))

	updater.StartPolling()

	updater.Idle()
}
