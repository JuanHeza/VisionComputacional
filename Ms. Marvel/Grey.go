package main

import (
	"fmt"
_	"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

var (
	UmbralA color.Palette
	QT = 4
	Umbral3 = color.Palette{
		color.RGBA{  0,  0,  0,255},
		color.RGBA{128,128,128,255},
		color.RGBA{255,255,255,255},
	}
	Umbral4 = color.Palette{
		color.RGBA{  0,  0,  0,255},
		color.RGBA{ 85, 85, 85,255},
		color.RGBA{170,170,170,255},
		color.RGBA{255,255,255,255},
	}
	Paleta = color.Palette{
		color.RGBA{  0,  0,  0,255},
		color.RGBA{255,  0,  0,255},
		color.RGBA{255,127,  0,255},
		color.RGBA{255,255,  0,255},
		color.RGBA{127,255,  0,255},
		color.RGBA{  0,255,  0,255},
		color.RGBA{  0,204,  0,255},
		color.RGBA{  0,255,127,255},
		color.RGBA{  0,255,255,255},
		color.RGBA{  0,127,255,255},
		color.RGBA{  0,  0,255,255},
		color.RGBA{127,  0,255,255},
		color.RGBA{255,  0,255,255},
		color.RGBA{255,  0,127,255},	
		color.RGBA{255,255,255,255},
	}
	MapaPal = map[color.Color]color.Color{
		color.RGBA{  0,  0,  0,255}:color.RGBA{255,255,255,255},//negro
		color.RGBA{255,  0,  0,255}:color.RGBA{255,255,255,255},//rojo
		color.RGBA{255,127,  0,255}:color.RGBA{255,255,255,255},//naranja
		color.RGBA{  0,255,255,255}:color.RGBA{255,255,255,255},//cian
		color.RGBA{  0,127,255,255}:color.RGBA{255,255,255,255},//celeste
		color.RGBA{  0,  0,255,255}:color.RGBA{255,255,255,255},//azul
		color.RGBA{127,  0,255,255}:color.RGBA{255,255,255,255},//purpura
		color.RGBA{255,  0,255,255}:color.RGBA{255,255,255,255},//morado
		color.RGBA{255,  0,127,255}:color.RGBA{255,255,255,255},//fuscia
	}
	g uint32
	IN = "SRC/Input/Mala Hierba.jpg"
	Out = "SRC/Input/Out.jpg"
)
/*
func main(){
	GreyMat(IN,Out)
}
*/
func GreyMat(In string, Out string){
	fmt.Println("HOLA PATOS")
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
	fmt.Println(img.ColorModel())
	fmt.Println("X:",b.Max.X,"Y:",b.Max.Y)
	imgSet := image.NewRGBA(b)

	for x := 0; x < b.Max.X; x++ {
		for y := 0; y < b.Max.Y; y++ {
			oldPixel := img.At(x, y)
			A := Paleta.Convert(oldPixel)
			B,err := MapaPal[A]
			if err != false{
				imgSet.Set(x,y,B)
			}else{
				switch(QT){
				case 3:
					UmbralA = Umbral3
					break
				case 4:
					UmbralA = Umbral4
					break
				}
				if UmbralA.Convert(oldPixel) == Blanco{
					imgSet.Set(x,y,Blanco)
				}else{
					imgSet.Set(x,y,oldPixel)
					//imgSet.Set(x,y,AnalogV.Convert(oldPixel))
				}
			}
		}
	}
	imgSet = Paleton(imgSet)
	imgSet = FiltroMedia(imgSet)
	imgSet = FiltroModa(imgSet)
	outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
		log.Println("Goal")
}