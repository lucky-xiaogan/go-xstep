package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"

	//"log"
	//"time"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

var consumerGroup = "1"

func main() {
	kfversion, err := sarama.ParseKafkaVersion("0.11.0.2") // kafkaVersion is the version of kafka server like 0.11.0.2
	if err != nil {
		log.Println(err)
	}

	config := sarama.NewConfig()
	config.Version = kfversion
	config.Consumer.Return.Errors = true

	// Start with a client
	client, err := sarama.NewClient([]string{"10.12.237.171:9092"}, config)
	if err != nil {
		log.Println(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(consumerGroup, client)
	if err != nil {
		log.Println(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{"test_log"}
		handler := exampleConsumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
/*	consumer, err := sarama.NewConsumer([]string{"10.12.237.171:9092","10.12.237.171:9093","10.12.237.171:9094"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("test_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)

	defer consumer.Close()
	for partition := range partitionList { // 遍历所有的分区
		fmt.Printf("partion:%d\n", partition)
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("test_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}

		//同步消费信息
		func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for message := range pc.Messages() {
				log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
			}
		}(pc)

		//defer pc.AsyncClose()
		//// 异步从每个分区消费信息
		//go func(pc sarama.PartitionConsumer) {
		//	for msg := range pc.Messages() {
		//		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s", msg.Partition, msg.Offset, msg.Key, msg.Value)
		//	}
		//}(pc)
	}
	time.Sleep(10 * time.Second)*/
}
