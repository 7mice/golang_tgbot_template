package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"goTgTemplate/app/di"
)

func SetupMainGroup(dispatcher *ext.Dispatcher, u *di.Initialization) {
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.All, u.Middlewares.TestMiddleware), 0)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("start", u.TestController.Test), 0)
}
