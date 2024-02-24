package main

import (
	"log"
	"math/rand"
	"os"
	"strings"

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

	fortniteWords := map[string]bool{
		"фортнайт боллс": true,
		"фортнайт болс":  true,
	}

	vlatWords := map[string]bool{
		"влат":  true,
		"влат)": true,
		"т":     true,
		"т)":    true,
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s chID %d", update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)

			msg := strings.ToLower(update.Message.Text)

			if strings.Contains(msg, "фортнайт") || strings.Contains(msg, "fortnite") || fortniteWords[msg] {
				video := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FilePath("video.mp4"))
				// video.Thumb = tgbotapi.FilePath("thumb.mp4")

				_, err := bot.Send(video)
				if err != nil {
					log.Println("error sending video:", err)
				}
			} else if strings.Contains(msg, "влат") || vlatWords[msg] {
				text := "Влат)"

				for i := 0; i < rand.Intn(50); i++ {
					text += ")"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

				bot.Send(msg)
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

// func randInt(min, max int) int {
// 	return min + rand.Int(max-min)
// }
