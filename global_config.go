// redis databases config

package main

import (
	"encoding/json"
	"io/ioutil"
)

const (
	Db_addr     = "127.0.0.1"
	Db_port     = "6379"
	Db_password = "random_foo"
	Description = "description"
	Gplus       = "google+ url"
	title       = "No Title"
)

type BlogConfig struct {
	Description string `json:"description"`
	Gplus       string `json:"gplus"`
	BlogTitle   string `json:"title"`
}
type DbConfig struct {
	Db_addr     string     `json:"db"`
	Db_port     string     `json:"port"`
	Db_password string     `json:"password"`
	Bconfig     BlogConfig `json:"BlogConfig"`
}

var GlobalCfg *DbConfig

func InitDbConfigWithPath(config_path string) (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	GlobalCfg = cfg

	return cfg, nil
}
func InitDbConfig() (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile("config.json")

	if err != nil {
		err := WriteConfigJsonToLocal()
		if err != nil {
			return nil, err
		} else {
			_cfg, _ = ioutil.ReadFile("config.json")
		}
	}

	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	GlobalCfg = cfg

	return cfg, nil
}
func WriteConfigJsonToLocal() error {

	DefaultConfig := &DbConfig{}
	DefaultConfig.Db_addr = Db_addr
	DefaultConfig.Db_port = Db_port
	DefaultConfig.Db_password = Db_password
	DefaultConfig.Bconfig.Description = Description
	DefaultConfig.Bconfig.Gplus = Gplus
	DefaultConfig.Bconfig.BlogTitle = title
	DefaultJson, _ := json.MarshalIndent(&DefaultConfig, "\n", "\t")

	return ioutil.WriteFile("config.json", []byte(DefaultJson), 0644)
}
