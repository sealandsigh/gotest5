package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sealandsigh/gotest5/logagent/conf"
	"github.com/sealandsigh/gotest5/logagent/etcd"
	"github.com/sealandsigh/gotest5/logagent/kafka"
	"gopkg.in/ini.v1"
)

// logAgent入口程序
var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	// 1. 读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			// 2. 发送到kafka
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 0. 加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Kafka_Address:", cfg.Section("kafka").Key("address"))
	// fmt.Println(cfg.Section("kafka").Key("topic"))
	// fmt.Println(cfg.Section("taillog").Key("path"))
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 2. 打开日志文件准备收集日志
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println("init taillog success.")
	// run()
}
