package main

import (
	"fmt"
	"github.com/Yiwen-Chan/qq-bot-api"
	"groupmanager/global"
	"groupmanager/modules"
	"regexp"
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
		if update.PostType == "notice" {
			log.Printf("[群管系统] 收到：%s,%s", update.NoticeType, update.GroupID)
			if update.NoticeType == "group_increase" {
				message := "欢迎新人入群~"
				bot.SendMessage(update.GroupID, "group", message)
			}
			if update.NoticeType == "group_decrease" {
				message := "非常遗憾，有人退群了~"
				bot.SendMessage(update.GroupID, "group", message)
			}
		}

		modules.Friend_add_request(bot, conf, update)
		modules.Group_add_request(bot, conf, update)

		if update.Message == nil {
			continue
		} else {
			log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

			r := regexp.MustCompile("申请头衔(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				content := parm[0][1]
				m, err := bot.SetChatMemberTitle(update.GroupID, update.UserID, content, 6000000000)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 申请头衔 ", m.Status)
				if m.Status == "ok" {
					message := "已修改"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}

			if update.UserID != conf.Master {
				continue
			}

			if update.Message.Text == "群管系统" {
				message := `=====>群管系统<=====
开启全员禁言
解除全员禁言
禁言[QQ]/[@QQ] [分钟]
解除禁言[QQ]/[@QQ]
踢出[QQ]/[@QQ]
退出群聊[群号]
修改群名片[QQ]/[@QQ] [内容]
设置群头衔[QQ]/[@QQ] [内容]
群聊转发[群] [内容]
私聊转发[QQ]/[@QQ] [内容]
注：[QQ]和[内容]之间带空格
===================`
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			if update.Message.Text == "开启全员禁言" {
				m, err := bot.RestrictAllChatMembers(update.GroupID, true)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 开启全员禁言 ", m.Status)
				if m.Status == "ok" {
					message := "已开启全群禁言"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			if update.Message.Text == "解除全员禁言" {
				m, err := bot.RestrictAllChatMembers(update.GroupID, false)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 解除全员禁言 ", m.Status)
				if m.Status == "ok" {
					message := "已解除全群禁言"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("禁言(.+)\\s(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				time := time.Duration(findInt(parm[0][2])) * time.Minute
				m, err := bot.RestrictChatMember(update.GroupID, target, time)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 禁言 ", m.Status)
				if m.Status == "ok" {
					message := "小黑屋收留成功"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("解除禁言(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				time := time.Duration(0) * time.Minute
				m, err := bot.RestrictChatMember(update.GroupID, target, time)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 解除禁言 ", m.Status)
				if m.Status == "ok" {
					message := "小黑屋释放成功"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("踢出(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				m, err := bot.KickChatMember(update.GroupID, target, false)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 踢出 ", m.Status)
				if m.Status == "ok" {
					message := "已踢出群聊，希望大家引以为戒"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("退出群聊(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				m, err := bot.LeaveChat(target, update.Message.Chat.Type, false)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 退出群聊 ", m.Status)
				if m.Status == "ok" {
					message := "已退出指定群聊"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人大失败~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("升为管理(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				m, err := bot.PromoteChatMember(update.GroupID, target, true)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 升为管理 ", m.Status)
				if m.Status == "ok" {
					message := "已绿~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("取消管理(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				m, err := bot.PromoteChatMember(update.GroupID, target, false)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 取消管理 ", m.Status)
				if m.Status == "ok" {
					message := "不绿了~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("修改群名片(.+)\\s(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				content := parm[0][2]
				m, err := bot.SetChatMemberCard(update.GroupID, target, content)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 修改群名片 ", m.Status)
				if m.Status == "ok" {
					message := "群名片已修改"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("设置群头衔(.+)\\s(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				content := parm[0][2]
				m, err := bot.SetChatMemberTitle(update.GroupID, target, content, 6000000000)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s %s", "[群管系统] 设置群头衔 ", m.Status)
				if m.Status == "ok" {
					message := "群头衔已设置"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				} else {
					message := "姬气人没有权限呢~"
					bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
				}
			}
			r = regexp.MustCompile("群聊转发(.+)\\s(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				content := parm[0][2]
				bot.SendMessage(target, "group", content)
				message := "发送成功喵~"
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
			r = regexp.MustCompile("私聊转发(.+)\\s(.+)")
			if r.MatchString(update.Message.Text) {
				parm := r.FindAllStringSubmatch(update.Message.Text, -1)
				target := findInt(parm[0][1])
				content := parm[0][2]
				bot.SendMessage(target, "private", content)
				message := "发送成功喵~"
				bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, message)
			}
		}
	}
}

func findInt(str string) int64 {
	var v int64 = 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			v = v*10 + int64(str[i]-'0')
		}
	}
	return v
}
