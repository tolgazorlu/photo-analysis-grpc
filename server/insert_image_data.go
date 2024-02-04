package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertImageData(DB *sql.DB, imageName string, imageData []byte, analysisResult string, joy float32, sorrow float32, anger float32, surprise float32) error {

	query := `
        INSERT INTO images (image_name, image_data, image_analysis, joy, sorrow, anger, surprise)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING image_id`
	var id int
	err := DB.QueryRow(query, imageName, imageData, analysisResult, joy, sorrow, anger, surprise)
	if err != nil {
		log.Printf("Failed to insert image data: %v", err)
	}

	fmt.Printf("Inserted image data with ID: %d\n", id)
	return nil
}
