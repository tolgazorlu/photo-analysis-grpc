package main

import (
	"context"
	"log"
	"os"
	"time"

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
	imageData, err := os.ReadFile("/Users/tolgazorlu/go/src/github.com/tolgazorlu/photo-analysis/bigbang.jpg")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	// Upload the image
	res, err := c.UploadImage(context.Background(), &pb.UploadImageRequest{
		ImageData: imageData,
		ImageName: "bigbang.jpg",
	})
	if err != nil {
		log.Fatalf("could not upload image: %v", err)
	}
	log.Printf("Uploaded image URL: %s", res.ImageId)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Assume this is an existing image ID in your database
	image_id := "1"
	image_data := []byte("...") // Your image data here
	image_name := "new_image_name.jpg"
	imageA_analysis := "new analysis data"

	r, err := c.UpdateImage(ctx, &pb.UpdateImageRequest{ImageId: image_id, ImageData: image_data, ImageName: image_name, ImageAnalysis: imageA_analysis})
	if err != nil {
		log.Fatalf("could not update image: %v", err)
	}
	log.Printf("Update Success: %t", r.GetSuccess())

}
