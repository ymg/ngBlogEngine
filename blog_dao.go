// posts dao

package main

import (
	"github.com/garyburd/redigo/redis"
	//"time"
)

type Post struct {
	Id, Title, Body, Date, Markdown string
}
type BlogDAO struct{}

var blogConn *redis.Pool

func (bd *BlogDAO) BlogDaoInit(c *DbConfig) error {

	/*var server string
	var password string

	blogConn := &redis.Pool{
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
	}*/

	return nil
}

func (bd *BlogDAO) NewPost(newpost *Post) error {

	conn := blogConn.Get()
	defer conn.Close()

	return nil
}

func (bd *BlogDAO) GetAll() (func() []Post, error) {

	return func() []Post {
		return nil
	}, nil
}

func (bd *BlogDAO) EditPost(newpost *Post) error {

	return nil
}

func (bd *BlogDAO) DeletePost(newpost *Post) error {

	return nil
}
