package middleware

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/internal/service"
	"goTgTemplate/pkg/logging"
)

type Middlewares interface {
	TestMiddleware(_ *gotgbot.Bot, ctx *ext.Context) error
}

type MiddlewaresImpl struct {
	testService service.TestService
}

func MiddlewaresInit() *MiddlewaresImpl {
	return &MiddlewaresImpl{}
}

func (u *MiddlewaresImpl) TestMiddleware(_ *gotgbot.Bot, ctx *ext.Context) error {
	logging.DefaultLogger.Info().Msg("Test middleware")
	//if need to stop -> return ext.EndGroups
	return ext.ContinueGroups
}
