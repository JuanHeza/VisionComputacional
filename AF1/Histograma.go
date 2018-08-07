package main
//escalar los datos del histograma
import (
	_"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	_ "fmt"
	"os"
)

var Histograma map[uint8]int

func Histo(In string, Out string) {
	Histograma = make(map[uint8]int)
	for i := 0 ; i < 256; i++{
		Histograma[uint8(i)] = 0
	}
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
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			r,g,b,_ := oldPixel.RGBA()
			y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Bin := (uint8(y/256))
			Histograma[Bin] ++
		}
	}

	ZR := image.Rect(0,0,255,255)
	imgSet := image.NewRGBA(ZR)
	for x := 0; x < 255; x++{
		REGHist:= uint8(Histograma[uint8(x)] >> 1)
		for y := 0; y < int(REGHist); y++{
		//for y := 0; y < Histograma[uint8(x)]; y++{
			pixel := color.Gray{255}
			imgSet.Set(x,255-y,pixel)
		} 	
	}
	outputfile, err := os.Create(Out)
	if err != nil {
		log.Fatal(err)
	}
		defer outputfile.Close()
	jpeg.Encode(outputfile,imgSet,nil)
}