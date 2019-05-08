package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewRNWBot makes and acivates r-n-w bot
func RunRNWBot(token string, dbg bool) {
	log.Printf("Start r-n-w bot")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("[DEBUG] %s", err)
	}

	bot.Debug = dbg

	log.Printf("[INFO] Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.InlineQuery != nil {
			inlineConfig, err := Process(update)
			if err != nil {
				log.Printf("[DEBUG] Error process update: %s", err)
				continue
			}
			_, err = bot.AnswerInlineQuery(*inlineConfig)
			if err != nil {
				log.Printf("[DEBUG] Cannot send msg with: %s", err)
			}
		}
	}
}
