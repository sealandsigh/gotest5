package taillog

// 专门从日志文件收集日志的模块

import (
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/sealandsigh/gotest5/logagent/kafka"
)

// var (
// 	tailObj *tail.Tail
// )

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() // 根据路径打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错

		Poll: true}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	go t.run() // 直接去采集日志发送到kafka
}

// func Init(fileName string) (err error) {
// 	config := tail.Config{
// 		ReOpen:    true,                                 // 重新打开
// 		Follow:    true,                                 // 是否跟随
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
// 		MustExist: false,                                // 文件不存在不报错

// 		Poll: true}
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("tail file failed, err:", err)
// 		return
// 	}
// 	return
// }

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志数据
			// 3.2 发往kafka
			// kafka.SendToKafka(t.topic, line.Text) // 函数调函数
			// 先把日志数据发到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取数据发到kafka
		}
	}
}

// func (t *TailTask) ReadChan() <-chan *tail.Line {
// 	return t.instance.Lines
// }

// func ReadLog() {
// 	var (
// 		line *tail.Line
// 		ok   bool
// 	)
// 	for {
// 		line, ok = <-tailObj.Lines
// 		if !ok {
// 			fmt.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
// 			time.Sleep(time.Second)
// 			continue
// 		}
// 		fmt.Println("line:", line.Text)
// 	}
// }
