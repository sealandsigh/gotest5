package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/sealandsigh/gotest5/logagent/conf"
	"github.com/sealandsigh/gotest5/logagent/etcd"
	"github.com/sealandsigh/gotest5/logagent/kafka"
	"github.com/sealandsigh/gotest5/logagent/taillog"
	"github.com/sealandsigh/gotest5/logagent/utils"
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
	// 为了实现每个logagent都拉取自己独立的配置，所以要以自己的ip地址进行区分
	ipstr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipstr)
	// 2.1 从etcd获取日志收集项目的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.getconf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	// 2.2 派一个哨兵去监视日志收集项目的变化(及时通知我的logagent实现热加载配置)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3.1 循环每一个日志收集项，创建TailOBJ
	// for _, logEntry := range logEntryConf {
	// 	// conf: *etcd.LogEntry
	// 	// logEntry.Path: 要收集的日志文件的路径
	// 	taillog.NewTailTask(logEntry.Path, logEntry.Topic)
	// }
	// 3 收集日志发往kafka
	taillog.Init(logEntryConf) // 因为newConfChan访问了tskMgr的newConfChan,这个channel是在taillog.Init(logEntryConf)执行的初始化

	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()

	// 2. 打开日志文件准备收集日志
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println("init taillog success.")
	// run()
}
