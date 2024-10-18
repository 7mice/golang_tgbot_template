package di

import (
	"goTgTemplate/internal/controller"
	"goTgTemplate/internal/middleware"
)

type Initialization struct {
	TestController controller.TestController
	Middlewares    middleware.Middlewares
}

func NewInitialization(testController controller.TestController,
	middleware middleware.Middlewares) *Initialization {
	init := &Initialization{
		TestController: testController,
		Middlewares:    middleware,
	}

	return init
}
