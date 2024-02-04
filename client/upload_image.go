package main

import (
	"context"
	"log"
	"os"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func uploadImage(c pb.PhotoManagementClient) {

	imageData, err := os.ReadFile("/Users/tolgazorlu/go/src/github.com/tolgazorlu/photo-analysis/bigbang.jpg")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	res, err := c.UploadImage(context.Background(), &pb.UploadImageRequest{
		ImageData: imageData,
		ImageName: "bigbang.jpg",
	})
	if err != nil {
		log.Fatalf("could not upload image: %v", err)
	}
	log.Printf("Uploaded image URL: %s", res.ImageId)

}
