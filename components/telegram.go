package components

import (
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var telegramBot *tb.Bot

// NewTelegramBot returns a new tg bot
func NewTelegramBot() *tb.Bot {
	if telegramBot != nil {
		return telegramBot
	}
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_KEY"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	b.Handle("/myid", func(m *tb.Message) {
		b.Send(m.Sender, m.Sender.Recipient())
	})
	return b
}

// GetUser returns the user object identified by @param:recipentID
func GetUser(recipientID string) tb.User {
	i, err := strconv.Atoi(recipientID)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tb.User{ID: i}
}
