// global_db_config_test

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

//initialization and unmarshalling of fetched json file
//that contains db details of the admin account
func TestConfigInitialization(t *testing.T) {

	//dummy json file
	DefaultConfig := &DbConfig{}
	DefaultConfig.Db_addr = "127.0.0.1"
	DefaultConfig.Db_port = "6379"
	DefaultConfig.Db_password = "RNG!"
	DefaultConfig.Bconfig.Description = "I write about code!"
	DefaultConfig.Bconfig.Gplus = "#"
	DefaultConfig.Bconfig.BlogTitle = "N/A"
	DefaultJson, _ := json.MarshalIndent(&DefaultConfig, "\n", "\t")

	ioutil.WriteFile("config.json", []byte(DefaultJson), 0644)

	defer os.Remove("config.json")

	/* test config initialization */

	if cfg, err := InitDbConfig(); err != nil || cfg == nil {
		t.Error("ADMIN DEFAULT PATH:\tfailed reading configuration file")
	} else {
		if cfg.Db_addr == "" || cfg.Db_port == "" || cfg.Db_password == "" {
			t.Log("ADMIN DEFAULT PATH:\tone of the address values is empty")
			t.Fail()
		}
	}
	if custom_cfg, cerr := InitDbConfigWithPath("config.json"); cerr != nil || custom_cfg == nil {
		t.Error("ADMIN CUSTOM PATH:\tfailed reading configuration file")
	} else {
		if custom_cfg.Db_addr == "" || custom_cfg.Db_port == "" || custom_cfg.Db_password == "" {
			t.Log("ADMIN CUSTOM PATH:\tone or more of the address values is empty")
			t.Fail()
		}
	}

}
