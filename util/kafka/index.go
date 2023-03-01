package kafka

import (
	"fmt"
	"log"
	"os"

	k "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewConsumer(
	brokers []string,
	groupId string,
) (c *k.Consumer) {
	if len(brokers) != 1 {
		log.Fatal("fail in init Kafka writer")
	}
	c, err := k.NewConsumer(&k.ConfigMap{
		"bootstrap.servers":               brokers[0],
		"group.id":                        groupId,
		"auto.offset.reset":               "smallest",
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		log.Fatal("fail in init Kafka writer")
	}
	return
}

func StartConsume(
	pollms int,
	consumer *k.Consumer,
	topics []string,
	callBack func(msg *k.Message),
) {
	run := true
	err := consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}

	for run {
		ev := consumer.Poll(pollms)
		switch e := ev.(type) {
		case *k.Message:
			callBack(e)
		case k.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}
