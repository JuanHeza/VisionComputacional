package main

import (
	"os"
	"log"
	"image"
	"image/color"
	"image/jpeg"
)

func Darkness(In string,Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			r,g,b,_ := oldPixel.RGBA()
			R := uint8(((float64(r)*.5))/256)
			G := uint8(((float64(g)*.5))/256)
			B := uint8(((float64(b)*.5))/256)
			pixel := color.RGBA{R,G,B,(100)}
			imgSet.Set(X,Y,pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}