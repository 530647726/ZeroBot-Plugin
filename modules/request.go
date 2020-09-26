package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"

	log "github.com/sirupsen/logrus"
)

func Request(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	FriendAddRequest(bot, conf, update)
	GroupAddRequest(bot, conf, update)
}

func FriendAddRequest(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "FriendAddRequest"
	if update.RequestType == "friend" {
		//触发命令
		commandType := "[群管系统-请求]"
		command := update.RequestType
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			message := conf.BackOk.FriendAddRequest
			bot.SendMessage(conf.Master, "private", message)
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func GroupAddRequest(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "GroupAddRequest"
	if update.RequestType == "group" {
		//触发命令
		commandType := "[群管系统-请求]"
		command := update.RequestType
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			message := conf.BackOk.GroupAddRequest
			bot.SendMessage(conf.Master, "private", message)
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
