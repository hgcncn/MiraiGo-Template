package lib

import (
	"fmt"
	"github.com/hgcncn/MiraiGo-Template/bot"
	"github.com/hgcncn/MiraiGo-Template/config"
	_ "golang.org/x/mobile/bind"
)

func init() {
	fmt.Println("bot as lib")
}

func InitBot(configJSONContent string, deviceJSONContent string) {
	config.InitWithContent([]byte(configJSONContent))
	bot.InitWithDeviceJSONContent([]byte(deviceJSONContent))
	bot.StartService()
	bot.UseProtocol(bot.IPad)
	bot.Login()
	bot.RefreshList()
}
