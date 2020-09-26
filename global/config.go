package global

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type JsonConfig struct {
	Host    string `json:"host"`
	Port    uint16 `json:"port"`
	Master  int64  `json:"master"`
	Manager struct {
		Manager0 int64 `json:"manager_0"`
		Manager1 int64 `json:"manager_1"`
		Manager2 int64 `json:"manager_2"`
		Manager3 int64 `json:"manager_3"`
		Manager4 int64 `json:"manager_4"`
	}
	Command struct {
		MuteOne             string `json:"mute_one"`
		MuteAll             string `json:"mute_all"`
		UnMuteOne           string `json:"un_mute_one"`
		UnMuteAll           string `json:"un_mute_all"`
		KickMember          string `json:"kick_member"`
		LeaveGroup          string `json:"leave_group"`
		PromoteManager      string `json:"promote_manager"`
		CancelManager       string `json:"cancel_manager"`
		SetTitle            string `json:"set_title"`
		SetCard             string `json:"set_card"`
		SendPrivate         string `json:"send_private"`
		SendGroup           string `json:"send_group"`
		MenuText            string `json:"menu_text"`
		FriendAddRequest    string `json:"friend_add_request"`
		GroupAddRequest     string `json:"group_add_request"`
		GroupIncreaseNotice string `json:"group_increase_notice"`
		GroupDecreaseNotice string `json:"group_decrease_notice"`
		NotifyNotice        string `json:"notify_notice"`
	} `json:"command"`
	Authority struct {
		MuteOne             string `json:"mute_one"`
		MuteAll             string `json:"mute_all"`
		UnMuteOne           string `json:"un_mute_one"`
		UnMuteAll           string `json:"un_mute_all"`
		KickMember          string `json:"kick_member"`
		LeaveGroup          string `json:"leave_group"`
		PromoteManager      string `json:"promote_manager"`
		CancelManager       string `json:"cancel_manager"`
		SetTitle            string `json:"set_title"`
		SetCard             string `json:"set_card"`
		SendPrivate         string `json:"send_private"`
		SendGroup           string `json:"send_group"`
		MenuText            string `json:"menu_text"`
		FriendAddRequest    string `json:"friend_add_request"`
		GroupAddRequest     string `json:"group_add_request"`
		GroupIncreaseNotice string `json:"group_increase_notice"`
		GroupDecreaseNotice string `json:"group_decrease_notice"`
		NotifyNotice        string `json:"notify_notice"`
	} `json:"switch"`
	BackOk struct {
		MuteOne             string `json:"mute_one"`
		MuteAll             string `json:"mute_all"`
		UnMuteOne           string `json:"un_mute_one"`
		UnMuteAll           string `json:"un_mute_all"`
		KickMember          string `json:"kick_member"`
		LeaveGroup          string `json:"leave_group"`
		PromoteManager      string `json:"promote_manager"`
		CancelManager       string `json:"cancel_manager"`
		SetTitle            string `json:"set_title"`
		SetCard             string `json:"set_card"`
		SendPrivate         string `json:"send_private"`
		SendGroup           string `json:"send_group"`
		MenuText            string `json:"menu_text"`
		FriendAddRequest    string `json:"friend_add_request"`
		GroupAddRequest     string `json:"group_add_request"`
		GroupIncreaseNotice string `json:"group_increase_notice"`
		GroupDecreaseNotice string `json:"group_decrease_notice"`
		NotifyNotice        string `json:"notify_notice"`
	} `json:"back_ok"`
	BackFail struct {
		MuteOne             string `json:"mute_one"`
		MuteAll             string `json:"mute_all"`
		UnMuteOne           string `json:"un_mute_one"`
		UnMuteAll           string `json:"un_mute_all"`
		KickMember          string `json:"kick_member"`
		LeaveGroup          string `json:"leave_group"`
		PromoteManager      string `json:"promote_manager"`
		CancelManager       string `json:"cancel_manager"`
		SetTitle            string `json:"set_title"`
		SetCard             string `json:"set_card"`
		SendPrivate         string `json:"send_private"`
		SendGroup           string `json:"send_group"`
		MenuText            string `json:"menu_text"`
		FriendAddRequest    string `json:"friend_add_request"`
		GroupAddRequest     string `json:"group_add_request"`
		GroupIncreaseNotice string `json:"group_increase_notice"`
		GroupDecreaseNotice string `json:"group_decrease_notice"`
		NotifyNotice        string `json:"notify_notice"`
	} `json:"back_fail"`
}

