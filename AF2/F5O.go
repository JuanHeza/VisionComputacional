/*
cls && go run F5O.go
*/

package main 

import (
_	"math"
	"fmt"
	"log"
)

type Punto struct{
	X,Y,V int
}

var AI,GMV, GMH, GM int
var MV1,Matriz [][]int
var Canal1, VR chan Punto
var MP1,MP2 [][]bool
var V,Control = 1,0

func Similitudes(A [][]int, B int) [][]bool{
	Matriz=A
	AI = B
	log.Println()
	initM()
	for h,i:= range Matriz{
		for H,I := range i{
			VR = make(chan Punto,GM)
			if MP1[h][H]!=true{
				Vecindad(H,h,I,V)
				V++
			}
		}
	}
	Marcar()
	Patch()
	MatrizSalida(1, Matriz)
	BorderLand()
	MatrizSalida(4,MV1)
	log.Println("WELL DONE!!\n")
	return MP1
}

func initM(){
	Control = 0
	GMV = len(Matriz) //tamaño de la matriz en filas
	GMH = len(Matriz[0]) //tamaño de la matriz en columnas
	GM = GMH * GMV // total de puntos
	Canal1 = make(chan Punto,GM)
	MP1 = make([][]bool,GMV)
	MP2 = make([][]bool,GMV)
	MV1 = make([][]int,GMV)
	for i:=0; i < GMV; i++{
		MP1[i] = make([]bool,GMH)
		MP2[i] = make([]bool,GMH)
		MV1[i] = make([]int ,GMH)
	}
}

func Vecindad(x, y, Q, V int){
	for IX := 0; IX <GMV; IX++{
		for IY:=0;IY<GMH;IY++{
			MP2[IX][IY]=false
		}
	}
	RangoU := Q + AI
	RangoD := Q - AI
	FD(x,y)
	for len(VR)!=0{
		JK := <- VR
		Y := JK.Y
		X := JK.X
		if Matriz[Y][X] > RangoD && Matriz[Y][X] <= RangoU && MP2[Y][X] == true && MP1[Y][X] != true{
			if float64(Control)/float64(GMV*GMH) >= 1{
				Control = 0
			}
			Control++
			Canal1 <- Punto{X,Y,V}
			FD(X,Y)
			MP1[Y][X]=true
			//Progress(float64(Control)/float64(GMV*GMH)*100)
			log.Printf("%.4f%s Procesed",(float64(Control)/float64(GMV*GMH))*100,"%")
		}
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

func FD(x,y int){
	YM := GMV-1
	XM := GMH-1
	MP2[y][x] = true
	if x > 0{
		MP2[y][x-1]=true
		VR <-Punto{x-1,y,0}
	}
	if y > 0{
		MP2[y-1][x]=true
		VR <-Punto{x,y-1,0}
	}
	if x < XM{
		MP2[y][x+1]=true
		VR <-Punto{x+1,y,0}
	}
	if y < YM{
		MP2[y+1][x]=true
		VR <-Punto{x,y+1,0}
	}
}

func Marcar(){
	for len(Canal1) != 0{
		E := <- Canal1
		MV1[E.Y][E.X] = E.V
	}
}

func BorderLand(){
	log.Println("Borderland")
	var JFK bool
	for IX := 0; IX <GMV; IX++{
		for IY:=0;IY<GMH;IY++{
			MP2[IX][IY]=false
		}
	}
	for x := 0; x < GMH; x++{
		for y := 0; y < GMV ; y++{
			JFK = false
			FD(x,y)
			for len(VR)!=0{
				F:= <- VR
				if MV1[F.Y][F.X] != MV1[y][x]{
					JFK=true
				}
			}
			MP1[y][x]=JFK
		}
	}
	MatrizSalida(2, MV1)
}

func Patch(){
	for IX := 0; IX <GMV; IX++{
		for IY:=0;IY<GMH;IY++{
			if MV1[IX][IY] == 0{
				V++
				MV1[IX][IY] = V
			}
		}
	}
}