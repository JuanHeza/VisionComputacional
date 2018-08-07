package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func Bin12(In string, Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal( err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b) 
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			r,g,b,_ := oldPixel.RGBA()
			y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Bin := (uint8(y/256))
			Bin=(Bin/21)
			pixel := color.Gray{uint8(Bin)*21}
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