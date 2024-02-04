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

	writer := io.Writer(os.Stdout)

	err := detectFaces(writer, in.ImageName, in.ImageData)
	if err != nil {
		log.Println("Error detecting faces:", err)
	}

	return &pb.UploadImageResponse{
		ImageId: "0",
	}, nil

}
