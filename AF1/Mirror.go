package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func Mirror(In string, Out string) {
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
	ZR := image.Rect(0,0,(b.Max.X)*2,b.Max.Y)
	imgSet := image.NewRGBA(ZR)
	QR := b.Max.X
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			imgSet.Set(X,Y,oldPixel)	
			imgSet.Set(2*(QR)-X,Y,oldPixel)		
		}
	}
	outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}