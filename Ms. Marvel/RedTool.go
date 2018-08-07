package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"strings"
	"image"
	"image/color"
	"encoding/csv"
)
var (
	Limite = 12
	Perc = 0.0005
	Q color.Gray
)
var Umbrales = map[int]color.Palette{
   2:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{255,255,255,255}},
   3:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{128,128,128,255},color.RGBA{255,255,255,255}},
   4:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{ 85, 85, 85,255},color.RGBA{170,170,170,255},color.RGBA{255,255,255,255}},
   6:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{ 51, 51, 51,255},color.RGBA{102,102,102,255},color.RGBA{153,153,153,255},color.RGBA{204,204,204,255},color.RGBA{255,255,255,255}},
   8:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{ 36, 36, 36,255},color.RGBA{ 72, 72, 72,255},color.RGBA{108,108,108,255},color.RGBA{144,144,144,255},color.RGBA{180,180,180,255},color.RGBA{216,216,216,255},color.RGBA{255,255,255,255}},
  16:color.Palette{color.RGBA{  0,  0,  0,255},color.RGBA{ 17, 17, 17,255},color.RGBA{ 34, 34, 34,255},color.RGBA{ 51, 51, 51,255},color.RGBA{ 68, 68, 68,255},color.RGBA{ 85, 85, 85,255},color.RGBA{102,102,102,255},color.RGBA{119,119,119,255},color.RGBA{136,136,136,255},color.RGBA{153,153,153,255},color.RGBA{170,170,170,255},color.RGBA{187,187,187,255},color.RGBA{204,204,204,255},color.RGBA{221,221,221,255},color.RGBA{238,238,238,255},color.RGBA{255,255,255,255}},
}

func Marcar(Canal1 *chan Punto,MV1 [][]int) [][]int{
	for len(*Canal1) != 0{
		E := <- *Canal1
		MV1[E.Y][E.X] = E.V
	}
	return MV1
}

func Patch(MV1 [][]int, GMV int, GMH int) ([][]int){
	for IX := 0; IX <GMV; IX++{
		for IY:=0;IY<GMH;IY++{
			if MV1[IX][IY] == 0{
				V++
				MV1[IX][IY] = V
			}
		}
	}
	return MV1
}

func MatrizSalida(Q int, M [][]int){
	var OUT [][]int
	var dir string
	switch (Q){
	case 1:
		OUT = M
		dir = "SRC/Output/Entrada.csv"
		break
	case 2:
		OUT = M
		dir = "SRC/Output/Salida.csv"
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

func Umbral(Base int, Matriz [][]int) [][]int{
  Paleta := Umbrales[Base]
  for z,i := range Matriz{
      for x,j := range i{
        Q.Y = uint8(j)        
        A := Paleta.Convert(Q)
        T,_,_,_ := A.RGBA()
        Matriz[z][x] = int(T>>8)
      }
  }
  return Matriz
}

func Mark(Data *chan int){
	if len(*Data)==cap(*Data){
		<- *Data
	}else{
		*Data <- 0
	}
}

func MatrixB(matriz [][]bool){
	for _,Q:=range matriz{
		fmt.Println(Q)
	}
	fmt.Println()
}

func Matrix(matriz [][]int){
	for _,Q:=range matriz{
		fmt.Println(Q)
	}
	fmt.Println()
}

func SetImage(Original image.Image) *image.RGBA{
	b := Original.Bounds()
	Salida := image.NewRGBA(b)
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := Original.At(X, Y)
			Salida.Set(X,Y,oldPixel)
		}
	}
	return Salida
}

func Uniformar(Can [][]int, In *image.RGBA) *image.RGBA{
	var Rosemount = make(map[int][]Punto)
	for X,XY := range Can{
		for Y,Val := range XY{
			Rosemount[Val] = append(Rosemount[Val],Punto{X,Y,0})
		}
	}
	for x,D := range Rosemount{
		//var Out = make(map[color.Color] int)
		//sumaR, sumaG, sumaB, mediaB, mediaG, mediaR := 0.0,0.0,0.0,0.0,0.0,0.0
		//Maximo := float64(len(D))
		var Pixel color.Color
		F := D[0]
		Pixel = In.At(F.X,F.Y)
		/*for _,F := range D{
			SuperPixel := In.At(F.X,F.Y)
			r,g,b,_ := SuperPixel.RGBA()
			sumaR += float64(r)
			sumaG += float64(g)
			sumaB += float64(b)
		}
		mediaR = sumaR/Maximo
		mediaG = sumaG/Maximo
		mediaB = sumaB/Maximo
		Pixel = color.RGBA{uint8(mediaR/256),uint8(mediaG/256),uint8(mediaB/256),(255)}
		*/
		for _,G := range D{
			In.Set(G.X,G.Y,Pixel)
		}
		if x < Limite{
			fmt.Println(Pixel)
		}
	}
	return In
}

func PIKO(InI [][]int) (map[int]int,[][]int){
	M := make(map[int]int)
	for i:=0; i<len(InI); i++{
		for j:=0; j<len(InI[0]); j++{
			x:=InI[i][j]
			M[x]++
		}
	}
	for Z,V := range M{
		log.Println("Beep")
		if V < 100{//int(float64(len(InI)*len(InI[0]))*Perc){
			for i:=0; i<len(InI); i++{
				for j:=0; j<len(InI[0]); j++{
					if InI[i][j] == Z{
						InI[i][j] = 0
					}
				}
			}
		}
	}
	//Stop()
	fmt.Println(int(float64(len(InI)*len(InI[0]))*Perc))
	return M, InI
}
func Stop(){
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}