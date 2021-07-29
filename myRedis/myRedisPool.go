package myRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		MaxIdle: 16, //最初的链接数量
		MaxActive: 100, //最大的链接数量

	}
}

func CallRedis()  {
	fmt.Println("")
}