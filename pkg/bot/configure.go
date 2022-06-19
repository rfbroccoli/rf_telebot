package bot

import (
	tele "gopkg.in/telebot.v3"
)

func ConfigureCommands(bot *tele.Bot) {
	bot.Handle("/register", registerHandler)
	bot.Handle("/when", whenHandler)
	bot.Handle("/help", helpHandler)
	bot.Handle("/attendance", attendanceHandler)
	bot.Handle(tele.OnText, textHandler)
}
