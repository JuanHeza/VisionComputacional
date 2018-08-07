package main
//escalar los datos del histograma
import (
	_"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

var Histograma map[uint8]int

func Histo(In string, Out string) uint8{
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
			//r,g,b,_ := oldPixel.RGBA()
			_,g,_,_ := oldPixel.RGBA()
			y := g//0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Bin := (uint8(y/256))
			Histograma[Bin] ++
		}
	}

	ZR := image.Rect(0,0,255,255)
	imgSet := image.NewRGBA(ZR)
	PixMax := uint8(0)
	Vl := uint8(0)
	Cont := 0
	ValConst := false
	for x := 0; x < 255; x++{
		REGHist:= uint8(Histograma[uint8(x)] >> 1)
		C:= uint8(Histograma[uint8(x)+1] >> 1)
		B:= uint8(Histograma[uint8(x)+2] >> 1)
		A:= uint8(Histograma[uint8(x)+3] >> 1)
		D:= uint8(Histograma[uint8(x)+4] >> 1)
		if ValConst == true {
			continue
		}
		if REGHist > 100 && REGHist > PixMax && REGHist > A&& REGHist > C&& REGHist > B&& REGHist > D{
			if ValConst != true{
				PixMax = REGHist
				Vl = uint8(x)
				Cont = 0
			}
		}else{
			if PixMax > 100{
			Cont++
				if Cont > 10{
					ValConst = true
				}
			}
		}
		for y := 0; y < int(REGHist); y++{			
			var pixel color.Color
			pixel = color.Gray{255}
			imgSet.Set(x,255-y,pixel)
		} 	
	}
	outputfile, err := os.Create(Out)
	if err != nil {
		log.Fatal(err)
	}
		defer outputfile.Close()
	jpeg.Encode(outputfile,imgSet,nil)
	return Vl
}