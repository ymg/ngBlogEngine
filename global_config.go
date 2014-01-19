// redis databases config

package main

import (
	"encoding/json"
	"io/ioutil"
)

const (
	Db_addr     = "127.0.0.1"
	Db_port     = "6379"
	Db_password = ""
	Description = "description"
	Gplus       = "google+ url"
)

type BlogConfig struct {
	Description string `json:"description"`
	Gplus       string `json:"gplus"`
}
type DbConfig struct {
	Db_addr     string     `json:"db"`
	Db_port     string     `json:"port"`
	Db_password string     `json:"password"`
	Bconfig     BlogConfig `json:"BlogConfig"`
}

func InitDbConfigWithPath(config_path string) (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
func InitDbConfig() (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile("config.json")

	if err != nil {
		err := WriteConfigJsonToLocal()
		if err != nil {
			return nil, err
		}
	}

	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
func WriteConfigJsonToLocal() error {

	DefaultConfig := &DbConfig{}
	DefaultConfig.Db_addr = Db_addr
	DefaultConfig.Db_port = Db_port
	DefaultConfig.Db_password = Db_password
	DefaultConfig.Bconfig.Description = Description
	DefaultConfig.Bconfig.Gplus = Gplus
	DefaultJson, _ := json.MarshalIndent(&DefaultConfig, "\n", "\t")

	return ioutil.WriteFile("config.json", []byte(DefaultJson), 0644)
}
