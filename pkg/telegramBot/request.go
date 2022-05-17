package telegramBot

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//request запрос в Телеграм бот
func (m *BotSetting) request() (api *tgbotapi.BotAPI, err error) {
	bot, err := tgbotapi.NewBotAPI(m.Token)
	if err != nil {
		return nil, errors.New("ошибка в подключении в Телеграм бот")
	}
	return bot, nil
}
