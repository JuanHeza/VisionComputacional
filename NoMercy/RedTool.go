package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"image/color"
	"encoding/csv"
)

var Q color.Gray
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

func Patch(MV1 [][]int, GMV int, GMH int) [][]int{
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