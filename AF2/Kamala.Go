package main
/*
cls && go run F5O.go Kamala.go Khan.go getfiles.go
08:54:16
*/
import (
	"fmt"
	_"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"encoding/csv"
	"strings"
)
var Peep [][]color.Color
var pixel color.Color

var RT [][]bool

//func main() {
func Kamala(In string, Out string,AI int, Selec int) {
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
	Peep := make([][]int, b.Max.X)
    for i := 0; i < b.Max.X; i++ {
        Peep[i] = make([]int, b.Max.Y)
	}
	imgSet := image.NewRGBA(b)
	fmt.Println("Y",b.Max.Y)
	fmt.Println("X",b.Max.X)
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			y := color.GrayModel.Convert(oldPixel)
			B,_,_,_:=y.RGBA()
			Peep[X][Y]=int(int(uint8(B)))	
		}
	}
	switch(Selec){
	case 1:
		RT = Similitudes(Umbral(8,Peep),AI)
		break
	case 2:
		//RT = Discontinuo(Peep,AI)
		break
	}
	MatrizSalida(3,Binit(RT))
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			if RT[X][Y]== true{
				pixel = color.RGBA{uint8(0),uint8(255),uint8(0),(255)}
			}else{
				DF := Peep[X][Y]
				pixel = color.RGBA{uint8(DF),uint8(DF),uint8(DF),(255)}
			}
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

func Binit(In [][]bool)[][]int{
	Out := make([][]int,len(In))
	for i:=0; i < len(In); i++{
		Out[i] = make([]int,len(In[i]))
	}
	for x,a := range In{
		for y,b := range a{
			if(b){
				Out[x][y]=1
			}else{
				Out[x][y]=0
			}
		}
	}
	return Out

}

func MatrizSalida(Q int, M [][]int){
	var OUT [][]int
	var dir string
	switch (Q){
	case 1:
		OUT = M
		dir = "C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/entrada.csv"
		break
	case 2:
		OUT = M
		dir = "C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/salida.csv"
		break
	case 4:
		OUT = M
		dir = "C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/Vecindarios.csv"
		break
	case 3:
		OUT = M
		dir = "C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/Bordes.csv"
		break
	}
	Outfile, err := os.Create(dir)
		if err != nil {
			log.Fatal(err)
		}
		defer Outfile.Close()

    writer := csv.NewWriter(Outfile)
    defer writer.Flush()
    for _, value := range OUT {
		st := strings.Fields(strings.Trim(fmt.Sprint(value),"[]"))
        err = writer.Write(st)
		if err != nil {
			log.Fatal(err)
		}
	}
}