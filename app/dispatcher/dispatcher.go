package dispatcher

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/app/di"
	"goTgTemplate/app/dispatcher/handlers"
	"goTgTemplate/internal/constants"
)

func Init(u *di.Initialization) *ext.Dispatcher {
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error:       constants.HandleError,
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	{
		handlers.SetupMainGroup(dispatcher, u)
	}
	return dispatcher
}
