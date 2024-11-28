package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
	"math/rand"
	"quran-mood-bot/utils"
)

// StartScheduler initializes the cron job and starts the scheduled task
func StartScheduler(bot *tgbotapi.BotAPI) {
	// Create a new cron job to run every day at 8 AM
	c := cron.New()
	c.AddFunc("0 8 * * *", func() {
		// Send random Ayat of the day
		SendRandomAyatOfTheDay(bot)
	})
	c.Start()
}

// SendRandomAyatOfTheDay sends a random Ayat message to a predefined chat/group
func SendRandomAyatOfTheDay(bot *tgbotapi.BotAPI) {
	// Load Ayat data
	ayatData := utils.LoadAyatData()

	// Get a random category from the Ayat data
	categories := []string{"Sad", "Motivation", "Love", "Patience", "Depression"}
	randomCategory := categories[rand.Intn(len(categories))]

	// Get the random Ayat from the selected category
	ayatList := ayatData[randomCategory]
	randomIndex := rand.Intn(len(ayatList))
	randomAyat := ayatList[randomIndex]

	// Format the message to send
	msg := fmt.Sprintf("Today's random Ayat for %s: \n\n%s", randomCategory, randomAyat)

	// Replace with actual chat ID or group ID
	chatID := int64(123456789) // Replace with actual chat ID or group ID
	SendMessage(bot, chatID, msg)
}
