package main

import (
	"bufio"
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func Colorfull(In string, Out string) {
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
			Peep[x][y] = img.At(x, y)
		}
	}

	Outfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
		defer Outfile.Close()
		w := bufio.NewWriter(Outfile)
		for _, line := range Peep {
			fmt.Fprintln(w, line)
		}
		w.Flush()
}