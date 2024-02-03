package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tolgazorlu/photo-analysis/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetImageDetail(ctx context.Context, req *pb.ImageDetailRequest) (*pb.ImageDetailResponse, error) {
	if req.GetImageId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Image ID is required")
	}

	sqlStatement := `SELECT image_analysis FROM images WHERE image_id = $1;`
	var imageAnalysis string
	err := DB.QueryRow(sqlStatement, req.GetImageId()).Scan(&imageAnalysis)
	if err != nil {
		log.Printf("Failed to get image detail: %v", err)
		return nil,
			fmt.Errorf("failed to get image: %v", err)
	}

	return &pb.ImageDetailResponse{ImageAnalysis: imageAnalysis}, nil
}
