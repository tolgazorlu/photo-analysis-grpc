package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) UpdateImage(ctx context.Context, req *pb.UpdateImageRequest) (*pb.UpdateImageResponse, error) {
	// Validate request
	if req.GetImageId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Image ID is required")
	}

	sqlStatement := `
    UPDATE images
    SET image_name = $2, image_data = $3, image_analysis = $4
    WHERE image_id = $1;`

	_, err := DB.Exec(sqlStatement, req.GetImageId(), req.GetImageData(), req.GetImageName(), req.GetImageAnalysis())
	if err != nil {
		log.Printf("Failed to update image: %v", err)
		return nil, fmt.Errorf("failed to update image: %w", err)
	}

	return &pb.UpdateImageResponse{Success: true}, nil

}
