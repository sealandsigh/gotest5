package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func mockKV_put() {}

func ExampleKV_put(cli *clientv3.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func ExampleKV_get(cli *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "sample_key")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ExampleKV_put(cli)
	ExampleKV_get(cli)
}
