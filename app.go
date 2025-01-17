package main

import (
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"

	"github.com/hgcncn/MiraiGo-Template/bot"
	"github.com/hgcncn/MiraiGo-Template/config"
	"github.com/hgcncn/MiraiGo-Template/utils"

	_ "github.com/hgcncn/MiraiGo-Template/modules/logging"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
}

func main() {
	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidWatch)

	// 登录
	err := bot.Login()
	if err != nil {
		panic(err)
	}
	token := base64.StdEncoding.EncodeToString(bot.GetToken()) // 存储快速登录使用的 Token, 如需使用快捷登录请解除本条注释
	logrus.Infof("token: %s", token)

	// 刷新好友列表，群列表
	bot.RefreshList()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	bot.Stop()
}
