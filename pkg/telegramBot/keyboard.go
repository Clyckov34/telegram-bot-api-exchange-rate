package telegramBot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("AUD"),
		tgbotapi.NewKeyboardButton("AZN"),
		tgbotapi.NewKeyboardButton("GBP"),
		tgbotapi.NewKeyboardButton("AMD"),
		tgbotapi.NewKeyboardButton("BYN"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("BGN"),
		tgbotapi.NewKeyboardButton("BRL"),
		tgbotapi.NewKeyboardButton("HUF"),
		tgbotapi.NewKeyboardButton("HKD"),
		tgbotapi.NewKeyboardButton("DKK"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("USD"),
		tgbotapi.NewKeyboardButton("EUR"),
		tgbotapi.NewKeyboardButton("INR"),
		tgbotapi.NewKeyboardButton("KZT"),
		tgbotapi.NewKeyboardButton("CAD"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("KGS"),
		tgbotapi.NewKeyboardButton("CNY"),
		tgbotapi.NewKeyboardButton("MDL"),
		tgbotapi.NewKeyboardButton("NOK"),
		tgbotapi.NewKeyboardButton("PLN"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("RON"),
		tgbotapi.NewKeyboardButton("XDR"),
		tgbotapi.NewKeyboardButton("SGD"),
		tgbotapi.NewKeyboardButton("TJS"),
		tgbotapi.NewKeyboardButton("TRY"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("TMT"),
		tgbotapi.NewKeyboardButton("UZS"),
		tgbotapi.NewKeyboardButton("UAH"),
		tgbotapi.NewKeyboardButton("CZK"),
		tgbotapi.NewKeyboardButton("SEK"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("CHF"),
		tgbotapi.NewKeyboardButton("ZAR"),
		tgbotapi.NewKeyboardButton("KRW"),
		tgbotapi.NewKeyboardButton("JPY"),
		tgbotapi.NewKeyboardButton("RUB"),
	),
)
