package dispatcher

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/pkg/logging"
	"time"
)

func Run(dispatcher *ext.Dispatcher, bot *gotgbot.Bot) {
	updater := ext.NewUpdater(dispatcher, nil)
	err := updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: 10 * time.Second,
			},
		},
	})
	if err != nil {
		logging.DefaultLogger.Fatal().Msgf("@%s has been started...\n", bot.User.Username) // Return error instead of panic
	}
	logging.DefaultLogger.Info().Msgf("@%s has been started...\n", bot.User.Username)
	updater.Idle()
}
