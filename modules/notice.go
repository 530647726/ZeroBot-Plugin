package modules

import (
	"fmt"
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"time"

	log "github.com/sirupsen/logrus"
)

func Notice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	GroupIncreaseNotice(bot, conf, update)
	GroupDecreaseNotice(bot, conf, update)
	NotifyNotice(bot, conf, update)
}

func GroupIncreaseNotice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "GroupIncreaseNotice"
	if update.NoticeType == "group_increase" {
		//触发命令
		commandType := "[群管系统-notice]"
		command := update.RequestType
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			message := conf.BackOk.GroupIncreaseNotice
			bot.SendMessage(conf.Master, "private", message)
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func GroupDecreaseNotice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "GroupDecreaseNotice"
	if update.NoticeType == "group_decrease" {
		//触发命令
		commandType := "[群管系统-notice]"
		command := update.RequestType
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			message := conf.BackOk.GroupDecreaseNotice
			bot.SendMessage(conf.Master, "private", message)
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func NotifyNotice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "NotifyNotice"
	if update.NoticeType == "notify" {
		//触发命令
		commandType := "[群管系统-notice]"
		command := update.RequestType
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			bot_info, err := bot.GetLoginInfo()
			if err != nil {
				log.Fatal(err)
			}
			if update.TargetID == bot_info.ID {
				time.Sleep(time.Second * 2)
				message := "请不要戳我 >_<"
				bot.SendMessage(update.GroupID, "group", message)
				message = fmt.Sprintf("[CQ:poke,qq=%v]", update.UserID)
				bot.SendMessage(update.GroupID, "group", message)
			}
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
