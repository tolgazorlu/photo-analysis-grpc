package main

import (
	"context"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) UploadImage(ctx context.Context, in *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {

	serveFrames(in.ImageData, in.Filename)

	return &pb.UploadImageResponse{
		ImageId: "13212312873912",
	}, nil

}
