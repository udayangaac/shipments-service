package config

import "github.com/udayangaac/shipments-service/yamlmgr"

// Configuration interface for read the different type of configuration
type Configuration interface {
	Read(m yamlmgr.Manager)
}

// Configuration implementation
type Configurations []Configuration

func (configs Configurations) Init(m yamlmgr.Manager) {
	for _, c := range configs {
		c.Read(m)
	}
}
