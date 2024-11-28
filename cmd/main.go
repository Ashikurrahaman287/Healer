package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"quran-mood-bot/handlers"
)

func main() {
	// Bot Token (directly in main.go)
	botToken := "Biye Korte Hobe :)" // Replace with your actual bot token

	// Initialize the Telegram bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	bot.Debug = true

	// Log the bot username
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up the command handlers
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)

	// Handle incoming messages
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Handle commands
		go handlers.HandleCommand(update.Message, bot)
	}
}
