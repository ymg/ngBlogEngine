// redis databases config

package main

import (
	"encoding/json"
	"io/ioutil"
)

type AdminJsonConfig struct {
	Db_addr string `json:"db"`
	Db_port string `json:"port"`
}
type BlogJsonConfig struct {
	Db_addr string `json:"db"`
	Db_port string `json:"port"`
}

func NewAdminConfigWithPath(config_path string) (*AdminJsonConfig, error) {

	cfg := new(AdminJsonConfig)

	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewAdminConfig() (*AdminJsonConfig, error) {

	cfg := new(AdminJsonConfig)

	_cfg, err := ioutil.ReadFile("admin_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewBlogConfigWithPath(config_path string) (*BlogJsonConfig, error) {

	cfg := new(BlogJsonConfig)

	_cfg, err := ioutil.ReadFile(config_path)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewBlogConfig() (*BlogJsonConfig, error) {

	cfg := new(BlogJsonConfig)

	_cfg, err := ioutil.ReadFile("admin_config.json")

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_cfg, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
