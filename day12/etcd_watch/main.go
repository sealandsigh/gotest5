package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// watch
	// 派一个哨兵监视 sample_key这个key的变化（新增删除修改）
	ch := cli.Watch(context.Background(), "sample_key")
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}

}
