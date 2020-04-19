package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

var producer sarama.AsyncProducer

// 初始化生产者
func InitProducer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		fmt.Errorf("unable to create kafka client: %v", err)
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
}

// 发送消息
func Send(topic, data string) {
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
	fmt.Println("kafka Produced message: [%s]", data)
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
