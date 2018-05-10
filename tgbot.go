package main

import (
	"./config"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"strings"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"log"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"fmt"
)

var owner = 167349417
var token = config.APIKEY

func start(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Hi")
}

func stop(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Bye")
}

func hi(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Hello to you too!")
}

func stickerDetect(b ext.Bot, u gotgbot.Update) {
	//b.SendMessage("Nice sticker...!", u.Message.Chat.Id)
	if _, err := u.EffectiveMessage.Delete(); err != nil {
		u.EffectiveMessage.ReplyMessage("Can't delete, you're in PM")
	} else {
		msg := b.NewSendableMessage(u.Message.Chat.Id, "Don't you *dare* send _stickers_ here!")
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	}
}

func text_detect(b ext.Bot, u gotgbot.Update) {
	u.EffectiveMessage.ReplyMessage("using new text func")
}

func filtersSet(b ext.Bot, u gotgbot.Update) {

	f := strings.Fields(u.Message.Text)
	if len(f) >= 3 {
		filt := f[1]
		txt := strings.Join(f[2:], " ")
		updater.Dispatcher.AddHandler(handlers.NewRegex("(?i)"+filt, func(bb ext.Bot, uu gotgbot.Update) { bb.SendMessage(uu.Message.Chat.Id, txt) }))
		u.EffectiveMessage.ReplyMessage("Added handler " + filt + "!")

	}
}

var updater = gotgbot.NewUpdater(token)

func main() {
	log.Println("Starting gotgbot...")
	if me, err := updater.Dispatcher.Bot.GetMe(); err != nil {
		log.Printf("Error: %+v\n", err)
	} else {
		log.Printf("%+v\n", me)
	}
	updater.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	updater.Dispatcher.AddHandler(handlers.NewCommand("stop", stop))
	updater.Dispatcher.AddHandler(handlers.NewCommand("test", test))
	updater.Dispatcher.AddHandler(handlers.NewCommand("filter", filtersSet))
	//updater.Dispatcher.AddHandler(Handlers.NewCommand("get", get))
	updater.Dispatcher.AddHandler(handlers.NewRegex("(?i)hello", hi))

	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Sticker, stickerDetect))
	//updater.Dispatcher.Add_handler(Handlers.NewMessage(Filters.UserID(12345), sticker_detect))
	//updater.Dispatcher.Add_handler(Handlers.NewMessage(Filters.Text, text_detect))

	updater.StartPolling()

	updater.Idle()
}

func test(b ext.Bot, u gotgbot.Update) {
	msg := b.NewSendableMessage(u.Message.Chat.Id, "This is a test with _markdown_")
	msg.ParseMode = parsemode.Markdown
	msg.ReplyToMessageId = u.Message.MessageId
	_, err := msg.Send()
	if err != nil {
		log.Println(err)
		log.Println("Did not send")
	}

	u.EffectiveMessage.ReplyMessage("Hey there")
}

func get(b ext.Bot, u gotgbot.Update) {
	chat, err := b.GetChat(u.EffectiveMessage.Chat.Id)
	if err != nil {
		log.Println("failed to get chat!")
		return
	}
	u.EffectiveMessage.ReplyMessage(fmt.Sprintf("%+v", chat))
	u.EffectiveMessage.ReplyMessage(fmt.Sprintf("%+v", u.EffectiveMessage))
}
