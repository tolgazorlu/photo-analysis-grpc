package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func updateImage(c pb.PhotoManagementClient) {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Assume this is an existing image ID in your database
	image_id := "1"
	image_data := []byte("...") // Your image data here
	image_name := "image_name_new.jpg"
	imageA_analysis := "new analysis data"

	r, err := c.UpdateImage(ctx, &pb.UpdateImageRequest{ImageId: image_id, ImageData: image_data, ImageName: image_name, ImageAnalysis: imageA_analysis})
	if err != nil {
		log.Fatalf("could not update image: %v", err)
	}
	log.Printf("Update Success: %t", r.GetSuccess())

}
