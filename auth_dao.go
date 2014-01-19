// admin_control_dao dao

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//server password and address
var password string
var server string

//conn pool to redis db
var dbConPool *redis.Pool

type AuthDAO struct{}

func (ad *AuthDAO) AuthDaoInit(dbcfg DbConfig) error {

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
func (ad *AuthDAO) EditAdminDetails(new_user string, new_password string) error {

	//hghgh := &Json_config{}
	//hghgh.Admin_db_addr = "admin_redis_db_addr_here"
	//hghgh.Blog_db_addr = "blog_redis_db_addr_here"

	//ff, _ := json.Marshal(&hghgh)

	//fmt.Println(string(ff))

	return nil
}
