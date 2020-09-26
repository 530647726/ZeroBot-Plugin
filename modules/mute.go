package modules

import (
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
)

func Mute(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	MuteOne(bot, conf, update)
	MuteAll(bot, conf, update)
	UnMuteOne(bot, conf, update)
	UnMuteAll(bot, conf, update)
}

func MuteOne(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "MuteOne"
	r := regexp.MustCompile(conf.Command.MuteOne)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-禁言]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			time := time.Duration(global.FindInt(parm[0][2])) * time.Minute
			//执行命令
			m, err := bot.RestrictChatMember(groupID, target, time)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.MuteOne
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.MuteOne
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func MuteAll(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "MuteAll"
	r := regexp.MustCompile(conf.Command.MuteAll)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-禁言]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			m, err := bot.RestrictAllChatMembers(groupID, true)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.MuteAll
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.MuteAll
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func UnMuteOne(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "UnMuteOne"
	r := regexp.MustCompile(conf.Command.UnMuteOne)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-禁言]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			parm := r.FindAllStringSubmatch(command, -1)
			target := global.FindInt(parm[0][1])
			time := time.Duration(0) * time.Minute
			//执行命令
			m, err := bot.RestrictChatMember(groupID, target, time)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.UnMuteOne
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.UnMuteOne
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}

func UnMuteAll(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	commandName := "UnMuteAll"
	r := regexp.MustCompile(conf.Command.UnMuteAll)
	if r.MatchString(update.Message.Text) {
		//触发命令
		commandType := "[群管系统-禁言]"
		command := update.Message.Text
		userID := update.UserID
		groupID := update.GroupID
		log.Printf("%v[收到] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		if global.Authority(commandName, userID, conf) {
			//命令解析
			//执行命令
			m, err := bot.RestrictAllChatMembers(groupID, false)
			if err != nil {
				log.Fatal(err)
			}
			//返回命令结果
			if m.Status == "ok" {
				message := conf.BackOk.UnMuteAll
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			} else {
				message := conf.BackFail.UnMuteAll
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			log.Printf("%v[完成] |触发命令|%v |触发者|%v |触发群聊|%v", commandType, command, userID, groupID)
		} else {
			log.Printf("%v[完成] |触发命令|%v !!!!!!!没有权限或者状态为关!!!!!!!", commandType, command)
		}
	}
}
