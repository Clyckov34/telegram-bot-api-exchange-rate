package main

import (
	"flag"
	"log"
	"telegram/pkg/telegramBot"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Token Bot Telegram")
}

func main() {
	flag.Parse()
	if err := telegramBot.App(token); err != nil {
		log.Println(err)
	}
}
