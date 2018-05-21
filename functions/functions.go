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
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"strconv"
)

var updater = gotgbot.NewUpdater(config.APIKEY)
var owner = config.OWNER
var sudoUser = config.SUDOUSER
var specialStatus = ""
var user2 string
var user3 string
var userid int
var username = ""

func RLeave(b ext.Bot, u gotgbot.Update) {
	splitter := strings.TrimPrefix(u.EffectiveMessage.Text, "/rleave ")

	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), fmt.Sprintf("%v", u.EffectiveUser.Id)) ||
		strings.Contains(fmt.Sprintf("%v", sudoUser), fmt.Sprintf("%v", u.EffectiveUser.Id)):
		i1, _ := strconv.Atoi(splitter)
		b.SendMessage(i1, "My owner requested me to leave this group.")
		b.LeaveChat(i1)
	}
}

func DataCheck(b ext.Bot, u gotgbot.Update) {
	user := u.EffectiveMessage.ReplyToMessage

	if user != nil {

		userid := u.EffectiveMessage.ReplyToMessage.From.Id
		username = strings.TrimPrefix(u.EffectiveMessage.ReplyToMessage.From.Username, "@")
		firstname := u.EffectiveMessage.ReplyToMessage.From.FirstName

		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		//HSET Shabier userid 167349417
		saveUsername, err := conn.Cmd("HSET", userid, "username", username).Int()
		if err != nil {
			log.Fatal(err)
		}
		saveUserID, err := conn.Cmd("HSET", username, "userid", userid).Int()
		if err != nil {
			log.Fatal(err)
		}
		//HSET 167349417 firstname Shabier
		saveFirstname, err := conn.Cmd("HSET", username, "firstname", firstname).Int()
		if err != nil {
			log.Fatal(err)
		}

		b.SendMessage(u.Message.Chat.Id, "ID:\t" + fmt.Sprintf("%v", userid) + "\nUsername: @"+
			username+ "\nFirstname:\t"+ firstname)
		fmt.Println(saveUsername, saveFirstname, saveUserID)

		//returns the userid based on username
		findUsername, err := conn.Cmd("HGET", userid, "username").Str()
		if err != nil {
			log.Fatal(err)
		}
		//returns the username based on the firstname
		findFirstname, err := conn.Cmd("HGET", username, "firstname").Str()
		if err != nil {
			log.Fatal(err)
		}

		findUserID, err := conn.Cmd("HGET", username, "userid").Str()
		if err != nil {
			log.Fatal(err)
		}

		b.SendMessage(u.Message.Chat.Id, "Found\n"+
			"UserID: "+ fmt.Sprintf("%v", findUserID)+
			"\nUsername: @"+ findUsername+
			"\nFirstname: "+ fmt.Sprintf("%v", findFirstname))

	}
}

func Checker(userid string, username string, firstname string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//HSET Shabier userid 167349417
	saveUsername, err := conn.Cmd("HSET", userid, "username", username).Int()
	if err != nil {
		log.Fatal(err)
	}
	saveUserID, err := conn.Cmd("HSET", username, "userid", userid).Int()
	if err != nil {
		log.Fatal(err)
	}
	//HSET 167349417 firstname Shabier
	saveFirstname, err := conn.Cmd("HSET", username, "firstname", firstname).Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(saveUsername, saveFirstname, saveUserID)
}

func Start(b ext.Bot, u gotgbot.Update) {
	user := u.EffectiveUser

	userid := fmt.Sprintf("%v", u.EffectiveMessage.From.Id)
	username = u.EffectiveMessage.From.Username
	firstname := u.EffectiveMessage.From.FirstName
	Checker(userid, username, firstname)

	switch {
	case u.EffectiveChat.Type == "private":
		mark := "Hey " + fmt.Sprintf("%v", user.FirstName) + ", my name is GoBot! " +
			"If you have any questions on how to use me, read /help - and then head to @GotgbotChat.\n\n" +
			"I'm a group manager bot maintained by " +
			"[this guy](https://t.me/shabier). I'm built in Golang, using " +
			"the [gotgbot](https://github.com/PaulSonOfLars/gotgbot) library " +
			"which is inspired on the [python-telegram-bot library](https://python-telegram-bot.org/), " +
			"and am fully opensource - you can find what makes me " +
			"tick [here](https://github.com/Shabier/Go_tgbot)\n\n" +
			"You can find the list of available commands with /help."
		msg := b.NewSendableMessage(u.Message.Chat.Id, mark)
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	default:
		b.ReplyMessage(u.Message.Chat.Id, "hoi", u.EffectiveMessage.MessageId)
	}
}

func Help(b ext.Bot, u gotgbot.Update) {
	switch {
	case u.EffectiveChat.Type == "private":
		b.ReplyMessage(u.Message.Chat.Id, "Not available yet.", u.EffectiveMessage.MessageId)
	default:
		b.ReplyMessage(u.Message.Chat.Id, "Not available yet.", u.Message.MessageId)
	}
}

