package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
)

//初始化ES 准备接收kafka那边发来的数据

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}

	fmt.Println("connect to es success")
	return
}

// 发送数据到ES
func SendToES(index string, data interface{}) error {
	// ?
	// 链式操作
	put1, err := client.Index().
		Index(index).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
