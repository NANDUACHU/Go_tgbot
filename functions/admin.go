package functions

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"log"
	"strings"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"github.com/mediocregopher/radix.v2/redis"
	"regexp"
	"strconv"
)

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

func Pin(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id

	switch {
	case u.EffectiveMessage.ReplyToMessage != nil:
		b.PinChatMessage(chatid, u.EffectiveMessage.ReplyToMessage.MessageId)
	default:
		b.ReplyMessage(chatid, "You don't seem to be replying to a message.", u.Message.MessageId)
	}
}

func UnPin(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	b.UnpinChatMessage(chatid)
}

func Kick(b ext.Bot, u gotgbot.Update) {
	chatid := u.Message.Chat.Id
	kicker := u.EffectiveUser

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
			victim := u.EffectiveMessage.ReplyToMessage.From

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
			victim := u.EffectiveMessage.ReplyToMessage.From

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
			victim := u.EffectiveMessage.ReplyToMessage.From
			
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

/**
  * doesnt work yet due lib
 */
func InviteLink(b ext.Bot, u gotgbot.Update) {
	fmt.Println("weg")
	fmt.Println(u.Message.Chat.InviteLink)
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