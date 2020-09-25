package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"

	log "github.com/sirupsen/logrus"
)

func Friend_add_request(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	if update.PostType == "request" {
		log.Printf("[群管系统] 收到：%s,%s", update.RequestType, update.UserID)
		if update.RequestType == "friend" {
			message := "有人加我了喵~"
			bot.SendMessage(conf.Master, "private", message)
		}
	}
}

func Group_add_request(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	if update.PostType == "request" {
		log.Printf("[群管系统] 收到：%s,%s", update.RequestType, update.UserID)
		if update.RequestType == "group" {
			message := "有人拉我进群了喵~"
			bot.SendMessage(conf.Master, "private", message)
		}
	}
}
