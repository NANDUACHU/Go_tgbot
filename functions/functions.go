package functions

import (
	"Go_tgbot/config"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"strings"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/ext/helpers"
	"regexp"
	"github.com/PaulSonOfLars/gotgbot/handlers"
)

var updater = gotgbot.NewUpdater(config.APIKEY)
var owner = config.OWNER
var sudoUser = config.SUDOUSER
var status = ""
var user2 = ""
var user3 = ""

func Start(b ext.Bot, u gotgbot.Update) {
	user := u.EffectiveUser
	switch {
	case u.EffectiveChat.Type == "private":
		mark :="Hey " + fmt.Sprintf("%v", user.FirstName) + ", my name is GoBot! " +
			"If you have any questions on how to use me, read /help - and then head to @GotgbotChat.\n\n"+
			"I'm a group manager bot maintained by I'm a group manager bot maintained by " +
			"[this guy](https://t.me/shabier). I'm built in Golang, using " +
			"the [gotgbot](https://github.com/PaulSonOfLars/gotgbot) library " +
			"which is inspired on the [python-telegram-bot library](https://python-telegram-bot.org/), " +
			"and am fully opensource - you can find what makes me " +
			"tick [here](https://github.com/Shabier/Go_tgbot)\n\n"+
			"You can find the list of available commands with /help."
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	default:
		b.SendMessage(u.Message.Chat.Id, "hoi")
	}
	}

func checker() {
	owner := fmt.Sprintf("%v", owner)
	sudoUser := fmt.Sprintf("%v", sudoUser)
	switch {
	case sudoUser == user2:
		status = "This person is one of my sudo users! Nearly as powerful as my owner - so watch it."
	case owner == user2:
		status = "This person is my owner - I would never do anything against them!"
	default:
		status = ""
	}
}

func Info(b ext.Bot, u gotgbot.Update) {

	if u.EffectiveMessage.ReplyToMessage != nil {
		user2 = fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id)
	} else {
		user2 = fmt.Sprintf("%v", u.EffectiveUser.Id)
	}

	switch {
	case u.EffectiveMessage.ReplyToMessage != nil:
		user := u.EffectiveMessage.ReplyToMessage.From
		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" +
			fmt.Sprintf("%v", user.Id) + ")\n"
		checker()
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark+
			status)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.DisableWebPreview = true
		msg.Send()
		status = ""
	default:
		user := u.EffectiveUser
		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" +
			fmt.Sprintf("%v", user.Id) + ")\n"
		checker()
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark+
			status)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.DisableWebPreview = true
		msg.Send()
	}
}

