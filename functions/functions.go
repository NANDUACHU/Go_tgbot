package functions

import (
	"Go_tgbot/config"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"strings"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/ext/helpers"
)

var updater = gotgbot.NewUpdater(config.APIKEY)

func Start(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Hi")

}

func Info(b ext.Bot, u gotgbot.Update) {
	if u.EffectiveMessage.ReplyToMessage != nil {
		user := u.EffectiveMessage.ReplyToMessage.From
		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent user link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" + fmt.Sprintf("%v", user.Id) + ")\n"
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.DisableWebPreview = true
		msg.Send()
	} else {
		user := u.EffectiveUser
		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent user link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" + fmt.Sprintf("%v", user.Id) + ")\n"
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.DisableWebPreview = true
		msg.Send()
	}
}

func Kick(b ext.Bot, u gotgbot.Update) {
	if u.EffectiveMessage.ReplyToMessage!= nil{
		b.KickChatMember(u.Message.Chat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
	}else{
		b.SendMessage(u.Message.Chat.Id, "You dont seem to reply to someone.")
	}
}

func Id(b ext.Bot, u gotgbot.Update) {
	mark := "This group's id is `" + fmt.Sprintf("%v", u.EffectiveChat.Id) + "`."
	msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
	msg.ParseMode = parsemode.Markdown
	msg.ReplyToMessageId = u.Message.MessageId
	msg.Send()
}

func InviteLink(b ext.Bot, u gotgbot.Update) {

	//b.ReplyMessage(u.Message.Chat.Id, u.EffectiveChat.InviteLink, u.Message.MessageId)

}

func Get(b ext.Bot, u gotgbot.Update) {

	//splitter := strings.TrimPrefix(u.EffectiveMessage.Text, "/get ")

}

func Stop(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Bye")
}

func Hi(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Hello to you too!")
}

func StickerDetect(b ext.Bot, u gotgbot.Update) {
	//b.SendMessage("Nice sticker...!", u.Message.Chat.Id)
	if _, err := u.EffectiveMessage.Delete(); err != nil {
		u.EffectiveMessage.ReplyMessage("Can't delete, you're in PM")
	} else {
		msg := b.NewSendableMessage(u.Message.Chat.Id, "Don't you *dare* send _stickers_ here!")
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	}
}

func Text_detect(b ext.Bot, u gotgbot.Update) {
	u.EffectiveMessage.ReplyMessage("using new text func")
}

func FiltersSet(b ext.Bot, u gotgbot.Update) {

	f := strings.Fields(u.Message.Text)
	if len(f) >= 3 {
		filt := f[1]
		txt := strings.Join(f[2:], " ")
		updater.Dispatcher.AddHandler(handlers.NewRegex("(?i)"+filt, func(bb ext.Bot, uu gotgbot.Update) { bb.SendMessage(uu.Message.Chat.Id, txt) }))
		u.EffectiveMessage.ReplyMessage("Added handler " + filt + "!")

	}
}
