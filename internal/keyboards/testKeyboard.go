package keyboards

import "github.com/PaulSonOfLars/gotgbot/v2"

func GetMainKeyboard() gotgbot.InlineKeyboardMarkup {
	return gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
			{Text: "Press me", CallbackData: "start_callback"},
		}}}
}
