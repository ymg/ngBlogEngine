// global_db_config_test

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

/*  initialization and unmarshalling of fetched json file
that contains db details of the admin account		*/
func TestConfigurationInitialization(t *testing.T) {

	BlogConfig := &BlogJsonConfig{}
	BlogConfig.Db_addr = "127.0.0.1"
	BlogConfig.Db_port = "6379"
	BlogJson, _ := json.Marshal(&BlogConfig)

	AdminConfig := &AdminJsonConfig{}
	AdminConfig.Db_addr = "192.168.23.161"
	AdminConfig.Db_port = "6379"
	AdminJson, _ := json.Marshal(&AdminConfig)

	ioutil.WriteFile("blog_config.json", []byte(BlogJson), 0644)
	ioutil.WriteFile("admin_config.json", []byte(AdminJson), 0644)
	defer os.Remove("admin_config.json")
	defer os.Remove("blog_config.json")

	/* test admin config initialization */

	if admin_cfg, err := NewAdminConfig(); err != nil || admin_cfg == nil {
		t.Error("ADMIN DEFAULT PATH:\tfailed reading configuration file")
	} else {
		if admin_cfg.Db_addr == "" || admin_cfg.Db_port == "" {
			t.Log("ADMIN DEFAULT PATH:\tone of the address values is empty")
			t.Fail()
		}
	}

	if custom_admin_cfg, cerr := NewAdminConfigWithPath("admin_config.json"); cerr != nil || custom_admin_cfg == nil {
		t.Error("ADMIN CUSTOM PATH:\tfailed reading configuration file")
	} else {
		if custom_admin_cfg.Db_addr == "" || custom_admin_cfg.Db_port == "" {
			t.Log("ADMIN CUSTOM PATH:\tone or more of the address values is empty")
			t.Fail()
		}
	}

	/* test blog config initialization */

	if blog_cfg, err := NewAdminConfig(); err != nil || blog_cfg == nil {
		t.Error("BLOG DEFAULT PATH:\tfailed reading configuration file")
	} else {
		if blog_cfg.Db_addr == "" || blog_cfg.Db_port == "" {
			t.Log("BLOG DEFAULT PATH:\tone of the address values is empty")
			t.Fail()
		}
	}

	if custom_blog_cfg, cerr := NewAdminConfigWithPath("blog_config.json"); cerr != nil || custom_blog_cfg == nil {
		t.Error("BLOG CUSTOM PATH:\tfailed reading configuration file")
	} else {
		if custom_blog_cfg.Db_addr == "" || custom_blog_cfg.Db_port == "" {
			t.Log("BLOG CUSTOM PATH:\tone or more of the address values is empty")
			t.Fail()
		}
	}

}
