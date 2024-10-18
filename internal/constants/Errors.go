package constants

import (
	"errors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/pkg/logging"
)

var (
	InternalError = errors.New("internal error")
)

func HandleError(b *gotgbot.Bot, _ *ext.Context, err error) ext.DispatcherAction {
	switch err {
	case InternalError:
		logging.DefaultLogger.Error().Msg(err.Error())
	default:
		logging.DefaultLogger.Error().AnErr("an unexpected error occurred: ", err)
	}

	return ext.DispatcherActionNoop
}
