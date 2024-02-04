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

	_, err = out.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalln("Error seeking file:", err)
	}

}
