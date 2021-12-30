package etcd

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd 初始化

var (
	cli *clientv3.Client
)

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
