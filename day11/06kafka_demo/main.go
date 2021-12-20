package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama第三⽅方库开发的kafka client
func main() {
	config := sarama.NewConfig()

	//config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀一个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Version = sarama.V0_11_0_2                         //kafka版本号
	config.Net.SASL.Enable = true
	config.Net.SASL.Mechanism = "PLAIN"
	config.Net.SASL.User = "xxxx"
	config.Net.SASL.Password = "xxxx"

	// 构造⼀一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "goweb_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"xxxx:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	println("kafka连接成功")
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
