package config

import (
	"fmt"
	"time"

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
	MaxOpenConn     int           `yaml:"max_open_conn"`
	MaxIdleConn     int           `yaml:"max_idle_conn"`
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time"`
}

func (sc *DatabaseConfig) Read(m yamlmgr.Manager) {
	path := fmt.Sprintf(`resource/config/database.yaml`)
	err := m.Read(path, &DatabaseConf)
	if err != nil {
		log.Fatal("Unable to read the database,yaml file")
	}
}
