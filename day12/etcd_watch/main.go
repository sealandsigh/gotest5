package main

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 派一个哨兵监视 sample_key这个key的变化（新增删除修改）
	ch := cli.Watch(ctx, "sample_key")
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		for wresp.Events {

		}
	}

}
