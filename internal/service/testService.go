package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goTgTemplate/internal/repository"
)

type TestService interface {
	Test(b *gotgbot.Bot, ctx *ext.Context) (string, error)
}

type TestServiceImpl struct {
	testRepository repository.TestRepository
}

func TestServiceInit(testRepository repository.TestRepository) *TestServiceImpl {
	return &TestServiceImpl{
		testRepository: testRepository,
	}
}

func (u *TestServiceImpl) Test(b *gotgbot.Bot, ctx *ext.Context) (string, error) {
	return u.testRepository.Test()
}
