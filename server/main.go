package main

import (
	"log"
	"net"

	pb "github.com/tolgazorlu/photo-analysis/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PhotoManagementServer
}

func main() {

	// config := &kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// }

	// topic := "coordinates"

	// producer, err := kafka.NewProducer(config)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < 10; i++ {
	// 	value := fmt.Sprintf("message-%d", i)
	// 	err := producer.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(value),
	// 	}, nil)
	// 	if err != nil {
	// 		fmt.Printf("Failed to produce message: %d: %v", i, err)
	// 	} else {
	// 		fmt.Printf("Produced message: %d: %s\n", i, value)
	// 	}
	// }

	// producer.Flush(15 * 1000)
	// producer.Close()

	connectToDB()
	defer DB.Close()

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Fatal to listen on: %v", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterPhotoManagementServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
