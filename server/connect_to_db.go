package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Image struct {
	ImageName    string `jsonb:"image_name"`
	ImageData    string `jsonb:"image_data"`
	AnaylsisData string `jsonb:"analysis_data"`
}

var DB *sql.DB

func connectToDB() {

	var err error
	DB, err = sql.Open("postgres", "user=root sslmode=disable password=secret host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfuly connected")
	}

	sqlStatement := `
    CREATE TABLE IF NOT EXISTS images (
		image_id SERIAL PRIMARY KEY,
		image_data BYTEA,
		image_name TEXT,
		image_analysis TEXT,
		joy FLOAT,
		sorrow FLOAT,
		anger FLOAT,
		surprise FLOAT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute SQL statement
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table checked and created (if not exists) successfully.")

}
