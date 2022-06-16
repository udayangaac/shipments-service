package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/yamlmgr"
)

var DatabaseConf DatabaseConfig

type DatabaseConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Database        string        `yaml:"database"`
	UserName        string        `yaml:"user_name"`
	Password        string        `yaml:"password"`
}

func (sc *DatabaseConfig) Read(m yamlmgr.Manager) {
	err := m.Read(`resource/config/database.yaml`, &DatabaseConf)
	if err != nil {
		log.Fatal("Unable to read the database,yaml file")
	}
}
