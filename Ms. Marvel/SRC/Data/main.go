package main

import (
"fmt"
"time"
)

func main() {
	fmt.Println("Hi")
	then := time.Now()
	Kamala("a.jpg", "Output.jpg",50,1)
	A := time.Now().Sub(then)
  fmt.Println("tiempo de ejecucion",A)
  //fmt.Printf("Porcentaje del tiempo base %3.2f%s",float64(A)*100/float64(31 * time.Minute),"%")
	fmt.Printf("Porcentaje del tiempo base %3.2f%s",float64(A)*100/float64(54 * time.Second),"%")
}

