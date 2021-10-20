package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"10.12.237.171:9092","10.12.237.171:9093","10.12.237.171:9094"}, nil)
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
	time.Sleep(10 * time.Second)
}
