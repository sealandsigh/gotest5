package main

import (
	"time"

	"github.com/sealandsigh/gotest5/day06/mylogger"
)

// 测试我们自己写 的日志库
func main() {
	a := 100
	b := "周林"
	// log := mylogger.NewLog("info")
	// 12 * 1024 是12kb 12*1024*1024 是 12MB
	log := mylogger.NewFileLogger("info", "./", "zhoulin.log", 12*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志包含a:%d b:%s", a, b)
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}
}
