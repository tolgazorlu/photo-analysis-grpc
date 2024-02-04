package main

import (
	"fmt"
	"io"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func AsyncSendToCloud(image_name string, image_data []byte) error {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:8082"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	topic := "image"

	delivery_chan := make(chan kafka.Event, 10000)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(image_data)},
		delivery_chan,
	)
	if err != nil {
		fmt.Printf("Produce: %v\n", err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
					writer := io.Writer(os.Stdout)

					err := detectFaces(writer, image_name, image_data)
					if err != nil {
						fmt.Printf("Error detectFaces(): %v\n", err)
					}
				}
			}
		}
	}()

	p.Flush(15 * 1000)

	return nil
}
