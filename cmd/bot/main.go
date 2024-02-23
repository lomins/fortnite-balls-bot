package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	token, ok := os.LookupEnv("TELEGRAM_APITOKEN")
	if !ok {
		log.Fatalf("Can't find .env or token")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	words := map[string]bool{
		"фортнайт боллс": true,
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s chID %d", update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)

			if words[update.Message.Text] {
				video := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FilePath("video.mp4"))
				// video.Thumb = tgbotapi.FilePath("thumb.mp4")

				_, err := bot.Send(video)
				if err != nil {
					log.Println("error sending video:", err)
				}
			}
		}
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
