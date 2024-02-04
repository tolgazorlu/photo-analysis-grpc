package main

import (
	"context"
	"fmt"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) UploadImage(ctx context.Context, in *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {

	// ! IF KAFKA DOESN'T WORK, COMMENT THE CODES BELOW

	if err := AsyncSendToCloud(in.ImageName, in.ImageData); err != nil {
		fmt.Printf("Error sending to vision api: %v\n", err)
	}

	serveFrames(in.ImageData, in.ImageName)

	// ! IF KAFKA DOESN'T, USE THE CODES BELOW

	// writer := io.Writer(os.Stdout)

	// err := detectFaces(writer, in.ImageName, in.ImageData)
	// if err != nil {
	// 	log.Println("Error detecting faces:", err)
	// }

	return &pb.UploadImageResponse{
		ImageId: "0",
	}, nil

}
