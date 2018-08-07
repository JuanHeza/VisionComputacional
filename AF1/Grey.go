package main

import (
	"fmt"
	"bufio"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

var Peep[][]color.Color
var Out "Matriz.csv"
func GreyMat(In string/*, Out string*/) string{
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
	
	Peep := make([][]color.Color, b.Max.X)
    for i := 0; i < b.Max.X; i++ {
        Peep[i] = make([]color.Color, b.Max.Y)
	}
	for x := 0; x < b.Max.X; x++ {
		for y := 0; y < b.Max.Y; y++ {
			oldPixel := img.At(x, y)
			Peep[x][y] = color.GrayModel.Convert(oldPixel)
		}
	}
	Outfile, err := os.Create(Out)
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