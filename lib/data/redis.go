package lib

import (
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

type RedisManager struct {
	connector map[*RedisConfig]*redis.Client
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Pwd  string `json:"pwd"`
	Db   int    `json:"db"`
}

var mgr *RedisManager
var once sync.Once
var getChan chan bool

func GetRedisManager() *RedisManager {
	once.Do(func() {
		mgr = &RedisManager{}
		mgr.connector = make(map[*RedisConfig]*redis.Client)
		getChan = make(chan bool, 1)
	})
	return mgr
}

func getChanOver() {
	<-getChan
}

func (*RedisManager) GetRedisClient(conf *RedisConfig) *redis.Client {
	getChan <- true
	defer getChanOver()
	if mgr.connector[conf] != nil {
		return mgr.connector[conf]
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
		Password: conf.Pwd,
		DB:       conf.Db,
	})
	mgr.connector[conf] = redisClient
	return redisClient
}
