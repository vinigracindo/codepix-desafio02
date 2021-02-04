package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}

	deliveryChan := make(chan kafka.Event)
	topic := "desafio02"
	msg := "Hello World!!!"

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}

	err = producer.Produce(message, deliveryChan)
	if err != nil {
		panic(err)
	}

	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message:", ev.TopicPartition)
			}
		}
	}

}
