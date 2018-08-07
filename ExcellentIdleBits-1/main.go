package main
/*
cls && go run main.go kamala.go f5o.go redtool.go grey.go moda.go Histograma.go
*/

import (
  "fmt"
"time"
)

func main() {
	fmt.Println("Hi")
	then := time.Now()
	var IN = "SRC/Input/Weed.jpg"
	var Out = "SRC/Output/Output.jpg"
	K := GreyMat(IN)
	fmt.Println(IN)
	fmt.Println(Out)
	Kamala(K,Out ,25,1)
  fmt.Println("Salida",Out)
	A := time.Now().Sub(then)
	fmt.Println("tiempo de ejecucion",A)
}