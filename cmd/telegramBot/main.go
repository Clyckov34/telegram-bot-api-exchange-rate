package main

import (
	"flag"
	"log"
	"os"
	"telegram/internal/check"
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

	if check.Flag(parameters.Token) != nil {
		os.Exit(1)
	} else {
		if err := telegramBot.App(parameters); err != nil {
			log.Panicln(err)
		}
	}
}
