package main

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/tolgazorlu/photo-analysis/proto"
)

func (s *Server) ListImages(ctx context.Context, in *pb.ListImagesRequest) (*pb.ListImagesResponse, error) {
	query := `SELECT image_id, image_analysis, (joy + sorrow + anger + surprise) / 4.0 AS average_emotion_score FROM images`

	// Add sorting
	sortBy := "average_emotion_score DESC" // Default sorting
	if in.SortBy != "" {
		// Validate and use in.SortBy if it's a valid field
		sortBy = in.SortBy
	}
	query = fmt.Sprintf("%s ORDER BY %s", query, sortBy)

	// Implement pagination
	offset := (in.PageNumber - 1) * in.PageSize
	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, in.PageSize, offset)

	// Execute query
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []*pb.ImageDetailResponse
	for rows.Next() {
		var image pb.ImageDetailResponse
		if err := rows.Scan(&image.ImageId, &image.ImageAnalysis, &image.AverageEmotion); err != nil {
			return nil, err
		}
		images = append(images, &image)
	}

	convertedImages := make([]*pb.ImageDetail, len(images))
	for i, img := range images {
		convertedImages[i] = ConvertImageDetailResponseToImageDetail(img)
	}

	totalPages := calculateTotalPages(DB, int(in.PageSize)) // Assuming in.PageSize is int32, convert to int for the function.

	// Construct response
	response := &pb.ListImagesResponse{
		ImageDetails: convertedImages,
		CurrentPage:  in.PageNumber,
		TotalPages:   int32(totalPages), // Implement this function based on your DB's total row count
	}
	return response, nil
}

func ConvertImageDetailResponseToImageDetail(detailResponse *pb.ImageDetailResponse) *pb.ImageDetail {
	return &pb.ImageDetail{
		ImageId:             detailResponse.ImageId,
		ImageAnalysis:       detailResponse.ImageAnalysis,
		AverageEmotionScore: detailResponse.AverageEmotion,
	}
}

// calculateTotalPages helps to calculate the total number of pages available based on the total number of records and the page size.
func calculateTotalPages(db *sql.DB, pageSize int) int {
	totalItems := 3                                      // For example, fetch the total count from the database.
	totalPages := (totalItems + pageSize - 1) / pageSize // Calculate total pages needed.
	return totalPages
}
