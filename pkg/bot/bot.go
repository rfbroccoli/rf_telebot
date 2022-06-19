package bot

import (
	"log"
	"time"

	"github.com/pwhb/rf_telebot/pkg/config"

	tele "gopkg.in/telebot.v3"
)

func StartBot() {
	token, err := config.GetToken()
	if err != nil {
		log.Println("No Token")
		return
	}
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	ConfigureCommands(b)

	log.Println("bot has started")
	b.Start()
}
