package global

import (
	"reflect"

	"fmt"
)

func FindInt(str string) int64 {
	var v int64 = 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			v = v*10 + int64(str[i]-'0')
		}
	}
	return v
}

func Authority(command string, uesrID int64, conf *JsonConfig) bool {
	fmt.Println(command)
	json := reflect.ValueOf(&conf.Authority)
	authority := string(json.Elem().FieldByName(command).String())
	master := conf.Master
	manager0 := conf.Manager.Manager0
	manager1 := conf.Manager.Manager1
	manager2 := conf.Manager.Manager2
	manager3 := conf.Manager.Manager3
	manager4 := conf.Manager.Manager4
	if authority == "master" {
		if uesrID == master {
			return true
		} else {
			return false
		}
	} else if authority == "manager" {
		if uesrID == master {
			return true
		} else if uesrID == manager0 {
			return true
		} else if uesrID == manager1 {
			return true
		} else if uesrID == manager2 {
			return true
		} else if uesrID == manager3 {
			return true
		} else if uesrID == manager4 {
			return true
		} else {
			return false
		}
	} else if authority == "all" {
		return true
	} else {
		return false
	}
}
