package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
)

func serveFrames(imgByte []byte, imgName string) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./photos/" + imgName)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 99
	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		log.Println(err)
	}

	// Since the image file is now written, we can call detectFaces
	// You need to reset the file pointer to the beginning of the file before reading it again
	_, err = out.Seek(0, io.SeekStart) // Reset the file pointer to the start
	if err != nil {
		log.Fatalln("Error seeking file:", err)
	}

	// Prepare a writer for detectFaces to output results, if needed
	// This could be the os.Stdout or any other writer where you want to output face detection results
	writer := io.Writer(os.Stdout) // Example: Change as needed

	// Now call detectFaces with the file name
	err = detectFaces(writer, imgName)
	if err != nil {
		log.Println("Error detecting faces:", err)
	}

}
