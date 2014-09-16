// admin_control_dao dao

package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type account struct {
	name, role string
}

type AuthDAO struct{}

//server password and address
var password string
var server string

//conn pool to redis db
var dbConPool *redis.Pool

func (ad *AuthDAO) InitAuthDao(dbcfg *DbConfig) error {

	server = fmt.Sprintf("%s:%s", dbcfg.Db_addr, dbcfg.Db_port)
	password = dbcfg.Db_password

	dbConPool = &redis.Pool{
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

	return nil
}
func (ad *AuthDAO) EditAdminDetails(current_user string) error {

	result, err := dbConPool.Get().Do("HGET", "account", "admin")

	if err != nil {
		return err
	}

	fmt.Println(string(result.([]byte)))

	return nil
}
