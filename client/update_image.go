package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func updateImage(client pb.PhotoManagementClient, imageID string, imageData []byte, imageName string, imageAnalysis string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.UpdateImage(ctx, &pb.UpdateImageRequest{
		ImageId:       imageID,
		ImageData:     imageData,
		ImageName:     imageName,
		ImageAnalysis: imageAnalysis,
	})
	if err != nil {
		log.Fatalf("Could not update image: %v", err)
		return err
	}
	log.Printf("Update success: %t", r.GetSuccess())
	return nil
}
