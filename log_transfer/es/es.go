package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

// LogData..
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

//初始化ES 准备接收kafka那边发来的数据

func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}

	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

// 发送数据到ES
// func SendToES(index string, data interface{}) error {
// 	// ?
// 	// 链式操作
// 	put1, err := client.Index().
// 		Index(index).
// 		BodyJson(data).
// 		Do(context.Background())
// 	if err != nil {
// 		// Handle error
// 		return err
// 	}
// 	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
// 	return err
// }

// 发送数据到ES 优化后
func SendToESChan(msg *LogData) {
	// ?
	ch <- msg
}

func sendToES() {
	for {
		select {
		case msg := <-ch:
			// 链式操作
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
