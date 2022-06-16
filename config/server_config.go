package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/yamlmgr"
)

var ServerConf ServerConfig

type ServerConfig struct {
	Jwt  JWT `yaml:"jwt"`
	Port int `yaml:"port"`
}

type JWT struct {
	Key      string `yaml:"key"`
	Duration int    `yaml:"duration"`
}

func (sc *ServerConfig) Read(m yamlmgr.Manager) {
	err := m.Read(`resource/config/server.yaml`, &ServerConf)
	if err != nil {
		log.Fatal("Unable to read the server.yaml file.")
	}
}
