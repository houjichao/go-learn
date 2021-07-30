package myRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,  //最初的链接数量
		MaxActive:   100, //最大的链接数量，不确定的可以用0，（0表示自动定义），按需分配
		IdleTimeout: 300, //链接关闭时间 300s (300s不使用自动关闭)
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword(""), redis.DialDatabase(0))
		},
	}
}

func CallRedis() {
	fmt.Println("----------redis demo start----------")
	client := pool.Get()

	defer client.Close()

	//下划线忽略返回值
	_, err := client.Do("set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}
	//这里的commandName大小写都可以
	result, err := redis.Int(client.Do("get", "abc"))
	if err != nil {
		fmt.Println("get abc failed:", err)
		return
	}
	fmt.Println("abc result:", result)
	pool.Close() //关闭连接池
	fmt.Println("----------redis demo end----------")
}
