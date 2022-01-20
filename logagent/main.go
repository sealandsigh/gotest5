package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sealandsigh/gotest5/logagent/conf"
	"github.com/sealandsigh/gotest5/logagent/etcd"
	"github.com/sealandsigh/gotest5/logagent/kafka"
	"github.com/sealandsigh/gotest5/logagent/taillog"
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
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
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
	// 2.1 从etcd获取日志收集项目的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.getconf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
	// 2.2 派一个哨兵去监视日志收集项目的变化(及时通知我的logagent实现热加载配置)

	// 3.1 循环每一个日志收集项，创建TailOBJ
	// for _, logEntry := range logEntryConf {
	// 	// conf: *etcd.LogEntry
	// 	// logEntry.Path: 要收集的日志文件的路径
	// 	taillog.NewTailTask(logEntry.Path, logEntry.Topic)
	// }
	// 3 收集日志发往kafka
	taillog.Init(logEntryConf)

	// 2. 打开日志文件准备收集日志
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println("init taillog success.")
	// run()
}
