package main 

import (
_	"math"
	"sync"
_	"fmt"
	"log"
)

type Punto struct{
	X,Y,V int
}

var AI,GMV, GMH, GM int
var MV1,Matriz [][]int
var Canal1, VR, Hilos chan Punto
var MP1,MP2,MP3 [][]bool
var V,Control,progress = 1,0,0.0

func Similitudes( A [][]int, B int) [][]bool{
	Matriz = A
	AI = B
	log.Println()
	initM()
	for h,i := range Matriz{
		for H,I := range i{
			if MP1[h][H] == true{
				continue
			}
			VR = make(chan Punto,GM)
			err := Vecindad(H,h,I,V)
			if err == true {
				log.Fatal(err)
			}
			V++
		}
	}
	MV1 = Marcar(&Canal1,MV1)
	MV1 = Patch(MV1, GMV, GMH)
	go MatrizSalida(1, Matriz)
	go MatrizSalida(2, MV1)
	log.Println("WELL DONE!!\n")
	return MP3
}

func initM(){
	Control = 0
	Hilos = make(chan Punto, 4)
	GMV = len(Matriz) //tamaño de la matriz en filas
	GMH = len(Matriz[0]) //tamaño de la matriz en columnas
	GM = GMH * GMV // total de puntos
	Canal1 = make(chan Punto,GM+5)
	MP1 = make([][]bool,GMV)
	MP2 = make([][]bool,GMV)
	MP3 = make([][]bool,GMV)
	MV1 = make([][]int,GMV)
	for i := 0; i < GMV; i++{
		MP1[i] = make([]bool,GMH)
		MP2[i] = make([]bool,GMH)
		MP3[i] = make([]bool,GMH)
		MV1[i] = make([]int ,GMH)
	}
}

func Vecindad(x, y, Q, V int) bool{
	for IX := 0; IX < GMV; IX++{
		for IY := 0;IY < GMH; IY++{
			MP2[IX][IY]=false
		}
	}
	RangoU := Q + AI
	RangoD := Q - AI
	var wg sync.WaitGroup
	FD(x,y,wg)
	for len(VR) > 0{
//		fmt.Println(len(VR),"\t",V)
		JK := <- VR
		Y := JK.Y
		X := JK.X
		if Matriz[Y][X] > RangoD && Matriz[Y][X] <= RangoU && MP2[Y][X] == true{
			if progress > 100{
				return true
			}
			Control++
			Canal1 <- Punto{X,Y,V}
			FD(X,Y,wg)
			progress = (float64(Control)/float64(GMV*GMH))*100
			log.Printf("%3.4f%s Procesed",progress,"%")
		}
		if Matriz[Y][X] <= RangoD || Matriz[Y][X] > RangoU{
//			fmt.Println("FLAG")
			MP3[Y][X] = true
		}
	}
	return false
}

func FD(x,y int,wg sync.WaitGroup){
	YM := GMV-1
	XM := GMH-1
	MP2[y][x] = true
	MP1[y][x] = true
	wg.Add(4)    
	go func(){
		if x > 0 {
			if MP1[y][x-1] != true{
				GodPro(x-1, y)
			}
		}   
		wg.Done()
	}()
	go func(){
		if y > 0 {
			if MP1[y-1][x] != true{
				GodPro(x, y-1)
			}
		}
		wg.Done()
	}()
	go func(){
		if x < XM {
			if MP1[y][x+1] != true{
				GodPro(x+1, y)
			}
		} 
		wg.Done()
	}()
	go func(){
		if y < YM  {
			if MP1[y+1][x] != true{  
				GodPro(x, y+1)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func GodPro(x int, y int){
	MP1[y][x] = true
	MP2[y][x] = true
	VR <- Punto{x,y,0}
}

