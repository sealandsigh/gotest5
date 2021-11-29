package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis error:%v\n", err)
	}
	fmt.Println("连接redis成功")
	//zset
	key := "rank"
	items := []*redis.Z{
		{Score: 90, Member: "PHP"},
		{Score: 96, Member: "Golang"},
		{Score: 97, Member: "Python"},
		{Score: 99, Member: "Java"},
	}
	// 把元素都追加到key
	rdb.ZAdd(rdb.Context(), key, items...)
	// 给golang加上10分
	newScore, err := rdb.ZIncrBy(rdb.Context(), key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby faild,err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
}
