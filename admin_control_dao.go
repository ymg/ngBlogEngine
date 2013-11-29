// admin_control_dao dao

package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var password string
var server string

func EditAdminDetails(new_user string, new_password string) error {

	password = ""
	server = "localhost:6379"

	pool := &redis.Pool{
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

	conn := pool.Get()
	defer conn.Close()

	//hghgh := &Json_config{}
	//hghgh.Admin_db_addr = "admin_redis_db_addr_here"
	//hghgh.Blog_db_addr = "blog_redis_db_addr_here"

	//ff, _ := json.Marshal(&hghgh)

	//fmt.Println(string(ff))

	return nil
}
