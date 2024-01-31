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

	// ctx := context.Background()

	// // Creates a client.
	// client, err := vision.NewImageAnnotatorClient(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// // Sets the name of the image file to annotate.
	// filename := "../cat.jpg"

	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatalf("Failed to read file: %v", err)
	// }
	// defer file.Close()
	// image, err := vision.NewImageFromReader(file)
	// if err != nil {
	// 	log.Fatalf("Failed to create image: %v", err)
	// }

	// labels, err := client.DetectLabels(ctx, image, nil, 10)
	// if err != nil {
	// 	log.Fatalf("Failed to detect labels: %v", err)
	// }

	// fmt.Println("Labels:")
	// for _, label := range labels {
	// 	fmt.Println(label.Description)
	// }
}
