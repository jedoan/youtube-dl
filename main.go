package main

import (
	"fmt"
	"log"
	"os/exec"

	"gopkg.in/telegram-bot-api.v4"

	"github.com/jedoan/youtube-dl/handler"

	//"io/ioutil"
	"path/filepath"
	"strings"
	//"io/ioutil"
	"os"
)

var (
	youBot handler.YouBot
	err    error
)

func getMp3File(url string, youBot handler.YouBot) {
	message := "Идет обработка и конвертация файла, пожалуйста подождите"
	msg := tgbotapi.NewMessage(youBot.Update.Message.Chat.ID, message)
	youBot.Bot.Send(msg)
	cmd := fmt.Sprintf("youtube-dl -x --audio-format='mp3' %s -o mp3/yourFile.mp3 ", url)
	c := exec.Command("bash", "-c", cmd)
	err := c.Run()
	if err != nil {
		log.Print(err)
	}
	files, _ := filepath.Glob("mp3/*.mp3")
	//files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		msg := tgbotapi.NewAudioUpload(youBot.Update.Message.Chat.ID, file)
		youBot.Bot.Send(msg)
		os.Remove(file)

	}
}

func main() {

	youBot.Bot, err = tgbotapi.NewBotAPI("427966213:AAEbSL9GDLN8P08D-TiEHx3hk_vA7pNdOIw")
	if err != nil {
		log.Panicf("Error connecting to Telegram: %v", err)
	}

	log.Printf("Authorized on account %s", youBot.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := youBot.Bot.GetUpdatesChan(u)

	if err != nil {
		log.Panicf("Error getting updates channel %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		youBot.Update = update
		command := strings.Split(youBot.Update.Message.Text, " ")[0]
		//chatStateMachine(&youBot)

		if command == "/d" {
			url := strings.Split(youBot.Update.Message.Text, " ")[1]
			getMp3File(url, youBot)
		}

		if command == "/start" || command == "/START" {
			message := `
			Бот извлекает аудиодорожки с видео на сайте youtube.com
		   для того, чтоб получить mp3 фаил аудио-дорожки необходимо ввести
			/d [ссылку на видео]
			например:  '/d https://www.youtube.com/watch?v=J1usv2Hn-pU'
			`
			msg := tgbotapi.NewMessage(youBot.Update.Message.Chat.ID, message)
			youBot.Bot.Send(msg)

		}
	}

}

// сделаю при необходимости
//func chatStateMachine(youBot  handler.YouBot) error {
//	return nil
//}
