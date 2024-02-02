package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertImageData(DB *sql.DB, imageName string, imageData []byte, analysisResult string) error {
	query := `
        INSERT INTO images (image_name, image_data, analysis_data)
        VALUES ($1, $2, $3)
        RETURNING id`
	var id int
	err := DB.QueryRow(query, imageName, imageData, analysisResult).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert image data: %v", err)
		return err
	}

	fmt.Printf("Inserted image data with ID: %d\n", id)
	return nil
}
