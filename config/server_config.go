package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/yamlmgr"
)

var ServerConf ServerConfig

type ServerConfig struct {
	Jwt  JWT `yaml:"jwt"`
	Port int `yaml:"metrics_port"`
}

type JWT struct {
	Key      string `yaml:"key"`
	Duration int    `yaml:"duration"`
}

func (sc *ServerConfig) Read(m yamlmgr.Manager) {
	path := fmt.Sprintf(`resource/config/server.yaml`)
	err := m.Read(path, &ServerConf)
	if err != nil {
		log.Fatal("Unable to read the server.yaml file.")
	}
}
