package main

import (
	"log"

	pb "github.com/tolgazorlu/photo-analysis/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPhotoManagementClient(conn)

	uploadImage(c)
	updateImage(c)

}
