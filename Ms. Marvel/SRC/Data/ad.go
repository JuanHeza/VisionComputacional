package main

import "fmt"
import "log"
import "image/Color"

var Max = 5
var BEEP,X,END chan int
func main(){/*
	END = make(chan int, 1)
	X = make(chan int, 100)
    for i:=0;i<100;i++{
		X <- i
	}
	Conteo()
//	time.Sleep(1000)
	<-END
	fmt.Println("BYE")
}

func Conteo ( ) {
	BEEP = make(chan int, 4)
	for len(X) > 0{ 
		go func(){
			BEEP <- 1
			log.Println(<-X)
			<- BEEP
		}()
	}
	END <- 0
	*/
}
