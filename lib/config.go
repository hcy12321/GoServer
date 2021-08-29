package lib

import (
	"encoding/json"
	"fmt"
	data "main/lib/data"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type ServerConfig struct {
	DefaultRedisConfig *data.RedisConfig `json:"defaultRedisConfig"`
}

var config *ServerConfig
var once sync.Once

func GetConfig() *ServerConfig {
	once.Do(func() {
		exePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println("load exePath err")
		}

		// 这里可以根据配置文件加载或重载不同配置
		configPath := path.Join(exePath, "config", "base.json")

		file, openErr := os.Open(configPath)
		if openErr != nil {
			fmt.Printf("open file %s err", configPath)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		config = &ServerConfig{}
		decoder.Decode(config)
	})
	return config
}
