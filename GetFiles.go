package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"io/ioutil"
)

var functions map[string][]string
var Titles map[string]MMPR 
var OutData map[string]int

type NewGen	struct{
	I,M,H,O string
}

type MMPR struct{
	Titulo string
	inf, sup int
}

func GetNames() NewGen{
	OutData := make(map[string]int)
	D :="./SRC/Output"
	files, err := ioutil.ReadDir(D)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if (!file.IsDir()){
			QR := strings.Split(file.Name(),".")
			if strings.Contains(QR[0],"-"){
				X:=strings.Split(QR[0],"-")
					switch(X[0]){
						case "OutI":
							OutData["OutI"], err = strconv.Atoi(X[1])
							break
						case "OutM":
							OutData["OutM"], err = strconv.Atoi(X[1])
							break
						case "OutH":
							OutData["OutH"], err = strconv.Atoi(X[1])
							break
						case "OutO":
							OutData["OutO"], err = strconv.Atoi(X[1])
							break
					}
				if err != nil {
					log.Panicln(err)
				}
			}
		}
	}
	Max := 0 
	for _,DR := range OutData{
		if DR > Max{
			Max = DR
		}
	}
	Max++
	var VR NewGen
	// Imagen de Salida
	VR.I =	fmt.Sprintf("SRC/Output/OutI-%d.jpg",Max)
	//Matriz Original
	VR.O =	fmt.Sprintf("SRC/Output/OutO-%d.txt",Max)
	//Matriz de Salida
	VR.M = fmt.Sprintf("SRC/Output/OutM-%d.txt",Max)
	//Histograma
	VR.H = fmt.Sprintf("SRC/Output/OutH-%d.jpg",Max)
	return VR
}

func GetTitle(prg string) (MMPR, bool){
	Titles = map[string]MMPR{
		"Aclarado":MMPR{Titulo:" Nivel de brillo: (1-255)", inf:1, sup:255},
		"Contraste":MMPR{Titulo:"Nivel de contraste: (1-3)", inf:1, sup:3},
		"Umbral Binario":MMPR{Titulo:"Nivel de umbral a trabajar (1-255)", inf:1, sup:255},
		"Filtro Orden de Rango":MMPR{Titulo:"Rango a Utilizar (0-8)", inf:0, sup:8},
	}
	X,K := Titles[prg]
	return X,K
}

func GetFiles(dir string,image bool) []string{
	functions = map[string][]string{
		"normal":{"Normalizacion Maxima","Normalizacion Minima"},
		"Bin2":{"Umbral Base 2"},
		"Bin3":{"Umbral Base 3"},
		"Bin4":{"Umbral Base 4"},
		"Bin6":{"Umbral Base 6"},
		"Bin8":{"Umbral Base 8"},
		"Bin12":{"Umbral Base 12"},
		"Bin16":{"Umbral Base 16"},
		"Dark":{"Oscurecimiento A"},
		"Light":{"Oscurecimiento B","Aclarado A","Aumentar Brillo","Minimizar Brillo"},
		"Greyscale":{"Escala a Grises"},
		"Mirror":{"Espejado"},
		"Nega":{"Negativo"},
		"Programas":{"Aclarado B","Contraste","Umbral Adaptativo","Ruido Gaussiano","Filtro Mediana","Filtro Media","Filtro Moda","Filtro Orden de Rango","Convolucion"},
		"Shift":{"Shifting-Norte","Shifting-Sur","Shifting-Oeste","Shifting-Este","Shifting-NorOeste","Shifting-SurEste","Shifting-NorEste","Shifting-SurOeste"},
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var items []string
	for _, file := range files {
		if (!file.IsDir()){
			Exclude := true
			if (!strings.Contains(file.Name(),"Histograma.go")==false){
				Exclude = false
			}else if (!strings.Contains(file.Name(),"colorful.go")==false){
				Exclude = false
			}else if (!strings.Contains(file.Name(),"GetFiles.go")==false){
				Exclude = false
			}else if (!strings.Contains(file.Name(),"Grey.go")==false){
				Exclude = false
			}else if (!strings.Contains(file.Name(),"desktop.ini")==false){
				Exclude = false
			}

			if(Exclude != false){
				if (!image){
					QR := strings.Split(file.Name(),".")
					Rosa := functions[QR[0]]
					for _,X := range Rosa{
						items = append(items,X)
					}
				}else{
					items = append(items,file.Name())
				}
			}
		}
	}
	return items
}