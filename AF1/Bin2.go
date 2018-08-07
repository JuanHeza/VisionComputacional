package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)
var in =""
var mascarax = [][]int {{-1,0,1},{-2,0,2},{-1,0,1}}
var mascaray = [][]int {{1,2,1},{0,0,0},{-1,-2,-1}}

func mascara (imagen image.Image){
	img, Gx, Gy, min, max, borde := aplica_m(mascarax,mascaray,imagen)
	return img, G, Gy, min, max, borde
}

func aplica_m(mascarax [][]int, mascaray[][]int, imagen image.Image){
	
}

func main(){
	imagen_f, Gx, Gy, min, max, borde, tomap := image(in)
	erosion, tomap := borde_del(imagen_f, tomap)
	imagen := borde_gr(erosion,tomap)
}

func image(img string){
	imagen := hacer_difusa(img)
	image, Gx, Gy, min, max, borde := mascara(image)
	img = im_normal(image, min, max, borde)
	save("normalizada.jpg")
	im_bin, tomap = b_n(img)
	save("binarizada.")
}












///////////////////////////////////////////////////////////////////
func Bin2(In string, Out string) {
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
			Bin := uint8(y/256)/128
			pixel := color.Gray{Bin*255}
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