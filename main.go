package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	//router := routers.InitRouter()
	////静态资源
	//router.Static("/static", "./static")
	//router.Run(":8081")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	addr := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "hello",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			break
		}
		msg.Value = sarama.ByteEncoder(value)
		fmt.Println(value)

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send message Fail")
		}
		fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}