func Pin(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id

	switch {
	case u.EffectiveMessage.ReplyToMessage != nil:
		b.PinChatMessage(chatid, u.EffectiveMessage.ReplyToMessage.MessageId)
	default:
		b.ReplyMessage(chatid, "You don't seem to be replying to a message.", u.Message.MessageId)
	}

}

func Stats(b ext.Bot, u gotgbot.Update) {
	userid := u.EffectiveUser.Id
	chatid := u.Message.Chat.Id

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), fmt.Sprintf("%v", userid)) ||
		strings.Contains(fmt.Sprintf("%v", sudoUser), fmt.Sprintf("%v", userid)):
		res, err := conn.Cmd("DBSIZE").Int()
		total := res / 2
		if err != nil {
			log.Fatal(err)
		}
		msg := b.NewSendableMessage(chatid, "You got a whopping amount of `"+
			fmt.Sprintf("%v", total)+ "` users")
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	}
}

func UnPin(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	b.UnpinChatMessage(chatid)
}

func checker() {
	owner := fmt.Sprintf("%v", config.OWNER)
	sudoUser := fmt.Sprintf("%v", config.SUDOUSER)
	switch {
	case sudoUser == user2:
		specialStatus = "This person is one of my sudo users! Nearly as powerful as my owner - so watch it."
	case owner == user2:
		specialStatus = "This person is my owner - I would never do anything against them!"
	default:
		specialStatus = ""
	}
}

func Info(b ext.Bot, u gotgbot.Update) {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if u.EffectiveMessage.ReplyToMessage != nil {
		user2 = fmt.Sprintf("%v", u.EffectiveMessage.ReplyToMessage.From.Id)
	} else {
		user2 = fmt.Sprintf("%v", u.EffectiveUser.Id)
	}

	switch {
	case u.EffectiveMessage.ReplyToMessage != nil:
		checker()
		user := u.EffectiveMessage.ReplyToMessage.From

		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" +
			fmt.Sprintf("%v", user.Id) + ")\n"

		msg := b.NewSendableMessage(u.Message.Chat.Id, mark+
			specialStatus)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	case strings.Contains(fmt.Sprintf("%v", u.EffectiveMessage), "/info @"):
		username := strings.TrimPrefix(u.EffectiveMessage.Text, "/info @")
		getUserID, err := conn.Cmd("HGET", username, "userid").Str()
		findFirstname, err := conn.Cmd("HGET", username, "firstname").Str()
		if err != nil {
			b.SendMessage(u.Message.Chat.Id, "I don't seem to have interacted with this user before - "+
				"please forward a message from them to give me control! (like a voodoo doll, I need a piece of "+
				"them to be able to execute certain commands...)")
		} else {
			user2 = getUserID
			checker()

			mark := "*User info:*\n" +
				"ID: `" + fmt.Sprintf("%v", getUserID) + "`\n" +
				"First name: " + findFirstname + "\n" +
				"Username: @" + helpers.EscapeMarkdown(username) + "\n" +
				"Permanent link for: [" + findFirstname + "](tg://user?id=" +
				fmt.Sprintf("%v", getUserID) + ")\n"

			msg := b.NewSendableMessage(u.Message.Chat.Id, mark+
				specialStatus)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.Send()
		}
	default:
		checker()
		user := u.EffectiveUser

		mark := fmt.Sprintf("*User info:*\n" +
			"ID: `" + fmt.Sprintf("%v", user.Id) + "`\n" +
			"First name: " + user.FirstName + "\n") +
			"Username: @" + helpers.EscapeMarkdown(user.Username) + "\n" +
			"Permanent link for: [" + user.FirstName + " " + user.LastName + "](tg://user?id=" +
			fmt.Sprintf("%v", user.Id) + ")\n"

		msg := b.NewSendableMessage(u.Message.Chat.Id, mark+
			specialStatus)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	}
}

