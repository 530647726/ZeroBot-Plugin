package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func Menu(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	MenuText(bot, conf, update)
}

func MenuText(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "MenuText"
	r := regexp.MustCompile(conf.Command.MenuText)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-菜单]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			message := conf.BackOk.MenuText
			bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			//返回命令结果
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
