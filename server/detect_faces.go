package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"fmt"

	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

// detectFaces gets faces from the Vision API for an image at the given file path.
func detectFaces(w io.Writer, file string, image_data []byte) error {

	log.Println("*** detectFaces was invoked! ***")
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	newFile := "/Users/tolgazorlu/go/src/github.com/tolgazorlu/photo-analysis/photos/" + file

	f, err := os.Open(newFile)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		log.Println(err)
		return err
	}
	annotations, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		log.Println(err)
		return err
	}

	// // Serialize the annotations to JSON
	annotationsJSON, err := json.Marshal(annotations)
	if err != nil {
		log.Printf("Error serializing annotations: %v", err)
		return err // Handle serialization errors appropriately
	}
	annotationsStr := string(annotationsJSON)

	// log.Println(annotationsStr)
	// log.Println(image_data)

	// Now, you can call insertImageData with the JSON string
	err = insertImageData(DB, file, image_data, annotationsStr)
	if err != nil {
		log.Fatalln("Error inserting image data")
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No faces found.")
	} else {
		fmt.Fprintln(w, "Faces:")
		for i, annotation := range annotations {
			fmt.Fprintln(w, "  Face", i)
			fmt.Fprintln(w, "    Anger:", annotation.AngerLikelihood)
			fmt.Fprintln(w, "    Joy:", annotation.JoyLikelihood)
			fmt.Fprintln(w, "    Surprise:", annotation.SurpriseLikelihood)
		}
	}
	return nil
}
