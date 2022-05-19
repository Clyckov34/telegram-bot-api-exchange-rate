package main

import (
	"flag"
	"log"
	"telegram/pkg/telegramBot"
)

var parameters = telegramBot.BotSetting{}

func init() {
	flag.StringVar(&parameters.Token, "t", "", "Token. Bot API Telegram")
	flag.IntVar(&parameters.Timeout, "to", 30, "Timeout. Интервал времяни обновления запросов")
	flag.BoolVar(&parameters.Debug, "d", false, "Debug, проверка логов")
}

func main() {
	flag.Parse()
	if err := telegramBot.App(parameters); err != nil {
		log.Println(err)
	}
}
