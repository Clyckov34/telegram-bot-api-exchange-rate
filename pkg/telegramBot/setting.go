package telegramBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//setting Настройка Бот телеграм
func (m *BotSetting) setting(tg *tgbotapi.BotAPI) (channel tgbotapi.UpdatesChannel) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = m.Timeout
	return tg.GetUpdatesChan(u)
}

//debug Дебаг. проверка
func (m *BotSetting) debug(tg *tgbotapi.BotAPI) {
	tg.Debug = true
}
