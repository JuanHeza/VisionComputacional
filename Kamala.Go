package main
/*
cls && go run F5O.go Kamala.go
*/
import (
	"fmt"
	_"bufio"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)
var Peep [][]color.Color
var pixel color.Color
var In, Out = "C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/Johnny.jpg","C:/Users/JuanEnrique/Desktop/Wack-A-Mole/Help!!/OUT.jpg"
//var AII,Selec = 50,1
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
		RT = Similitudes(Peep,AI)
		break
	case 2:
		//RT = Discontinuo(Peep,AI)
		break
	}
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