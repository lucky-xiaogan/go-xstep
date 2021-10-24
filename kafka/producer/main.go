package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回


	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.3.104:9092", "192.168.3.104:9093", "192.168.3.104:9094"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	for i := 0; i < 100; i++ {
		// 构造一个消息
		value := fmt.Sprintf("this is test log:%d", i)
		msg := &sarama.ProducerMessage{}
		msg.Topic = "test_log"
		msg.Value = sarama.StringEncoder(value)
		// 发送消息
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}

}