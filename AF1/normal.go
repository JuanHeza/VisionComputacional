package main

import(
    "math"
	"os"
	"log"
	"image"
	"image/color"
	"image/jpeg"
)

func NorMi(In string, Out string) {
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
            Mini := math.Min(float64(r),math.Min(float64(g),float64(b)))
            M := uint8(Mini/256)
			pixel := color.RGBA{M,M,M,(255)}
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
func NorMa(In string, Out string) {
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
            Maxi := math.Max(float64(r),math.Max(float64(g),float64(b)))
            M := uint8(Maxi/256)
			pixel := color.RGBA{M,M,M,(255)}
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