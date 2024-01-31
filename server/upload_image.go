package main

import (
	"context"
	"log"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) UploadImage(ctx context.Context, in *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.UploadImageResponse{
		ImageId: "13212312873912",
	}, nil
}
