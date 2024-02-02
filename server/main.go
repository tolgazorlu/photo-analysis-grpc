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

	connectToDB()

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
