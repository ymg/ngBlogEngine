package main

import "testing"

var cfg *DbConfig
var dao AuthDAO

func TestInitAuthDao(t *testing.T) {

	cfg, _ = InitDbConfig()

	if cfg == nil {
		t.Error("configuration file could not be retrieved")
	}

	err := dao.InitAuthDao(cfg)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEditAdminPassword(t *testing.T) {

	cfg, _ = InitDbConfig()

	if cfg == nil {
		t.Error("configuration file could not be retrieved")
	}

	err := dao.InitAuthDao(cfg)

	if err != nil {
		t.Error(err.Error())
	}

	dao.EditAdminPassword(struct {
		NewPassword     string
		CurrentPassword string
	}{
		"admin",
		"123",
	})

	//dao.EditAdminUsername("admin", "epic!")

}
