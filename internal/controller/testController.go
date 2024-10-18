package controller

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/internal/keyboards"
	"goTgTemplate/internal/service"
	"goTgTemplate/pkg/logging"
)

type TestController interface {
	Test(b *gotgbot.Bot, ctx *ext.Context) error
	TestCallback(b *gotgbot.Bot, ctx *ext.Context) error
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
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<b>Hello dboss!</b>"), &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: keyboards.GetMainKeyboard(),
	})
	if err != nil {
		logging.DefaultLogger.Error().AnErr("Error on test callback: ", err)
	}
	return nil
}

func (u *TestControllerImpl) TestCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery

	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "You pressed a button!",
	})
	if err != nil {
		logging.DefaultLogger.Error().AnErr("Error on test callback: ", err)
	}

	_, _, err = cb.Message.EditText(b, "You edited the start message.", nil)
	if err != nil {
		logging.DefaultLogger.Error().AnErr("Error on test callback: ", err)
	}
	return nil
}
