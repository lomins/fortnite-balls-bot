package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
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
			re := regexp.MustCompile(" га+д")

			switch {
			case strings.Contains(msg, "фортнайт") || strings.Contains(msg, "fortnite"):
				fortniteHandler(bot, update)
			case strings.Contains(msg, "влат") || vlatWords[msg]:
				vlatHandler(bot, update)
			case strings.Contains(msg, "judlemain") || strings.Contains(msg, "саня") || strings.Contains(msg, "сане") ||
				strings.Contains(msg, "сань"):
				sanyaHandler(bot, update)
			case strings.Contains(msg, "вадим") || strings.Contains(msg, "Ebatel_mamok_2014"):
				vadimHandler(bot, update)
			case strings.Contains(msg, "god") || strings.Contains(msg, " гад") || re.MatchString(msg):
				ohMyGodHandler(bot, update)
			case strings.Contains(msg, "satoru") || strings.Contains(msg, "сатору") ||
				strings.Contains(msg, "годжо") || strings.Contains(msg, "godzo"):
				godzoHandler(bot, update)
			case strings.Contains(msg, "Карин") || strings.Contains(msg, "Karin"):
				karinaHandler(bot, update)
			}
		}
	}
}

func fortniteHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	video := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FilePath("video.mp4"))
	// video.Thumb = tgbotapi.FilePath("thumb.mp4")

	_, err := bot.Send(video)
	if err != nil {
		log.Println("error sending video:", err)
	}
}

func vlatHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Влат)"

	for i := 0; i < rand.Intn(50); i++ {
		text += ")"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	bot.Send(msg)
}

func sanyaHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "К Саньку?"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	bot.Send(msg)
}

func vadimHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("vadimHuy.jpg"))

	_, err := bot.Send(photo)
	if err != nil {
		log.Println("error sending photo:", err)
	}
}

func ohMyGodHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	num := rand.Intn(4)
	filePath := "ohmygod" + strconv.Itoa(num) + ".ogg"
	audio := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FilePath(filePath))

	_, err := bot.Send(audio)
	if err != nil {
		log.Println("error sending photo:", err)
	}
}

func godzoHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	video := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FilePath("satoru.mp4"))
	// video.Thumb = tgbotapi.FilePath("thumb.mp4")

	_, err := bot.Send(video)
	if err != nil {
		log.Println("error sending video:", err)
	}
}

func karinaHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	audio := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FilePath("druzhishe.ogg"))

	_, err := bot.Send(audio)
	if err != nil {
		log.Println("error sending photo:", err)
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
