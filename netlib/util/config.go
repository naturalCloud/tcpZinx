package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	Host            string
	Port            int
	MaxConn         uint32
	Name            string
	MaxBufSize      uint32
	WorkPoolSize    uint32 //工作的协程池数量
	MaxTaskQueueLen uint32 //单个work 队列最大消息的长度

}

var ServerConf *ServerConfig

func Init(wdPath string) {
	ServerConf = &ServerConfig{
		Host:            "127.0.0.1",
		Port:            8889,
		MaxConn:         3,
		Name:            "test",
		MaxTaskQueueLen: 1024,
		WorkPoolSize:    10,
	}

	path, err := os.Getwd()
	file, err := ioutil.ReadFile(path + wdPath + "/config/speed.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, ServerConf)
	if err != nil {
		panic(err)
	}

}
