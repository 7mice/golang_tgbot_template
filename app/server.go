package app

import (
	"goTgTemplate/app/connections"
	"goTgTemplate/app/di"
	"goTgTemplate/app/dispatcher"
)

func Start() {
	u := di.Init()
	botInst := connections.GetBotInstance()
	dispatch := dispatcher.Init(u)
	dispatcher.Run(dispatch, botInst)
}
