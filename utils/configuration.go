package utils

import (
	"encoding/json"
	"io/ioutil"
)

func (self *Configuration) Init() *Configuration {
	if b, err := self.readConfiguration(); err != nil {
		panic(err)
	} else {
		return b
	}
}

func (self *Configuration) readConfiguration() (*Configuration, error) {
	if content, err := ioutil.ReadFile(self.ConfigPath); err != nil {
		return nil, err
	} else {
		var conf Configuration
		if err = json.Unmarshal(content, &conf); err != nil {
			return nil, err
		} else {
			conf.Redis = NewFactory(conf.Redis.Host)
			return &conf, nil
		}
	}
}
