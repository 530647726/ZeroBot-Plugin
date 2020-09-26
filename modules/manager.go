package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func Manager(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	PromoteManager(bot, conf, update)
	CancelManager(bot, conf, update)
}

func PromoteManager(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "PromoteManager"
	r := regexp.MustCompile(conf.Command.PromoteManager)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-管理]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			//执行命令
			m, err := bot.PromoteChatMember(groupID, target, true)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.PromoteManager
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.PromoteManager
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func CancelManager(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "CancelManager"
	r := regexp.MustCompile(conf.Command.CancelManager)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-管理]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			//执行命令
			m, err := bot.PromoteChatMember(groupID, target, false)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.CancelManager
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.CancelManager
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
