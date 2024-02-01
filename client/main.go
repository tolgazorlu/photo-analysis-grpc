package main

import (
	"context"
	"log"
	"os"

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

	// Read the image file
	imageData, err := os.ReadFile("/Users/tolgazorlu/go/src/github.com/tolgazorlu/photo-analysis/cat.jpg")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	// Upload the image
	res, err := c.UploadImage(context.Background(), &pb.UploadImageRequest{
		ImageData: imageData,
		Filename:  "image.jpg",
	})
	if err != nil {
		log.Fatalf("could not upload image: %v", err)
	}
	log.Printf("Uploaded image URL: %s", res.ImageId)
}
