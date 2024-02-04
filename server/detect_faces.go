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

	var analysis []string
	var joy float32
	var sorrow float32
	var anger float32
	var surprise float32

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No faces found.")
	} else {
		fmt.Fprintln(w, "Faces:")
		for i, annotation := range annotations {
			analysis = append(analysis, (fmt.Sprint(i) + ": { Anger: " + annotation.AngerLikelihood.String() + " Joy: " + annotation.JoyLikelihood.String() + " Surprise: " + annotation.SurpriseLikelihood.String() + "}"))
			joy = float32(annotation.JoyLikelihood)
			sorrow = float32(annotation.SorrowLikelihood)
			anger = float32(annotation.AngerLikelihood)
			surprise = float32(annotation.SurpriseLikelihood)
			fmt.Print(analysis)
		}
	}

	annotationsJSON, err := json.Marshal(analysis)
	if err != nil {
		log.Printf("Error serializing annotations: %v", err)
		return err
	}
	annotationsStr := string(annotationsJSON)

	err = insertImageData(DB, file, image_data, annotationsStr, joy, sorrow, anger, surprise)
	if err != nil {
		log.Fatalln("Error inserting image data")
	}

	return nil
}
