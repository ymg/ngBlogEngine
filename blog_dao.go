// posts dao

package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Post struct {
	Id, Title, Body, Date, Markdown string
}
type BlogDAO struct{}

var blogConn *redis.Pool

func (bd *BlogDAO) InitBlogDao(cfg *DbConfig) error {

	var server string
	var password string

	server = strings.Join([]string{cfg.Db_addr, cfg.Db_port}, ":")
	password = cfg.Db_password

	blogConn = &redis.Pool{
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

func (bd *BlogDAO) NewPost(newpost *Post) error {

	conn := blogConn.Get()
	defer conn.Close()

	JsonPost, err := json.Marshal(newpost)

	if err != nil {
		return err
	}

	if _, err := conn.Do("HSET", "posts", newpost.Id, JsonPost); err != nil {
		return err
	}

	return nil
}

func (bd *BlogDAO) GetAll(page int) (*[]Post, error) {

	conn := blogConn.Get()
	defer conn.Close()

	postL := []Post{}

	list, err := conn.Do("HGETALL", "posts")

	if err != nil {
		return nil, err
	}

	//var IPosts []interface{}

	if err != nil {
		return nil, err
	} else {
		idx := 1
		for _, v := range list.([]interface{}) {
			idx++
			p := Post{}

			err := json.Unmarshal(v.([]byte), &p)

			if err == nil {
				postL = append(postL, p)
			}
		}
	}

	return &postL, nil
}

func (bd *BlogDAO) Get(post *Post) (*Post, error) {

	conn := blogConn.Get()
	defer conn.Close()

	fetched, err := conn.Do("HGET", "posts", post.Id)

	if err != nil {
		return nil, err
	}

	var parsedPost *Post

	err = json.Unmarshal(fetched.([]byte), &parsedPost)

	if err != nil {
		return nil, err
	}

	return parsedPost, nil
}

func (bd *BlogDAO) UpdatePost(post *Post) error {

	p, _ := bd.Get(post)

	p.Body = post.Body
	p.Id = strings.Replace(strings.ToLower(post.Title), " ", "-", -1)
	p.Markdown = p.Markdown
	p.Title = post.Title

	conn := blogConn.Get()
	defer conn.Close()

	JsonPost, err := json.Marshal(post)

	if err != nil {
		return err
	}

	if _, err := conn.Do("HMSET", "posts", post.Id, JsonPost); err != nil {
		return err
	}

	return nil
}

func (bd *BlogDAO) DeletePost(post *Post) error {
	conn := blogConn.Get()
	defer conn.Close()

	if _, err := conn.Do("HDEL", "posts", post.Id); err != nil {
		return err
	}

	return nil
}
