package taillog

import (
	"fmt"
	"time"

	"github.com/sealandsigh/gotest5/logagent/etcd"
)

var tskMgr *taillogMgr

// tailTask 管理者
type taillogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry:    logEntryConf, //把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道，如果没值就阻塞了
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1 配置新增
			// 2 配置删除
			// 3 配置变更
			fmt.Println("新配置来了!", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
