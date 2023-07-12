package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BOT *tgbotapi.BotAPI
var UPDATE_CHANNEL tgbotapi.UpdatesChannel
var COMMAND_MAP = make(map[string]func() string)
var CHAN_ID int64 = 0

func Start(ChannelID int64, APIKey string, offset int, timeout int) error {
	var err error
	BOT, err = tgbotapi.NewBotAPI(APIKey)
	if err != nil {
		return err
	}
	CHAN_ID = ChannelID

	BOT.Debug = false

	log.Println("Telegram Authorized on account %s", BOT.Self.UserName)

	u := tgbotapi.NewUpdate(offset)
	u.Timeout = timeout

	UPDATE_CHANNEL = BOT.GetUpdatesChan(u)

	for update := range UPDATE_CHANNEL {
		if update.ChannelPost != nil {
			function, ok := COMMAND_MAP[update.ChannelPost.Text]
			if !ok {
				msg := tgbotapi.NewMessage(CHAN_ID, "Unknown command")
				_, err = BOT.Send(msg)
				if err != nil {
					log.Println(err)
				}
				continue
			}

			msg := tgbotapi.NewMessage(CHAN_ID, function())
			_, err = BOT.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}

	}

	return nil
}

func SendMessage(msg ...interface{}) error {
	_, err := BOT.Send(tgbotapi.NewMessage(CHAN_ID, fmt.Sprint(msg...)))
	if err != nil {
		return err
	}
	return nil
}
