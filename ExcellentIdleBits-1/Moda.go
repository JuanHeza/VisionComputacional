package main 

import(
  "fmt"
_ "os"
_ "log"
 	"math"
_ "time"
  "image"
  "image/color"
)

type Pato struct{
	X,Y int
	V float64
}

var (
	AnalogV = color.Palette{
		color.RGBA{255,255,  0,255},
		color.RGBA{127,255,  0,255},
		color.RGBA{  0,255,  0,255},
		color.RGBA{100,255,100,255},
		color.RGBA{  0,255,127,255},
		color.RGBA{  0,204,  0,255},
		color.RGBA{255,255,255,255},
	}
  SPC = 1
  Size = int(math.Pow(float64(1+SPC*2),2))
	VecindadM []Pato
	Value Pato
	Blanco = color.RGBA{255,255,255,255}
)

func FiltroModa(In *image.RGBA)  *image.RGBA {
	var Kit  *image.RGBA
	fmt.Println("In Moda")

	b := In.Bounds()
	Kit = image.NewRGBA(b)
	for X := SPC; X < b.Max.X-SPC; X++ {
		for Y := SPC; Y < b.Max.Y-SPC; Y++ {
			Q := 0
			VecindadM = make([]Pato,Size)
			for x := -SPC; x <= SPC; x++{
				for y := -SPC; y <= SPC; y++{
					oldPixel := In.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					v := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					VecindadM[Q] = Pato{X:X+x, Y:Y+y, V:v}
					Q++
				}
			}
			moda := 0.0
			modaCant := 0
			for k:=0;k<Size;k++{
				tempCant := 0
				for e:=0; e<Size;e++{
					if (VecindadM[e].V == VecindadM[k].V){ tempCant++}
				}
				if (tempCant > modaCant){
					moda = VecindadM[k].V
					modaCant = tempCant
					Value = VecindadM[k]
				}
			}
			if moda == 255{
				Kit.Set(X,Y,Blanco)
			}else{
				Kit.Set(X,Y,In.At(Value.X,Value.Y))
			}		
		}
	}
	fmt.Println("Return Moda")
	return Kit
}

func FiltroMedia(In *image.RGBA)  *image.RGBA {
	var Kit  *image.RGBA
	fmt.Println("In Media")

	var sumaR,sumaG,sumaB float64
	var mediaR,mediaG,mediaB float64
	b := In.Bounds()
	Kit = image.NewRGBA(b)
	for X := SPC; X < b.Max.X-SPC; X++ {
		for Y := SPC; Y < b.Max.Y-SPC; Y++ {
			sumaR,sumaG,sumaB = 0,0,0
			for x := -SPC; x <= SPC; x++{
				for y := -SPC; y <= SPC; y++{
					oldPixel := In.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					sumaR += float64(r)
					sumaG += float64(g)
					sumaB += float64(b)
				}
			}
			mediaR = sumaR/float64(Size)
			mediaG = sumaG/float64(Size)
			mediaB = sumaB/float64(Size)
			var pixel color.Color
			pixel = color.RGBA{uint8(mediaR/256),uint8(mediaG/256),uint8(mediaB/256),(255)}
			Kit.Set(X,Y,pixel)		
		}
	}
	fmt.Println("Return Media")
	return Kit
}

func Paleton(In *image.RGBA) *image.RGBA{
	var Kit *image.RGBA
	var NewPixel color.Color
	fmt.Println("In Paleton")
	b := In.Bounds()
	Kit = image.NewRGBA(b)
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			oldPixel := In.At(X, Y)
			//if oldPixel != Blanco{
				NewPixel = AnalogV.Convert(oldPixel)
				//Kit.Set(X,Y,NewPixel)
			//}
			
			if NewPixel == Blanco{
				Kit.Set(X,Y,NewPixel)		
			}else{
				Kit.Set(X,Y,oldPixel)
			}
			
		}
	}
	fmt.Println("Return Paleton")
	return Kit
}