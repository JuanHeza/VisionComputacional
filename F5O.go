/*
cls && go run F5O.go
*/

package main 

import (
_	"math"
	"fmt"
	"log"
	"time"
)

type Punto struct{
	X,Y,V int
}
var AI,GMV, GMH, GM int
var MV1,Matriz [][]int
var Canal1, VR chan Punto

var MP1,MP2 [][]bool
var V = 1

func Similitudes(A [][]int, B int) [][]bool{
	Matriz = A
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
	time.Sleep(10 * time.Millisecond)	
	Marcar()
	BorderLand()
/*	Matrix(Matriz)
	Matrix(MV1)
	MatrixB(MP1)*/
	log.Println("WELL DONE!!\n")
	return MP1
}

func initM(){
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
			Canal1 <- Punto{X,Y,V}
			FD(X,Y)
			MP1[Y][X]=true
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
	MatrixB(MP2)
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
			fmt.Println(JFK)
			MP1[y][x]=JFK
		}
	}
}