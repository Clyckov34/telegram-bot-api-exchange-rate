package telegramBot

import (
	"fmt"
	"telegram/internal/centerBank"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//App приложения запуска бота
func App(token string) error {
	var tgAPI = BotSetting{
		Token:   token,
		Timeout: 60,
	}

	bot, err := tgAPI.request()
	if err != nil {
		return err
	}

	updates := tgAPI.setting(bot)
	if err = command(bot, updates); err != nil {
		return err
	}

	return nil
}

//command команды телеграмм бота
func command(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		var currency string

		switch update.Message.Text {
		case "/start":
			msg.Text = "Добро пожаловать! Здесь Вы можете узнать курсы валют ЦБ к соотношению $USD. \n\n- Меню курсов валют: /course \n- Помощь: /help"
		case "/help":
			msg.Text = "- Меню курсов валют: /course \nПосле ввода команды откроется меню с курсами валют, выбираете нужный курс. \n- Поддреживает ввод курса на прямую\nНапример: RUB"
		case "/course":
			msg.Text = "Меню курсов валют"
			msg.ReplyMarkup = numericKeyboard
		case "AUD":
			currency = "AUD"
		case "AZN":
			currency = "AZN"
		case "GBP":
			currency = "GBP"
		case "AMD":
			currency = "AMD"
		case "BYN":
			currency = "BYN"
		case "BGN":
			currency = "BGN"
		case "BRL":
			currency = "BRL"
		case "HUF":
			currency = "HUF"
		case "HKD":
			currency = "HKD"
		case "DKK":
			currency = "DKK"
		case "USD":
			currency = "USD"
		case "EUR":
			currency = "EUR"
		case "INR":
			currency = "INR"
		case "KZT":
			currency = "KZT"
		case "CAD":
			currency = "CAD"
		case "KGS":
			currency = "KGS"
		case "CNY":
			currency = "CNY"
		case "MDL":
			currency = "MDL"
		case "NOK":
			currency = "NOK"
		case "PLN":
			currency = "PLN"
		case "RON":
			currency = "RON"
		case "XDR":
			currency = "XDR"
		case "SGD":
			currency = "SGD"
		case "TJS":
			currency = "TJS"
		case "TRY":
			currency = "TRY"
		case "TMT":
			currency = "TMT"
		case "UZS":
			currency = "UZS"
		case "UAH":
			currency = "UAH"
		case "CZK":
			currency = "CZK"
		case "SEK":
			currency = "SEK"
		case "CHF":
			currency = "CHF"
		case "ZAR":
			currency = "ZAR"
		case "KRW":
			currency = "KRW"
		case "JPY":
			currency = "JPY"
		case "RUB":
			currency = "RUB"
		default:
			msg.Text = "No Command"
		}

		resCurrency, err := centerBank.GetCourse(currency)
		if err != nil {
			msg.Text = err.Error()
		}

		if currency != "" && resCurrency != "" {
			msg.Text = fmt.Sprintf("USD / %v = %v\n", currency, resCurrency)
		}

		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
