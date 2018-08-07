package main

import (
	"fmt"
    "os"
)

func main(){
	// CMD : X.exe In.jpg Out.jpg Accion
	//Archivo de entrada
	In := os.Args[1]
	//Archivo de salida
	Out := os.Args[2]
	//Seleccion de programa
	Select := os.Args[3]
	fmt.Println("Program:",Select,"Image:",In,"Output",Out)/*
	switch(Select){
		case "Umbral Base 2":
			Bin2(In,Out)
			fmt.Println("Umbral Base 2")
			break
		case "Umbral Base 3":
			Bin3(In,Out)
			fmt.Println("Umbral Base 3")
			break
		case "Umbral Base 4":
			Bin4(In,Out)
			fmt.Println("Umbral Base 4")
			break
		case "Umbral Base 6":
			Bin6(In,Out)
			fmt.Println("Umbral Base 6")
			break
		case "Umbral Base 8":
			Bin8(In,Out)
			fmt.Println("Umbral Base 8")
			break
		case "Umbral Base 12":
			Bin12(In,Out)
			fmt.Println("Umbral Base 12")
			break
		case "Umbral Base 16":
			Bin16(In,Out)
			fmt.Println("Umbral Base 16")
			break
		case "Negativo":
			Nega(In,Out)
			fmt.Println("Negativo")
			break
		case "Oscurecimiento A":
			Darkness(In,Out)
			fmt.Println("Oscurecimiento A")
			break
		case "Oscurecimiento B":
			Dark(In,Out)
			fmt.Println("Oscurecimiento B")
			break
		case "Aclarado":
			Light(In,Out)
			fmt.Println("Aclarado")
			break
		case "Aumentar Brillo":
			SatP(In,Out)
			fmt.Println("Aumentar Brillo")
			break
		case "Minimizar Brillo":
			SatM(In,Out)
			fmt.Println("Minimizar Brillo")
			break
		case "Escala a Grises":
			GreyScale(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Espejado":
			Mirror(In,Out)
			fmt.Println("Espejado")
			break
		case "Shifting-Oeste":
			SO(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-Este":
			SE(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-Norte":
			SN(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-Sur":
			SS(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-NorOeste":
			SNO(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-SurEste":
			SSE(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-NorEste":
			SNE(In,Out)
			fmt.Println("Escala a Grises")
			break
		case "Shifting-SurOeste":
			SSO(In,Out)
			fmt.Println("Escala a Grises")
			break
		default: 
			fmt.Println("NONE")
	}*/
}