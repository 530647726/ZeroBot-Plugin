package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func Title(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	SetTitle(bot, conf, update)
	SetCard(bot, conf, update)
}

func SetTitle(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "SetTitle"
	r := regexp.MustCompile(conf.Command.SetTitle)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-头衔]"
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
			m, err := bot.SetChatMemberTitle(groupID, target, content, 6000000000)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.SetTitle
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.SetTitle
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func SetCard(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "SetCard"
	r := regexp.MustCompile(conf.Command.SetCard)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-头衔]"
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
			m, err := bot.SetChatMemberCard(update.GroupID, target, content)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.SetCard
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.SetCard
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
