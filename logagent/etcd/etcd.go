package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd 初始化

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:string`
	Topic string `json:string`
}

// 初始化etcd的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

// 从etcd中获取配置项
func GetConf(key string) (LogEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err := json.Unmarshal(ev.Value, &LogEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, err:%v", err)
		}
	}
	return
}
