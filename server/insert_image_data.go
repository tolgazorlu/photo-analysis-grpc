package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertImageData(DB *sql.DB, imageName string, imageData []byte, analysisResult string) error {

	query := `
        INSERT INTO images (image_name, image_data, image_analysis)
        VALUES ($1, $2, $3)
        RETURNING image_id`
	var id int
	err := DB.QueryRow(query, imageName, imageData, analysisResult)
	if err != nil {
		log.Printf("Failed to insert image data: %v", err)
	}

	fmt.Printf("Inserted image data with ID: %d\n", id)
	return nil
}
