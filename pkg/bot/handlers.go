package bot

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/pwhb/rf_telebot/pkg/models"
	tele "gopkg.in/telebot.v3"
)

var helloArray = [4]string{"hello", "hi", "hey", "ဟယ်လို"}

func startHandler(ctx tele.Context) error {
	sender := ctx.Message().Sender
	hello := fmt.Sprintf("hello %v %v", sender.FirstName, sender.LastName)
	ctx.Send(hello)
	ctx.Send("ဘယ်လိုကူညီပေးရမလဲ")
	return helpHandler(ctx)
}

func registerHandler(ctx tele.Context) error {
	if ctx.Chat().Type != tele.ChatPrivate {
		return ctx.Send(noRegisterString)
	}
	msg := ctx.Message()
	sender := msg.Sender
	oldStudent, _ := models.GetOneStudent(sender.ID)
	if oldStudent != nil {
		ctx.Send("☑ register လုပ်ပြီးသားဖြစ်ပါတယ်။ ထပ်လုပ်ရန်မလိုပါ")
		response := models.GetInfoString(oldStudent)
		return ctx.Send(response)
	}

	student := models.Student{
		TelegramId:       sender.ID,
		FirstName:        sender.FirstName,
		LastName:         sender.LastName,
		TelegramUsername: sender.Username,
		BatchName:        batch17.BatchName,
		ClassName:        batch17.ClassName,
	}
	// response := fmt.Sprintf("Student: %v", student)
	models.CreateNewStudent(&student)
	ctx.Send("✅ register လုပ်လိုက်ပါပြီ။")
	response := models.GetInfoString(&student)
	return ctx.Send(response)
}

func whenHandler(ctx tele.Context) error {
	now := time.Now()
	today := now.Weekday()
	log.Println(now.Clock())
	switch today {
	case time.Saturday:
	case time.Sunday:
		return ctx.Send("Tuesday ညနေ 6:45 P.M.")
	case time.Monday:
	case time.Thursday:
		return ctx.Send("မနက်ဖြန် ညနေ 6:45 P.M.")
	case time.Wednesday:
		return ctx.Send("Friday ညနေ 6:45 P.M.")
	case time.Tuesday:
	case time.Friday:
		return ctx.Send("ဒီနေ့ ညနေ 6:45 P.M.")
	}
	return ctx.Send("")
}

func helpHandler(ctx tele.Context) error {
	chatType := ctx.Chat().Type
	// log.Printf("private == %v : %v", private, private == tele.ChatPrivate)
	if chatType == tele.ChatPrivate {
		return ctx.Send(privateString)
	}
	return ctx.Send(groupString)
}

func attendanceHandler(ctx tele.Context) error {
	return ctx.Send(attendanceString)
}

func textHandler(ctx tele.Context) error {
	msg := ctx.Message()
	text := msg.Text
	for _, keyword := range helloArray {
		match, _ := regexp.MatchString(keyword, text)
		if match {
			response := fmt.Sprintf("%v %v %v", keyword, msg.Sender.FirstName, msg.Sender.LastName)
			return ctx.Send(response)
		}
	}
	return ctx.Send("sorry နားမလည်ပါ")
}
