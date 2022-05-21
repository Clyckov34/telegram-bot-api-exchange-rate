package telegramBot

import (
	"fmt"
	"telegram/internal/centerBank"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//App приложения запуска бота
func App(parameters BotSetting) (err error) {
	var tgAPI = BotSetting{
		Token:   parameters.Token,
		Timeout: parameters.Timeout,
		Debug:   parameters.Debug,
	}

	bot, err := tgAPI.request()
	if err != nil {
		return err
	}

	updates := tgAPI.setting(bot)

	if tgAPI.Debug {
		tgAPI.debug(bot)
	}

	if err = command(bot, updates); err != nil {
		return err
	}

	return nil
}

//command команды телеграмм бота
func command(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) (err error) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		var currencie string

		switch update.Message.Text {
		case "/start":
			msg.Text = welcom
		case "/help":
			msg.Text = descriptionHelp + "\n" + currencyList
		case "/course":
			msg.Text = "Меню курсов валют"
			msg.ReplyMarkup = numericKeyboard
		case "AUD":
			currencie = "AUD"
		case "AZN":
			currencie = "AZN"
		case "GBP":
			currencie = "GBP"
		case "AMD":
			currencie = "AMD"
		case "BYN":
			currencie = "BYN"
		case "BGN":
			currencie = "BGN"
		case "BRL":
			currencie = "BRL"
		case "HUF":
			currencie = "HUF"
		case "HKD":
			currencie = "HKD"
		case "DKK":
			currencie = "DKK"
		case "USD":
			currencie = "USD"
		case "EUR":
			currencie = "EUR"
		case "INR":
			currencie = "INR"
		case "KZT":
			currencie = "KZT"
		case "CAD":
			currencie = "CAD"
		case "KGS":
			currencie = "KGS"
		case "CNY":
			currencie = "CNY"
		case "MDL":
			currencie = "MDL"
		case "NOK":
			currencie = "NOK"
		case "PLN":
			currencie = "PLN"
		case "RON":
			currencie = "RON"
		case "XDR":
			currencie = "XDR"
		case "SGD":
			currencie = "SGD"
		case "TJS":
			currencie = "TJS"
		case "TRY":
			currencie = "TRY"
		case "TMT":
			currencie = "TMT"
		case "UZS":
			currencie = "UZS"
		case "UAH":
			currencie = "UAH"
		case "CZK":
			currencie = "CZK"
		case "SEK":
			currencie = "SEK"
		case "CHF":
			currencie = "CHF"
		case "ZAR":
			currencie = "ZAR"
		case "KRW":
			currencie = "KRW"
		case "JPY":
			currencie = "JPY"
		case "RUB":
			currencie = "RUB"
		default:
			msg.Text = "No Command"
		}

		resCurrency, err := centerBank.GetCourse(currencie)
		if err != nil {
			msg.Text = err.Error()
		}

		if currencie != "" && resCurrency != "" {
			msg.Text = fmt.Sprintf("1 USD =  %v %v\n", resCurrency, currencie)
		}

		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
