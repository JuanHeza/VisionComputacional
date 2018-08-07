package main
/*
cls && go build F5O.go Kamala.go Khan.go getfiles.go redtool.go
cls && go run main.go kamala.go f5o.go redtool.go grey.go moda.go 
*/
import (
	"fmt"
	_"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	_"encoding/csv"
	_"strings"
)

var (
	Peep [][]color.Color
	pixel color.Color
	NivelUmbral = 3
	PK [][]int
	RT [][]bool
)
func Kamala(In image.Image, Out string,AI int, Selec int) {
	img := In
	b := img.Bounds()
	Peep := make([][]int, b.Max.X)
    for i := 0; i < b.Max.X; i++ {
        Peep[i] = make([]int, b.Max.Y)
	}
	imgSet := image.NewRGBA(b)
	imgSet = SetImage(img)

	fmt.Println("Y",b.Max.Y)
	fmt.Println("X",b.Max.X)
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			y := color.GrayModel.Convert(oldPixel)
			_,B,_,_:=y.RGBA()
			Peep[X][Y]=int(int(uint8(B)))	
		}
	}
	switch(Selec){
	case 1:
		RT,PK = Similitudes(Umbral(NivelUmbral,Peep),AI)
		break
	case 2:
		RT,PK = Similitudes(Umbral(NivelUmbral,Peep),AI)
		break
	}
//	imgSet = Uniformar(PK,imgSet)
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			if RT[X][Y]== true || PK[X][Y] == 0{
				pixel = color.RGBA{uint8(255),uint8(255),uint8(255),(255)}
			//}else{
				//DF := 0//Peep[X][Y]
				//pixel = color.RGBA{uint8(0),uint8(0),uint8(0),(255)}
			//}
			imgSet.Set(X,Y,pixel)	
			}
			oldPixel := img.At(X, Y)
			R,G,B,_ := oldPixel.RGBA()

			if Eva(uint8(R/256),uint8(G/256),uint8(B/256)) == true {
				imgSet.Set(X,Y,Blanco)	
			}
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func Eva(R,G,B uint8) bool{
	Rango := uint8(15)
	if R > G-Rango && R < G + Rango && R > B -Rango && R < B + Rango{
		return true
	}
	if G > R-Rango && G < R + Rango && G > B -Rango && G < B + Rango{
		return true
	}
	if B > G-Rango && B < G + Rango && B > R -Rango && B < R + Rango{
		return true
	}
	return false
}