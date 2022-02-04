package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func mockKV_put() {}

func ExampleKV_put(cli *clientv3.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// etcd put 3
	// value := `[{"path":"/tmp/nginx.log","topic":"web_log"},{"path":"/tmp/redis.log","topic":"redis_log"},{"path":"/tmp/mysql.log","topic":"mysql_log"}]`
	// etcd put 2
	value := `[{"path":"/tmp/nginx.log","topic":"web_log"},{"path":"/tmp/redis.log","topic":"redis_log"}]`
	// _, err := cli.Put(ctx, "/logagent/collect_config", value)
	// etcd put ip
	_, err := cli.Put(ctx, "/logagent/192.168.31.39/collect_config", value)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
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
}
