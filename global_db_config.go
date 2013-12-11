// redis databases config

package main

import (
	"encoding/json"
	"io/ioutil"

	//"github.com/garyburd/redigo/redis"
)

type DbConfig struct {
	Db_addr string `json:"db"`
	Db_port string `json:"port"`
}

func InitAdminConfigWithPath(config_path string) (*DbConfig, error) {

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
func InitAdminConfig() (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile("admin_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitBlogConfigWithPath(config_path string) (*DbConfig, error) {

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
func InitBlogConfig() (*DbConfig, error) {

	cfg := new(DbConfig)
	_cfg, err := ioutil.ReadFile("blog_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
