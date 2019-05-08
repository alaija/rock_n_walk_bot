package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/alaija/rock_n_walk_bot/bot/rnw"
)

func Process(update tgbotapi.Update) (*tgbotapi.InlineConfig, error) {
	userId := update.InlineQuery.From.ID
	if userId != 34776086 && userId != 111212377 {
		return &tgbotapi.InlineConfig{}, nil
	}
	txt := ""
	ds := rnw.Seatup(update.InlineQuery.Query)
	for _, d := range *ds {
		txt = txt + d.Description()
	}

	artical := createArticle(txt)
	inlineConfig := createConfig(update.InlineQuery.ID, artical)

	return &inlineConfig, nil
}

func createArticle(text string) tgbotapi.InlineQueryResultArticle {
	return tgbotapi.InlineQueryResultArticle{
		Type:        "article",
		ID:          "only result",
		Title:       "Seat up it?",
		Description: text,
		InputMessageContent: tgbotapi.InputTextMessageContent{
			Text: text,
		},
	}
}

func createConfig(queryId string, artical tgbotapi.InlineQueryResultArticle) tgbotapi.InlineConfig {
	return tgbotapi.InlineConfig{
		InlineQueryID: queryId,
		Results:       castToInterfaceSlice([]tgbotapi.InlineQueryResultArticle{artical}),
		CacheTime:     0,
	}
} 
func castToInterfaceSlice(iqra []tgbotapi.InlineQueryResultArticle) []interface{} {
	s := make([]interface{}, len(iqra))
	for i, v := range iqra {
		s[i] = v
	}

	return s
}
