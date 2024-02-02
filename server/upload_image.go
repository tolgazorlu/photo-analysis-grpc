package main

import (
	"context"
	"io"
	"log"
	"os"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) UploadImage(ctx context.Context, in *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {

	serveFrames(in.ImageData, in.ImageName)

	// Prepare a writer for detectFaces to output results, if needed
	// This could be the os.Stdout or any other writer where you want to output face detection results
	writer := io.Writer(os.Stdout) // Example: Change as needed

	// Now call detectFaces with the file name
	err := detectFaces(writer, in.ImageName, in.ImageData)
	if err != nil {
		log.Println("Error detecting faces:", err)
	}

	return &pb.UploadImageResponse{
		ImageId: "13212312873912",
	}, nil

}
