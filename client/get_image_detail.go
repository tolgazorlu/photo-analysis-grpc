package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func getImageDetail(c pb.PhotoManagementClient) {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Assume this is an existing image ID in your database
	image_id := "2"

	r, err := c.GetImageDetail(ctx, &pb.ImageDetailRequest{ImageId: image_id})
	if err != nil {
		log.Fatalf("could not get image detail: %v", err)
	}
	log.Printf("Image Detail: %s", r.GetImageAnalysis())

}
