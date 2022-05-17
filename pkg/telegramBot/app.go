package telegramBot

import (
	"fmt"
	"telegram/internal/centerBank"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	welcom          = fmt.Sprintln("Добро пожаловать! Здесь Вы можете узнать курсы валют ЦБ к соотношению $USD. \n\n- Меню курсов валют: /course \n- Помощь: /help")
	descriptionHelp = fmt.Sprintln("- Меню курсов валют: /course \nПосле ввода команды откроется меню с курсами валют, выбираете нужный курс. \n- Поддреживает ввод курса на прямую\nНапример: RUB")
	currencyList    = fmt.Sprintln(`
	- 'AUD' - Австралийский доллар
	- 'AZN' - Азербайджанский манат
	- 'GBP' - Фунт стерлингов
	- 'AMD' - Армянский драм
	- 'BYN' - Белорусский рубль
	- 'BGN' - Болгарский Лев
	- 'BRL' - Бразильский реал
	- 'HUF' - Венгерский форинт
	- 'HKD' - Гонконгский доллар
	- 'DKK' - Датская крона
	- 'USD' - Доллар США
	- 'EUR' - Евро
	- 'INR' - Индийская рупия
	- 'KZT' - Казахстанский тенге
	- 'CAD' - Канадский доллар
	- 'KGS' - Киргизский сом
	- 'CNY' - Китайский юань
	- 'MDL' - Молдавский лей
	- 'NOK' - Норвежская крона
	- 'PLN' - Польский злотый
	- 'RON' - Румынский лей
	- 'XDR' - Международный Валютный Фонд
	- 'SGD' - Сингапурский доллар
	- 'TJS' - Таджикский сомони
	- 'TRY' - Турецкая лира
	- 'TMT' - Туркменский манат
	- 'UZS' - Узбекский сум
	- 'UAH' - Украинская гривна
	- 'CZK' - Чешская крона
	- 'SEK' - Шведская крона
	- 'CHF' - Швейцарский франк
	- 'ZAR' - Южноафриканский рэнд
	- 'KRW' - Южнокорейская вона
	- 'JPY' - Японская иена
	- 'RUB' - Российский рубль`)
)

//App приложения запуска бота
func App(token string) (err error) {
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
			msg.Text = fmt.Sprintf("USD / %v = %v\n", currencie, resCurrency)
		}

		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
