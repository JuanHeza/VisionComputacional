package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"sort"
    "time"
	"os"
)

func Convolucion(In string, Out string) {
	var Kernel = [9]float64{0,1,0,1,-8,1,0,1,0}
	var Vecindad []float64
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
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			Q := 0
			Vecindad = make([]float64,9)
			for x := -1; x < 2; x++{
				for y := -1; y < 2; y++{
					oldPixel := img.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					Vecindad[Q] = y
					Q++
				}
			}
			var PI float64
			for i:=0;i<9;i++{
				PI += (Kernel[i]*Vecindad[i])
			}
			var pixel color.Color
				pixel = color.Gray{uint8(PI)}
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

func FiltroMediana(In string, Out string) {
	var Vecindad []float64
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
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			SuperPixel := img.At(X, Y)
			r,g,b,_ := SuperPixel.RGBA()
			RT := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Q := 0
			Vecindad = make([]float64,9)
			for x := -1; x < 2; x++{
				for y := -1; y < 2; y++{
					oldPixel := img.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					Vecindad[Q] = y
					Q++
				}
			}
			var pixel color.Color
			mediana := Vecindad[4]
			difa := mediana - RT
			difb := RT - mediana
			if difa > 20 || difb > 20 {
				pixel = color.Gray{uint8(mediana)}
			}else{
				pixel = color.Gray{uint8(RT)}
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

func FiltroMedia(In string, Out string) {
	var Vecindad []float64
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
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			SuperPixel := img.At(X, Y)
			r,g,b,_ := SuperPixel.RGBA()
			RT := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Q := 0
			Vecindad = make([]float64,9)
			for x := -1; x < 2; x++{
				for y := -1; y < 2; y++{
					oldPixel := img.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					Vecindad[Q] = y
					Q++
				}
			}

			suma := 0.0			
			
			for k:=0;k<9;k++{
				suma += Vecindad[k]
			}
			media := suma/9
			var pixel color.Color
			difa := media - RT
			difb := RT - media
			if difa > 20 || difb > 20 {
				pixel = color.Gray{uint8(media)}
			}else{
				pixel = color.Gray{uint8(RT)}
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

func FiltroModa(In string, Out string) {
	var Vecindad []float64
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
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			Q := 0
			Vecindad = make([]float64,9)
			for x := -1; x < 2; x++{
				for y := -1; y < 2; y++{
					oldPixel := img.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					Vecindad[Q] = y
					Q++
				}
			}
			moda := 0.0
			modaCant := 0
			for k:=0;k<9;k++{
				tempCant := 0
				for e:=0; e<9;e++{
					if (Vecindad[e]==Vecindad[k]){ tempCant++}
				}
				if (tempCant > modaCant){
					moda = Vecindad[k]
					modaCant = tempCant
				}
			}
			pixel := color.Gray{uint8(moda)}
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

func FiltroOrdenRango(In string, Out string, R int) {
	var Vecindad []float64
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
	for X := 1; X < b.Max.X-1; X++ {
		for Y := 1; Y < b.Max.Y-1; Y++ {
			SuperPixel := img.At(X, Y)
			r,g,b,_ := SuperPixel.RGBA()
			RT := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			Q := 0
			Vecindad = make([]float64,9)
			for x := -1; x < 2; x++{
				for y := -1; y < 2; y++{
					oldPixel := img.At(X+x, Y+y)
					r,g,b,_ := oldPixel.RGBA()
					y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
					Vecindad[Q] = y
					Q++
				}
			}
			sort.Float64s(Vecindad)
			
			var pixel color.Color
			difa := RT - Vecindad[R]
			difb := Vecindad[R] - RT
			if difa > 20 || difb >20 {
				pixel = color.Gray{uint8(Vecindad[R])}
			}else{
				pixel = color.Gray{uint8(RT)}
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

func Contraste(In string, Out string, Cn float64) {
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
			y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			var pixel color.Color
			if Cn*y >255{
				pixel = color.Gray{255}
			}else{
				pixel = color.Gray{uint8(y*Cn)}
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

func Gaussiano(In string, Out string) {
	var prob = 0.01
	var thres = 1-prob
	rand.Seed( time.Now().UnixNano())
	rnd := rand.Float64()
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
			var pixel color.Color
			if rnd > thres{
				pixel = color.Gray{255}
			}else if rnd < prob{
				pixel = color.Gray{0}
			}else{
				pixel = oldPixel
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

func UmbralAdap(In string, Out string, Um uint8) {
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
			y := 0.299*float64(r)+0.587*float64(g)+0.114*float64(b)
			var pixel color.Color
			if uint8(y)>Um{
				pixel = color.Gray{255}
			}else{
				pixel = color.Gray{0}
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

func Aclarar(In string, Out string, Br uint8) {
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
			R := Br+uint8(r/0x101)
			G := Br+uint8(g/0x101)
			B := Br+uint8(b/0x101)
			if R > 255{ R = 255}
			if B > 255{ B = 255}
			if G > 255{ G = 255}
			pixel := color.RGBA{R,G,B,(255)}
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