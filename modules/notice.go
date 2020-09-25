package modules

import (
	"fmt"
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"time"

	log "github.com/sirupsen/logrus"
)

func Group_increase_notice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	if update.PostType == "notice" {
		log.Printf("[群管系统] 收到：%s,%s", update.NoticeType, update.GroupID)
		if update.NoticeType == "group_increase" {
			message := "欢迎新人入群~"
			bot.SendMessage(update.GroupID, "group", message)
		}
	}
}

func Group_decrease_notice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	if update.PostType == "notice" {
		log.Printf("[群管系统] 收到：%s,%s", update.NoticeType, update.GroupID)
		if update.NoticeType == "group_decrease" {
			message := "非常遗憾，有人退群了~"
			bot.SendMessage(update.GroupID, "group", message)
		}
	}
}

func Notify_notice(bot *qqbotapi.BotAPI, conf *global.JsonConfig, update qqbotapi.Update) {
	if update.PostType == "notice" {
		log.Printf("[群管系统] 收到：%s,%s", update.NoticeType, update.GroupID)
		if update.NoticeType == "notify" {
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

		}
	}
}
