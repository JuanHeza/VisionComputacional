package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main (){
	_, err := exec.LookPath("SDoo.exe")
	if err != nil {
		log.Fatal("File Not Found")
	}else{
		var Data = []string {"1","2","3"}
		fmt.Println(Data)
		cmd := exec.Command("SDoo",Data...)
	
		log.Printf("Running command and waiting for it to finish...")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
	}
}