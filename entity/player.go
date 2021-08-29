package entity

import (
	"bytes"
	"fmt"
	lib "main/lib"
	libData "main/lib/data"
	"strconv"
)

type PlayerEntity struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
	Gold int    `json:"gold"`
	Lv   int    `json:"lv"`
	Exp  int    `json:"exp"`
}

func (player *PlayerEntity) Load(data interface{}) bool {
	uid, ok := data.(string)
	if !ok {
		fmt.Printf("load player err")
		return false
	}
	redisMgr := libData.GetRedisManager()
	client := redisMgr.GetRedisClient(lib.GetConfig().DefaultRedisConfig)
	buffer := bytes.Buffer{}
	buffer.WriteString("player_")
	buffer.WriteString(uid)
	redisKey := buffer.String()

	ret, err := client.HGetAll(redisKey).Result()
	if err != nil {
		fmt.Printf("hgetall err %v", err)
	}
	player.Uid = uid
	if val, ok := ret["name"]; ok {
		player.Name = val
	} else {
		player.Name = ""
	}

	if val, ok := ret["gold"]; ok {
		player.Gold, _ = strconv.Atoi(val)
	} else {
		player.Gold = 1000
	}

	if val, ok := ret["lv"]; ok {
		player.Lv, _ = strconv.Atoi(val)
	} else {
		player.Lv = 1
	}

	if val, ok := ret["exp"]; ok {
		player.Exp, _ = strconv.Atoi(val)
	} else {
		player.Exp = 0
	}
	return true
}

func (player *PlayerEntity) Save() {
	redisMgr := libData.GetRedisManager()
	client := redisMgr.GetRedisClient(lib.GetConfig().DefaultRedisConfig)
	buffer := bytes.Buffer{}
	buffer.WriteString("player_")
	buffer.WriteString(player.Uid)
	redisKey := buffer.String()

	m := make(map[string]interface{})
	m["name"] = player.Name
	m["gold"] = player.Gold
	m["lv"] = player.Lv
	m["exp"] = player.Exp
	client.HMSet(redisKey, m)
}
