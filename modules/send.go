package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func Send(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	SendPrivate(bot, conf, update)
	SendGroup(bot, conf, update)
}

func SendPrivate(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "SendPrivate"
	r := regexp.MustCompile(conf.Command.SendPrivate)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-转发]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			content := parm[0][2]
			//执行命令
			bot.SendMessage(target, "private", content)
			//返回命令结果
			message := conf.BackOk.SendPrivate
			bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		}
	}
}

func SendGroup(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "SendGroup"
	r := regexp.MustCompile(conf.Command.SendGroup)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-转发]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			content := parm[0][2]
			//执行命令
			bot.SendMessage(target, "group", content)
			//返回命令结果
			message := conf.BackOk.SendGroup
			bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
