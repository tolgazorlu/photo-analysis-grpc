package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func getImageFeed(c pb.PhotoManagementClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	page_size := int32(1)
	page_number := int32(1)

	r, err := c.GetImageFeed(ctx, &pb.ListImagesRequest{PageSize: page_size, PageNumber: page_number})
	if err != nil {
		log.Fatalf("could not get image feed: %v", err)
	}

	for _, imageDetail := range r.ImageDetails {
		log.Printf("Image Feed: %v", imageDetail)
	}

}
