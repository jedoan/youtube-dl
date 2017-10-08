package handler

import (
	"gopkg.in/telegram-bot-api.v4"
)


type YouBot struct {
	Update tgbotapi.Update
	Bot    *tgbotapi.BotAPI
}