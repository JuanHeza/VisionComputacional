package main

import (
	"fmt"
)
var(
	Vl Punto
	Max int
	Puntos = []Punto {
		{1,1},{2,2},{3,3},{0,1},{1,4},{-1,-2},{-1,2},
	} 
)

type Punto struct{
	X,Y float64
}

type Tri struct{
	X,Y,Z float64
}

func main() {
	var Hough = make(map[Punto][]Punto)
	for _,A := range Puntos{
		for _,B := range Puntos{
			if A == B{
				continue
			}
		MC := MCPlane(A,B)
		HoughAdd(MC,A,Hough)
		fmt.Println(A,"-",B," ->",MC)
		}
	}
	for q,r := range Hough{
		fmt.Println(q," - ",r)
		if Max < len(r){
			Max = len(r)
			Vl = q
		}
	}
	fmt.Println("Winner",Vl,"Len",Max)
	for _,A := range Puntos{
		fmt.Print(Punto{A.X,A.Y},"\t")
		fmt.Println(EqualPoint(Punto{A.X,A.Y},Hough[Vl]))
	}
}

func MCPlane(A Punto,B Punto) Tri{
	X1, Y1 := A.X, A.Y
	X2, Y2 := B.X, B.Y
	if X1 == X2{
		return Tri{X1,0,0}
	}
	if Y1 == Y2{
		return Tri{0,Y1,0}
	}
	V := (Y2 - Y1) / (X2 - X1) 
	a := V//X
	c := (-X1 * V)+Y1//Z
	b := V * X1 + c//Y
	return Tri{float64(int(1000*a))/1000, float64(int(1000*c))/1000, b}
}

func EqualPoint(P1 Punto, P []Punto) bool{
	Result := false
	for _,P2 := range P{
		if P1.X == P2.X && P1.Y == P2.Y{
			return true
		}
	}
	return Result
}

func HoughAdd(MC Tri,XY Punto, H map[Punto][]Punto){ 
	if /*XY.Y == (MC.X * XY.X) +MC.Y && */Exist(XY,H[Punto{MC.X,MC.Y}]) != true{
		H[Punto{MC.X,MC.Y}] = append(H[Punto{MC.X,MC.Y}],XY)
	}
}

func Exist(A Punto, H []Punto) bool{
	for _,B := range H{
		if A.X == B.X && A.Y == B.Y{
			return true
		}
	}
	return false
}