func Kick(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	kicker := u.EffectiveUser
	victim := u.EffectiveMessage.ReplyToMessage.From

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
			b.KickChatMember(chatid, victim.Id)
			mark :=
				"[" + kicker.FirstName + " " + kicker.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", kicker.Id) + ")" +
					"kicked " +
					"[" + victim.FirstName + " " +
					victim.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", victim.Id) + ")"
			msg := b.NewSendableMessage(chatid, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()
		default:
			b.ReplyMessage(chatid, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(chatid, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func Ban(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	banner := u.EffectiveUser
	victim := u.EffectiveMessage.ReplyToMessage.From

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
			b.KickChatMember(chatid, u.EffectiveMessage.ReplyToMessage.From.Id)
			mark :=
				"[" + banner.FirstName + " " + banner.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", banner.Id) + ")" +
					"banned " +
					"[" + victim.FirstName + " " +
					victim.LastName + "](tg://user?id=" + fmt.Sprintf("%v", victim.Id) + ")"
			msg := b.NewSendableMessage(chatid, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()
		default:
			b.ReplyMessage(chatid, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(chatid, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func UnBan(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	banner := u.EffectiveUser
	victim := u.EffectiveMessage.ReplyToMessage.From

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
			b.UnbanChatMember(chatid, u.EffectiveMessage.ReplyToMessage.From.Id)
			mark :=
				"[" + banner.FirstName + " " + banner.LastName + "](tg://user?id=" +
					fmt.Sprintf("%v", banner.Id) + ")" +
					"unbanned " +
					"[" + victim.FirstName + " " +
					victim.LastName + "](tg://user?id=" + fmt.Sprintf("%v", victim.Id) + ")"
			msg := b.NewSendableMessage(chatid, mark)
			msg.ParseMode = parsemode.Markdown
			msg.ReplyToMessageId = u.Message.MessageId
			msg.DisableWebPreview = true
			msg.Send()
		default:
			b.ReplyMessage(chatid, "You don't seem to be referring to a user.", u.Message.MessageId)
		}
	default:
		b.ReplyMessage(chatid, "Who dis non-admin telling me what to do?", u.Message.MessageId)
	}
}

func Id(b ext.Bot, u gotgbot.Update) {
	userIdPrivate := u.EffectiveUser
	chatid := u.Message.Chat.Id

	switch {
	case u.EffectiveMessage.ReplyToMessage != nil && u.EffectiveMessage.ReplyToMessage.ForwardFrom == nil:
		forwarder := u.EffectiveMessage.ReplyToMessage.From
		mark := "[" + forwarder.FirstName + "](tg://user?id=" + fmt.Sprintf("%v", forwarder.Id) + ")" +
			"'s id is `" + fmt.Sprintf("%v", forwarder.Id) + "`."
		msg := b.NewSendableMessage(chatid, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	case u.EffectiveMessage.ReplyToMessage != nil && u.EffectiveMessage.ReplyToMessage.ForwardFrom != nil:
		forwarder := u.EffectiveMessage.ReplyToMessage.From
		orgin := u.EffectiveMessage.ReplyToMessage.ForwardFrom

		mark := "The original sender, [" + orgin.FirstName + "](tg://user?id=" + fmt.Sprintf("%v", orgin.Id) + ")" +
			", has an ID of `" + fmt.Sprintf("%v", orgin.Id) + "`. " +
			"The forwarder, [" + forwarder.FirstName + "](tg://user?id=" + fmt.Sprintf("%v", forwarder.Id) + "), " +
			"has an ID of `" + fmt.Sprintf("%v", forwarder.Id) + "`"
		msg := b.NewSendableMessage(chatid, mark)
		msg.ReplyToMessageId = u.Message.MessageId
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	case u.EffectiveChat.Type == "private":
		mark := "Your user id is `" + fmt.Sprintf("%v", userIdPrivate.Id) + "`."
		msg := b.NewSendableMessage(chatid, mark)
		msg.ReplyToMessageId = u.Message.MessageId
		msg.ParseMode = parsemode.Markdown
		msg.Send()
	default:
		count1, _ := b.GetChatMembersCount(u.Message.Chat.Id)
		count := fmt.Sprintf("%v", count1)
		mark := "This group's title is `" + u.Message.Chat.Title + "`, it has `" + count + "` members" +
			" and it's ID is: `" + fmt.Sprintf("%v", u.EffectiveChat.Id) + "`."
		msg := b.NewSendableMessage(chatid, mark)
		msg.ParseMode = parsemode.Markdown
		msg.ReplyToMessageId = u.Message.MessageId
		msg.Send()
	}
}

/**
  * doesnt work yet due lib
 */
func InviteLink(b ext.Bot, u gotgbot.Update) {
	fmt.Println("weg")
	fmt.Println(u.Message.Chat.InviteLink)
}

func Get(b ext.Bot, u gotgbot.Update) {

	//splitter := strings.TrimPrefix(u.EffectiveMessage.Text, "/get ")
	b.ReplyMessage(u.Message.Chat.Id, "SoonTM", u.EffectiveMessage.MessageId)

}

func AdminCheck(b ext.Bot, u gotgbot.Update) {
	check, _ := b.GetChatAdministrators(u.Message.Chat.Id)
	reg := regexp.MustCompile(`[^0-9]+`)
	res := reg.ReplaceAllString(fmt.Sprintf("%s\n", check), "${1}")
	user := fmt.Sprintf("%v", u.EffectiveUser.Id)
	chatid := u.Message.Chat.Id

	switch {
	case strings.Contains(fmt.Sprintf("%v", owner), user):
		b.ReplyMessage(chatid, "This person is my owner - I would never do anything against them!",
			u.EffectiveMessage.MessageId)
	case strings.Contains(fmt.Sprintf("%v", sudoUser), user):
		b.ReplyMessage(chatid, "This user is one of my sudo users, I would never do anything against them!",
			u.EffectiveMessage.MessageId)
	case strings.Contains(res, user):
		b.ReplyMessage(chatid, "This user is an admin", u.EffectiveMessage.MessageId)
	default:
		b.ReplyMessage(chatid, "You're a member", u.EffectiveMessage.MessageId)
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
