package main

import (
	"fmt"
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"groupmanager/modules"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Printf("GroupManager 正在启动")
	var conf *global.JsonConfig
	if global.PathExists("config.json") {
		conf = global.Load("config.json")
	}
	if conf == nil {
		err := global.DefaultConfig().Save("config.json")
		if err != nil {
			log.Fatalf("创建默认配置文件时出现错误: %v", err)
			return
		}
		log.Infof("默认配置文件已生成, 请编辑 config.json 后重启程序.")
		time.Sleep(time.Second * 5)
		return
	}
	if conf.Master == 0 {
		log.Warnf("请修改 config.json 以添加主人QQ.")
		time.Sleep(time.Second * 5)
		return
	}
	log.Printf("GroupManager 加载配置完毕")
	// Whether to use WebSocket or LongPolling depends on the address.
	// To use WebSocket, the address should be something like "ws://localhost:6700"
	url := fmt.Sprintf("ws://%v:%v", conf.Host, conf.Port)
	bot, err := qqbotapi.NewBotAPI("", url, "")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewUpdate(0)
	u.PreloadUserInfo = true
	updates, err := bot.GetUpdatesChan(u)

	log.Printf("GroupManager 启动完毕，正在运行")

	for update := range updates {

		//message类信息触发
		if update.PostType == "message" {
			log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

			if update.Message.Chat.Type == "group" {
				modules.Mute(bot, conf, update)
				modules.Kick(bot, conf, update)
				modules.Manager(bot, conf, update)
				modules.Title(bot, conf, update)
				modules.Menu(bot, conf, update)
				modules.Send(bot, conf, update)
			} else if update.Message.Chat.Type == "private" {
				modules.Menu(bot, conf, update)
				modules.Send(bot, conf, update)
			}

		}
		//request类信息触发
		if update.PostType == "request" {
			modules.Request(bot, conf, update)
		}
		//notice类信息触发
		if update.PostType == "notice" {
			modules.Notice(bot, conf, update)
		}
	}
}
