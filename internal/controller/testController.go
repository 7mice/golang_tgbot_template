package controller

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/internal/service"
)

type TestController interface {
	Test(b *gotgbot.Bot, ctx *ext.Context) error
}

type TestControllerImpl struct {
	testService service.TestService
}

func TestControllerInit(testService service.TestService) *TestControllerImpl {
	return &TestControllerImpl{
		testService: testService,
	}
}

func (u *TestControllerImpl) Test(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<b>Hello boss!</b>"), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		// return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
