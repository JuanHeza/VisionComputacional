package main
/*
cls && go run main.go kamala.go f5o.go redtool.go grey.go moda.go
*/

import (
"fmt"
"time"
)

func main() {
	fmt.Println("Hi")
	then := time.Now()
	var IN = "SRC/Input/Weed.jpg"
	//var IN = "SRC/Input/Mala Hierba.jpg"
	var Out = "SRC/Input/Out.jpg"
	GreyMat(IN,Out)
	fmt.Println(IN)
	fmt.Println(Out)
	Kamala(Out, "Output.jpg",25,1)
	A := time.Now().Sub(then)
	fmt.Println("tiempo de ejecucion",A)
}

