// redis databases config

package main

import (
	"encoding/json"
	"io/ioutil"

	//"github.com/garyburd/redigo/redis"
)

type JsonConfig struct {
	Db_addr string `json:"db"`
	Db_port string `json:"port"`
}

func NewAdminConfigWithPath(config_path string) (*JsonConfig, error) {

	cfg := new(JsonConfig)
	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
func NewAdminConfig() (*JsonConfig, error) {

	cfg := new(JsonConfig)
	_cfg, err := ioutil.ReadFile("admin_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewBlogConfigWithPath(config_path string) (*JsonConfig, error) {

	cfg := new(JsonConfig)
	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
func NewBlogConfig() (*JsonConfig, error) {

	cfg := new(JsonConfig)
	_cfg, err := ioutil.ReadFile("blog_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
