package connections

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"goTgTemplate/internal/config"
	"goTgTemplate/pkg/logging"
)

func GetBotInstance() *gotgbot.Bot {
	token := config.DefaultConfig.BotToken
	if token == "" {
		logging.DefaultLogger.Fatal().Msg("TOKEN environment variable is empty")
	}

	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {

		logging.DefaultLogger.Fatal().AnErr("failed to create new bot: ", err)
	}
	return bot
}
