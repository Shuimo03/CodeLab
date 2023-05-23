package main

import (
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"time"
)

func AsyncProduce(topic string, limit int) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	asyncProducer, err := sarama.NewAsyncProducer([]string{"192.168.207.130:30244", "192.168.207.131:30244"}, config)
	if err != nil {
		log.Fatal("NewAsyncProducer error:", err)
	}
	defer asyncProducer.AsyncClose()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(str),
		}
		asyncProducer.Input() <- msg
	}

	for i := 0; i < limit; i++ {
		select {
		case msg := <-asyncProducer.Successes():
			log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", msg.Partition, msg.Offset, msg.Value)
		}
	}

}

func SyncProduce(topic string, limit int) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{"192.168.207.130:30244", "192.168.207.131:30244"}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))

		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, sendError := producer.SendMessage(msg)
		if sendError != nil {
			log.Println("SendMessage err: ", sendError)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}

func Consumer(topic string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"192.168.207.130:30244", "192.168.207.131:30244"}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer consumer.Close()
	//多分区消费
	//partition, err := consumer.Partitions(topic)
	//if err != nil {
	//	log.Fatal("Partitions err:", err)
	//}
	//
	//var wg sync.WaitGroup
	//wg.Add(len(partition))
	//
	//for _, partitionId := range partition {
	//	go

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}

	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}

func main() {
	//SyncProduce("test-topic", 20)
	Consumer("test-topic")
	//AsyncProduce("test-topic", 20)
}