func DefaultConfig() *JsonConfig {
	return &JsonConfig{
		Host:   "127.0.0.1",
		Port:   2333,
		Master: 12345678,
		Manager: struct {
			Manager0 int64 `json:"manager_0"`
			Manager1 int64 `json:"manager_1"`
			Manager2 int64 `json:"manager_2"`
			Manager3 int64 `json:"manager_3"`
			Manager4 int64 `json:"manager_4"`
		}{
			Manager0: 12345678,
			Manager1: 12345678,
			Manager2: 12345678,
			Manager3: 12345678,
			Manager4: 12345678,
		},
		Command: struct {
			MuteOne             string `json:"mute_one"`
			MuteAll             string `json:"mute_all"`
			UnMuteOne           string `json:"un_mute_one"`
			UnMuteAll           string `json:"un_mute_all"`
			KickMember          string `json:"kick_member"`
			LeaveGroup          string `json:"leave_group"`
			PromoteManager      string `json:"promote_manager"`
			CancelManager       string `json:"cancel_manager"`
			SetTitle            string `json:"set_title"`
			SetCard             string `json:"set_card"`
			SendPrivate         string `json:"send_private"`
			SendGroup           string `json:"send_group"`
			MenuText            string `json:"menu_text"`
			FriendAddRequest    string `json:"friend_add_request"`
			GroupAddRequest     string `json:"group_add_request"`
			GroupIncreaseNotice string `json:"group_increase_notice"`
			GroupDecreaseNotice string `json:"group_decrease_notice"`
			NotifyNotice        string `json:"notify_notice"`
		}{
			MuteOne:             "禁言(.+)\\s(.+)",
			MuteAll:             "开启全员禁言",
			UnMuteOne:           "解除禁言(.+)",
			UnMuteAll:           "解除全员禁言",
			KickMember:          "踢出(.+)",
			LeaveGroup:          "退出群聊(.+)",
			PromoteManager:      "升为管理(.+)",
			CancelManager:       "取消管理(.+)",
			SetTitle:            "设置群头衔(.+)\\s(.+)",
			SetCard:             "修改群名片(.+)\\s(.+)",
			SendPrivate:         "私聊转发(.+)\\s(.+)",
			SendGroup:           "群聊转发(.+)\\s(.+)",
			MenuText:            "群管系统",
			FriendAddRequest:    "占位置，无作用~",
			GroupAddRequest:     "占位置，无作用~",
			GroupIncreaseNotice: "占位置，无作用~",
			GroupDecreaseNotice: "占位置，无作用~",
			NotifyNotice:        "占位置，无作用~",
		},
		Authority: struct {
			MuteOne             string `json:"mute_one"`
			MuteAll             string `json:"mute_all"`
			UnMuteOne           string `json:"un_mute_one"`
			UnMuteAll           string `json:"un_mute_all"`
			KickMember          string `json:"kick_member"`
			LeaveGroup          string `json:"leave_group"`
			PromoteManager      string `json:"promote_manager"`
			CancelManager       string `json:"cancel_manager"`
			SetTitle            string `json:"set_title"`
			SetCard             string `json:"set_card"`
			SendPrivate         string `json:"send_private"`
			SendGroup           string `json:"send_group"`
			MenuText            string `json:"menu_text"`
			FriendAddRequest    string `json:"friend_add_request"`
			GroupAddRequest     string `json:"group_add_request"`
			GroupIncreaseNotice string `json:"group_increase_notice"`
			GroupDecreaseNotice string `json:"group_decrease_notice"`
			NotifyNotice        string `json:"notify_notice"`
		}{
			MuteOne:             "manager",
			MuteAll:             "manager",
			UnMuteOne:           "manager",
			UnMuteAll:           "manager",
			KickMember:          "master",
			LeaveGroup:          "master",
			PromoteManager:      "master",
			CancelManager:       "master",
			SetTitle:            "manager",
			SetCard:             "manager",
			SendPrivate:         "master",
			SendGroup:           "master",
			MenuText:            "manager",
			FriendAddRequest:    "off",
			GroupAddRequest:     "off",
			GroupIncreaseNotice: "all",
			GroupDecreaseNotice: "all",
			NotifyNotice:        "all",
		},
		BackOk: struct {
			MuteOne             string `json:"mute_one"`
			MuteAll             string `json:"mute_all"`
			UnMuteOne           string `json:"un_mute_one"`
			UnMuteAll           string `json:"un_mute_all"`
			KickMember          string `json:"kick_member"`
			LeaveGroup          string `json:"leave_group"`
			PromoteManager      string `json:"promote_manager"`
			CancelManager       string `json:"cancel_manager"`
			SetTitle            string `json:"set_title"`
			SetCard             string `json:"set_card"`
			SendPrivate         string `json:"send_private"`
			SendGroup           string `json:"send_group"`
			MenuText            string `json:"menu_text"`
			FriendAddRequest    string `json:"friend_add_request"`
			GroupAddRequest     string `json:"group_add_request"`
			GroupIncreaseNotice string `json:"group_increase_notice"`
			GroupDecreaseNotice string `json:"group_decrease_notice"`
			NotifyNotice        string `json:"notify_notice"`
		}{
			MuteOne:             "小黑屋收留成功~",
			MuteAll:             "小黑屋收留成功~",
			UnMuteOne:           "小黑屋释放成功~",
			UnMuteAll:           "小黑屋释放成功~",
			KickMember:          "已经踢出，希望大家引以为戒~",
			LeaveGroup:          "已退出群聊~",
			PromoteManager:      "已绿~",
			CancelManager:       "不绿了~",
			SetTitle:            "群头衔设置完毕~",
			SetCard:             "群名片设置完毕~",
			SendPrivate:         "转发成功~",
			SendGroup:           "转发成功~",
			MenuText:            "施工中~",
			FriendAddRequest:    "有人加我好友了~",
			GroupAddRequest:     "有人申请加入群聊~",
			GroupIncreaseNotice: "欢迎加入群聊",
			GroupDecreaseNotice: "有人退群了",
			NotifyNotice:        "请不要戳我 >_<",
		},
		BackFail: struct {
			MuteOne             string `json:"mute_one"`
			MuteAll             string `json:"mute_all"`
			UnMuteOne           string `json:"un_mute_one"`
			UnMuteAll           string `json:"un_mute_all"`
			KickMember          string `json:"kick_member"`
			LeaveGroup          string `json:"leave_group"`
			PromoteManager      string `json:"promote_manager"`
			CancelManager       string `json:"cancel_manager"`
			SetTitle            string `json:"set_title"`
			SetCard             string `json:"set_card"`
			SendPrivate         string `json:"send_private"`
			SendGroup           string `json:"send_group"`
			MenuText            string `json:"menu_text"`
			FriendAddRequest    string `json:"friend_add_request"`
			GroupAddRequest     string `json:"group_add_request"`
			GroupIncreaseNotice string `json:"group_increase_notice"`
			GroupDecreaseNotice string `json:"group_decrease_notice"`
			NotifyNotice        string `json:"notify_notice"`
		}{
			MuteOne:             "姬气人没有权限呢~",
			MuteAll:             "姬气人没有权限呢~",
			UnMuteOne:           "姬气人没有权限呢~",
			UnMuteAll:           "姬气人没有权限呢~",
			KickMember:          "姬气人没有权限呢~",
			LeaveGroup:          "姬气人大失败~",
			PromoteManager:      "姬气人没有权限呢~",
			CancelManager:       "姬气人没有权限呢~",
			SetTitle:            "姬气人没有权限呢~",
			SetCard:             "姬气人没有权限呢~",
			SendPrivate:         "占位置，无作用~",
			SendGroup:           "占位置，无作用~",
			MenuText:            "占位置，无作用~",
			FriendAddRequest:    "占位置，无作用~",
			GroupAddRequest:     "占位置，无作用~",
			GroupIncreaseNotice: "占位置，无作用~",
			GroupDecreaseNotice: "占位置，无作用~",
			NotifyNotice:        "占位置，无作用~",
		},
	}
}

func Load(p string) *JsonConfig {
	if !PathExists(p) {
		log.Warnf("尝试加载配置文件 %v 失败: 文件不存在", p)
		return nil
	}
	c := JsonConfig{}
	err := json.Unmarshal([]byte(ReadAllText(p)), &c)
	if err != nil {
		log.Warnf("尝试加载配置文件 %v 时出现错误: %v", p, err)
		log.Infoln("原文件已备份")
		os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		return nil
	}
	return &c
}

func (c *JsonConfig) Save(p string) error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	WriteAllText(p, string(data))
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ReadAllText(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}

func WriteAllText(path, text string) {
	_ = ioutil.WriteFile(path, []byte(text), 0644)
}
