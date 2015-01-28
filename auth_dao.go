// admin_control_dao dao

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type account struct {
	name, role string
}

type credStruct struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type AuthDAO struct{}

//server password and address
var password string
var server string
var hashUtil *Hasher

//conn pool to redis db
var authDbCon *redis.Pool

func (ad *AuthDAO) InitAuthDao(dbcfg *DbConfig) error {

	server = fmt.Sprintf("%s:%s", dbcfg.Db_addr, dbcfg.Db_port)
	password = dbcfg.Db_password

	authDbCon = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	conn := authDbCon.Get()
	defer conn.Close()

	hashUtil = new(Hasher)

	exists, err := redis.Int(conn.Do("HEXISTS", "account", "admin"))

	if err == nil && exists != 1 {

		pass, salt, err := hashUtil.NewHash("admin")

		if err == nil {
			credStruct := &credStruct{
				pass,
				salt,
			}

			jsonCred, _ := json.Marshal(credStruct)

			conn.Do("HSETNX", "account", "admin", jsonCred)
		}
	}

	return nil
}
func (ad *AuthDAO) EditAdminPassword(PassDetails struct {
	NewPassword     string
	CurrentPassword string
}) error {

	conn := authDbCon.Get()
	defer conn.Close()

	result, err := conn.Do("HGET", "account", "admin")

	if err != nil {
		return err
	}

	var dat map[string]string

	if err := json.Unmarshal(result.([]byte), &dat); err != nil {
		return err
	}

	currentP, _ := dat["password"]
	currentS, _ := dat["salt"]

	if err := hashUtil.CompareHash(currentP, currentS, PassDetails.CurrentPassword); err == nil {

		p, s, err := hashUtil.NewHash(PassDetails.NewPassword)

		newCred := credStruct{p, s}

		updatedCreds, err := json.Marshal(newCred)

		if err != nil {
			return err
		}

		if _, err := conn.Do("HSET", "account", "admin", updatedCreds); err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}
func (ad *AuthDAO) AuthenticateUser(UserDetails struct {
	Username string
	Password string
}) error {

	conn := authDbCon.Get()
	defer conn.Close()

	result, err := redis.Strings(conn.Do("HGETALL", "account"))

	if err != nil {
		return err
	}

	for i := range result {

		if string(result[i]) == UserDetails.Username {

			var userData map[string]string

			jsonBytes := []byte(result[i+1])

			if err := json.Unmarshal(jsonBytes, &userData); err != nil {
				return err
			}

			currentP, _ := userData["password"]
			currentS, _ := userData["salt"]

			if err := hashUtil.CompareHash(currentP, currentS, UserDetails.Password); err == nil {
				return nil
			}
		}
	}

	return errors.New("Failed Finding User Account")
}

/*
func (ad *AuthDAO) EditAdminUsername(current string, newUser string) error {

	conn := authDbCon.Get()
	defer conn.Close()

	result, err := conn.Do("HGET", "account", current)

	if err != nil {
		return err
	}

	success, _ := redis.Int(conn.Do("HSETNX", "account", newUser, result))

	fmt.Println(success)

	if success == 1 {
		conn.Do("HDEL", "account", current)
	}

	return nil
}
*/
