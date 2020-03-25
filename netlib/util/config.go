package util

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	Host       string
	Port       int
	MaxConn    uint32
	Name       string
	MaxBufSize uint32
}

var ServerConf *ServerConfig

func init() {
	ServerConf = &ServerConfig{
		Host:    "127.0.0.1",
		Port:    8889,
		MaxConn: 3,
		Name:    "test",
	}
	file, err := ioutil.ReadFile("config/speed.json")
	if err != nil {
		return
		panic(err)
	}

	err = json.Unmarshal(file, ServerConf)
	if err != nil {
		panic(err)
	}

}
