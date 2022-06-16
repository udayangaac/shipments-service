// Package yamlmgr manage the yaml files (read/write).
package yamlmgr

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Manager interface {
	Read(path string, i interface{}) (readErr error)
	Write(path string, i interface{}) (err error)
}

type yamlManager struct{}

func NewYamlManager() Manager {
	return &yamlManager{}
}

func (ym *yamlManager) Read(path string, i interface{}) (readErr error) {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("./%v", path))
	if err != nil {
		readErr = err
		return
	}
	err = yaml.Unmarshal(yamlFile, i)
	if err != nil {
		readErr = err
		return
	}
	return
}

func (ym *yamlManager) Write(path string, i interface{}) (writeErr error) {
	panic("Not implemented")
}
