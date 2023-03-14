package kafka

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gogf/gf/util/gconv"
)

type KafkaClient struct {
	Producer sarama.SyncProducer
}

func NewKafkaClient() KafkaClient {
	return KafkaClient{
		Producer: GetProducer(),
	}
}

var Client = NewKafkaClient()

func GetProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 5 * time.Second
	config.Metadata.AllowAutoTopicCreation = true
	//kafkaConf := vipConfig.ConfigServer.Get("KAFKA_HOSTS")

	producer, err := sarama.NewSyncProducer([]string{gconv.String("kafka:9092")}, config)
	if err != nil {
		log.Fatal("F", err)
	}
	return producer
}

func (kafkaClent KafkaClient) Send(topic, msg string) error {
	defer kafkaClent.Producer.Close()
	// 创建消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	// 发送消息
	partition, offset, err := kafkaClent.Producer.SendMessage(message)
	if err != nil {
		log.Fatal("Failed to send message: ", err)
		return err
	}

	// 打印分区和偏移量
	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}
