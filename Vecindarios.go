/*
cls && go run vecindarios.go
*/

package main 

import (
	"math"
	"fmt"
	"log"
	"time"
)

var Matriz = [][]int {
	{135, 132, 128, 126, 121, 117},
	{130, 128, 125, 122, 118, 115},
	{126, 124, 122, 119, 116, 113},
	{121, 120, 116, 115, 114, 110},
}
var MV1, MV2 [][]int
var MH1, MH2 [][]int
var GMV, GMH, Matrices int
type Super struct{
	Vector 	[]int
	ID		int 
}
var MP1,MP2 [][]bool
var AI = 2

func main(){
	GMV = len(Matriz)//tamaño de la matriz en filas
	GMH = len(Matriz[0])//tamaño de la matriz en columnas
	MV1 = make([][]int,GMV)
	MV2 = make([][]int,GMV)
	MH1 = make([][]int,GMH)
	MH2 = make([][]int,GMH)
	MP1 = make([][]bool,GMV)
	MP2 = make([][]bool,GMV)
	for i:=0; i < GMV; i++{
		MP1[i] = make([]bool,GMH)
		MP2[i] = make([]bool,GMH)
	}

	log.Println("\tGMV:",GMV," GMH:",GMH)
	SendChan(Matriz, 1)
	SendChan(transponer(Matriz),2)
	
	Pertenencia(MV2,1)
	Pertenencia(transponer(MH2),2)

}
func SendChan (Matriz [][]int, Dir int){
	Matrices = 0
	Canal1 := make(chan Super, len(Matriz))
	for i := 0; i<len(Matriz);i++{
		Canal1 <- Super{Matriz[i], i}
	}
	for len(Canal1)>0{
		Diferencia(<-Canal1,1,Dir)
		time.Sleep(1 * time.Millisecond)
	}
	switch(Dir){
		case 1:		
			Matrix(Matriz)
			Matrix(MV1)
			Matrix(MV2)
			break
		case 2:			
			Matrix(transponer(Matriz))
			Matrix(transponer(MH1))
			Matrix(transponer(MH2))
			break
	}
	fmt.Println("WELL DONE!!\n")
}

func Diferencia (fector Super, SC int,Dir int){
	var FD  int
	vector := fector.Vector
	ID := fector.ID
	raton := make([]int,len(vector)-1)
	for q:=0;q<len(vector)-1;q++{
		raton[q]=vector[q]-vector[q+1]
	}
	switch(Dir){
	case 1:
		FD = GMH
		switch(SC){
		case 1:
			MV1[ID]=raton
			break	
		case 2:
			MV2[ID]=raton
			break
		}
		break	
	case 2:
		FD = GMV
		switch(SC){
		case 1:
			MH1[ID]=raton
			break	
		case 2:
			MH2[ID]=raton
			break
		}
		break
	}
	if (FD-len(raton))<2{
		Diferencia(Super{raton,ID},2,Dir)
	}
}


func Matrix(matriz [][]int){
	for _,Q:=range matriz{
		fmt.Println(Q)
	}
	fmt.Println()
}
func MatrixB(matriz [][]bool){
	for _,Q:=range matriz{
		fmt.Println(Q)
	}
	fmt.Println()
}

func transponer(m [][]int) [][]int{
	r := make([][]int,len(m[0]))
	for x,_ :=range r {
		r[x] = make([]int,len(m))
	}
	for y, s := range m{
		for x,e := range s{
			r[x][y]=e
		}
	}
	return r
}

func Pertenencia(matriz [][]int,dir int){
	var PCH [][]bool
	switch(dir){
	case 1:
		PCH = MP1
		break
	case 2:
		PCH = MP2
		break
	}
	for T,t := range matriz{
		for L,l := range t{
			fmt.Println(math.Abs(float64(l))," >= ",float64(AI))
			PCH[T][L] = math.Abs(float64(l)) >= float64(AI)
		}
	}
	MatrixB(PCH)
}

/*
DISCONTINUIDAD
diferencia vertical
3   4   2   5   4
2   3   3   4   3
2   2   3   3   3
1   4   1   1   4

diferencia vertical 2
-1  2   -3  -1
-1  0   -1  -1
0   -1  0   0
3   -3  0   -3

diferencia vertical 2 
AI <= |2|
-1   2      -1           -3  
-1   0  -1  -1
 0  -1   0   0
         0          -3          4   -4

         4 vecindarios

diferencia horizontal
5   4   3   4   3   2
4   4   3   3   2   3
5   4   6   4   2   3

diferencia horizontal 2
1   0   0   1   1   -1
-1  0   -3  -1  0   0   

diferencia horizontal 2
 1  0    0   1  1  -1
-1  0       -1  0   0    -3

        2 Vecindarios


*/