func Kick(b ext.Bot, u gotgbot.Update) {
	check, _ := b.GetChatAdministrators(u.Message.Chat.Id)
	reg := regexp.MustCompile(`[^0-9]+`)
	res := reg.ReplaceAllString(fmt.Sprintf("%s\n", check), "${1}")
	user := fmt.Sprintf("%v", u.EffectiveUser.Id)
	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(fmt.Sprintf("%v", sudoUser), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(res, user):
		switch {
		case u.EffectiveMessage.ReplyToMessage != nil:
			b.KickChatMember(u.Message.Chat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
			user := u.EffectiveUser
			mark :=
				"[" + u.EffectiveUser.FirstName + " " + u.EffectiveUser.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", user.Id) + ")" +
					"kicked " +
					"[" + u.EffectiveMessage.ReplyToMessage.From.FirstName + " " + u.EffectiveMessage.ReplyToMessage.From.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id) + ")"
			msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()

		default:
			b.ReplyMessage(u.Message.Chat.Id, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(u.Message.Chat.Id, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func Ban(b ext.Bot, u gotgbot.Update) {
	check, _ := b.GetChatAdministrators(u.Message.Chat.Id)
	reg := regexp.MustCompile(`[^0-9]+`)
	res := reg.ReplaceAllString(fmt.Sprintf("%s\n", check), "${1}")
	user := fmt.Sprintf("%v", u.EffectiveUser.Id)
	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(fmt.Sprintf("%v", sudoUser), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(res, user):
		switch {
		case u.EffectiveMessage.ReplyToMessage != nil:
			b.KickChatMember(u.Message.Chat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
			user := u.EffectiveUser
			mark :=
				"[" + u.EffectiveUser.FirstName + " " + u.EffectiveUser.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", user.Id) + ")" +
					"banned " +
					"[" + u.EffectiveMessage.ReplyToMessage.From.FirstName + " " + u.EffectiveMessage.ReplyToMessage.From.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id) + ")"
			msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()

		default:
			b.ReplyMessage(u.Message.Chat.Id, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(u.Message.Chat.Id, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func UnBan(b ext.Bot, u gotgbot.Update) {
	check, _ := b.GetChatAdministrators(u.Message.Chat.Id)
	reg := regexp.MustCompile(`[^0-9]+`)
	res := reg.ReplaceAllString(fmt.Sprintf("%s\n", check), "${1}")
	user := fmt.Sprintf("%v", u.EffectiveUser.Id)
	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(fmt.Sprintf("%v", sudoUser), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(res, user):
		switch {
		case u.EffectiveMessage.ReplyToMessage != nil:
			b.UnbanChatMember(u.Message.Chat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
			user := u.EffectiveUser
			mark :=
				"[" + u.EffectiveUser.FirstName + " " + u.EffectiveUser.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", user.Id) + ")" +
					"unbanned " +
					"[" + u.EffectiveMessage.ReplyToMessage.From.FirstName + " " + u.EffectiveMessage.ReplyToMessage.From.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id) + ")"
			msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()

		default:
			b.ReplyMessage(u.Message.Chat.Id, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(u.Message.Chat.Id, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func Id(b ext.Bot, u gotgbot.Update) {
	switch {
	case u.EffectiveMessage.ReplyToMessage != nil:
		mark := "[" + u.EffectiveMessage.ReplyToMessage.From.FirstName + "](tg://user?id=" +
			fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id) + ")" +
			"'s id is `" + fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id) + "`."
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	default:
		mark := "This group's id is `" + fmt.Sprintf("%v", u.EffectiveChat.Id) + "`."
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	}
}

func InviteLink(b ext.Bot, u gotgbot.Update) {
	fmt.Println("weg")
	fmt.Println(u.Message.Chat.InviteLink)
}

func Get(b ext.Bot, u gotgbot.Update) {

	//splitter := strings.TrimPrefix(u.EffectiveMessage.Text, "/get ")

}

func AdminCheck(b ext.Bot, u gotgbot.Update) {
	check, _ := b.GetChatAdministrators(u.Message.Chat.Id)
	reg := regexp.MustCompile(`[^0-9]+`)
	res := reg.ReplaceAllString(fmt.Sprintf("%s\n", check), "${1}")
	user := fmt.Sprintf("%v", u.EffectiveUser.Id)

	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), user):
		b.ReplyMessage(u.Message.Chat.Id, "This person is my owner - I would never do anything against them!", u.EffectiveMessage.MessageId)
	case strings.Contains(fmt.Sprintf("%v", sudoUser), user):
		b.ReplyMessage(u.Message.Chat.Id, "This user is one of my sudo users, I would never do anything against them!", u.EffectiveMessage.MessageId)
	case strings.Contains(res, user):
		b.ReplyMessage(u.Message.Chat.Id, "This user is an admin", u.EffectiveMessage.MessageId)
	default:
		b.ReplyMessage(u.Message.Chat.Id, "You're a member", u.EffectiveMessage.MessageId)
	}
}

func Stop(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Bye")
}

func Hi(b ext.Bot, u gotgbot.Update) {
	b.SendMessage(u.Message.Chat.Id, "Hello to you too!")
}

func StickerDetect(b ext.Bot, u gotgbot.Update) {
	////b.SendMessage("Nice sticker...!", u.Message.Chat.Id)
	//if _, err := u.EffectiveMessage.Delete(); err != nil {
	//	u.EffectiveMessage.ReplyMessage("Can't delete, you're in PM")
	//} else {
	//	msg := b.NewSendableMessage(u.Message.Chat.Id, "Don't you *dare* send _stickers_ here!")
	//	msg.ParseMode = parsemode.Markdown
	//	msg.Send()
	//}
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
