package kafka_test

import (
	"fmt"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestListTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	for _, p := range partitions {
		fmt.Print(p.Topic)
	}
}

func TestCreateTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	err = conn.CreateTopics(kafka.TopicConfig{Topic: "maudit", NumPartitions: 3, ReplicationFactor: 1})
	if err != nil {
		t.Fatal(err)
	}
}
