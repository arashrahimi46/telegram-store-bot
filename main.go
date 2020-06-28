package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	_ "strings"
	"testBot/flow"
	"time"
)


func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  "1225787378:AAHclDu-sBmf2eimfStH1BtKk-EH8xX_N8k",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}
	botFlow := flow.GetBotFlow(bot)
	bot.Handle("/start", func(m *tb.Message) {
		log.Println("starting the botFlow for", m.Sender.Recipient())
		if err := botFlow.Start(m.Sender, "به فروشگاه تستی من خوش آمدید برای انتخاب محصول خود نام یک محصول را وارد نمایید"); err != nil {
			log.Println("failed to start the conversation", err)
		}
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		botFlow.Process(m)
	})
	log.Println("starting...", bot.Me.Username)
	bot.Start()
}

