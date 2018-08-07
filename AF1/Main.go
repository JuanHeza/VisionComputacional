// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"time"
	"image"
	"image/draw"	
	"strconv"
	_ "image/jpeg"
	_ "image/png"
	"fmt"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

var Selection,file string

type SuperData struct{
	In, Out, Hst, Mat, Ori string
}
var TMNT SuperData

type OutLayout struct{
	img1, img2, img3 gxui.Image
	labelO, labelN gxui.Label
}

var inf, sup int = 0,0

// Number picker uses the gxui.DefaultAdapter for driving a list
func numberPicker(theme gxui.Theme, overlay gxui.BubbleOverlay, window gxui.Window, driver gxui.Driver) gxui.Control {
	var items,GOs []string
	//Data := make(chan bool)
	go func(){
		items=GetFiles("./SRC/Input",true) 
	}()
	go func(){
		GOs=GetFiles(".",false) 
	}()
	time.Sleep(1 * time.Second)
	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(items)

	adapterD := gxui.CreateDefaultAdapter()
	adapterD.SetItems(GOs)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	layout1 := theme.CreateLinearLayout()
	layout1.SetDirection(gxui.TopToBottom)
	layout2 := theme.CreateLinearLayout()
	layout2.SetDirection(gxui.TopToBottom)
	layout2.SetVisible(false)
	layout2.SetHorizontalAlignment(1)

	label0 := theme.CreateLabel()
	label0.SetText("Actions:")
	layout1.AddChild(label0)

	dropList := theme.CreateDropDownList()
	dropList.SetAdapter(adapterD)
	dropList.SetBubbleOverlay(overlay)
	layout1.AddChild(dropList)

	Textos := theme.CreateLinearLayout()
	Textos.SetDirection(gxui.TopToBottom)

	labelM := theme.CreateLabel()
	labelM.SetText("Input Data:")
	textoM := theme.CreateTextBox()
	textoM.SetText("1")
	
	Textos.AddChild(labelM)
	Textos.AddChild(textoM)
	Textos.SetVisible(false)
	layout1.AddChild(Textos)

	label1 := theme.CreateLabel()
	label1.SetText("Images:")
	layout1.AddChild(label1)

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(gxui.Vertical)
	layout1.AddChild(list)

	dropList.OnSelectionChanged(func(item gxui.AdapterItem) {
		if item != nil {
			Selection = fmt.Sprintf("%s",item)
			Val,Exist := GetTitle(Selection)
			fmt.Println(Val)			
			fmt.Println(Selection,":",Exist)
			if(Exist==true){
				labelM.SetText(Val.Titulo)
				textoM.SetText(strconv.Itoa(Val.inf))
				inf = Val.inf
				sup = Val.sup
				Textos.SetVisible(true)
			}else{
				textoM.SetText("0")
				Textos.SetVisible(false)
			}
		}
	})

	button := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action()})
		layout1.AddChild(b)
		return b
	}

	button("Aceptar",
		func() { 
			then := time.Now()
			Input := textoM.Text()
			Digit,err := strconv.Atoi(Input)
			if (err != nil){
				textoM.SetText("No Es Un Numero")
			}else if(Digit > sup || Digit < inf){
				textoM.SetText("Valor Fuera De Los Limites")
			}else{
			fmt.Println("ALL OK")
			layout2.RemoveAll()
			TMNT.Out, TMNT.Hst, TMNT.Mat, TMNT.Ori = Ejecution(Digit,Selection,TMNT.In, layout2)
			if (TMNT.In != ""){
					img1 := theme.CreateImage()
					img2 := theme.CreateImage()
					img3 := theme.CreateImage()
					labelO := theme.CreateLabel()
					labelN := theme.CreateLabel()
					DougOut := OutLayout{img1,img2,img3,labelO,labelN}
					layout2 = OutputLayout(theme,layout2, TMNT,DougOut)
					fmt.Println("Cambiando Datos")
				}else{
					fmt.Println("Faltan Datos")
				}
			fmt.Println("Tiempo De Ejecucion",time.Now().Sub(then))
		}
		},
	)

	label2 := theme.CreateLabel()
	label2.SetMargin(math.Spacing{T: 20})
	label2.SetText("Selected Image:")
	layout1.AddChild(label2)

	selected1 := theme.CreateImage()
	selected1.SetExplicitSize(math.Size{W: 220, H: 220})
	selected1.SetAspectMode(1)
	layout1.AddChild(selected1)

	Close := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action()})
		layout1.AddChild(b)
		return b
	}


	Close("Cerrar",
		func() { 
			fmt.Println("BYE")
			window.Close()
		},
	)

	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		if item != nil {
			file = fmt.Sprintf("./SRC/Input/%s",item)
			TMNT.In=file
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("Failed to Select image '%s': %v\n", file, err)
				os.Exit(1)
			}

			source, _, err := image.Decode(f)
			
			// Copy the image to a RGBA format before handing to a gxui.Texture
			rgba := image.NewRGBA(source.Bounds())
			draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
			texture := driver.CreateTexture(rgba, 1)
			selected1.SetTexture(texture)
		}
	})

	layout.AddChild(layout1)
	layout.AddChild(layout2)
	return layout
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	window := theme.CreateWindow(1200, 600, "Vision")
	holder := theme.CreatePanelHolder()
	holder.AddPanel(numberPicker(theme, overlay, window, driver), "Z-626")

	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	//window.SetFullscreen(true)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}

func Ejecution(Q int,Program string, file string, layout1 gxui.LinearLayout) (string,string,string, string){
	fmt.Println("Program:",Program,"Image:",file)
	VR := GetNames()
	DataOk:=true
	switch(Program){
		case "Umbral Base 2":
			Bin2(file,VR.I)
			break
		case "Umbral Base 3":
			Bin3(file,VR.I)
			break
		case "Umbral Base 4":
			Bin4(file,VR.I)
			break
		case "Umbral Base 6":
			Bin6(file,VR.I)
			break
		case "Umbral Base 8":
			Bin8(file,VR.I)
			break
		case "Umbral Base 12":
			Bin12(file,VR.I)
			break
		case "Umbral Base 16":
			Bin16(file,VR.I)
			break
		case "Negativo":
			Nega(file,VR.I)
			break
		case "Oscurecimiento A":
			GreyScale(file,VR.I)
			Darkness(VR.I,VR.I)
			break
		case "Oscurecimiento B":
			GreyScale(file,VR.I)
			Dark(VR.I,VR.I)
			break
		case "Aclarado A":
			GreyScale(file,VR.I)
			Light(VR.I,VR.I)
			break
		case "Aumentar Brillo":
			GreyScale(file,VR.I)
			SatP(VR.I,VR.I)
			break
		case "Minimizar Brillo":
			GreyScale(file,VR.I)
			SatM(VR.I,VR.I)
			break
		case "Escala a Grises":
			GreyScale(file,VR.I)
			break
		case "Espejado":
			Mirror(file,VR.I)
			break
		case "Shifting-Oeste":
			SO(file,VR.I)
			break
		case "Shifting-Este":
			SE(file,VR.I)
			break
		case "Shifting-Norte":
			SN(file,VR.I)
			break
		case "Shifting-Sur":
			SS(file,VR.I)
			break
		case "Shifting-NorOeste":
			SNO(file,VR.I)
			break
		case "Shifting-SurEste":
			SSE(file,VR.I)
			break
		case "Shifting-NorEste":
			SNE(file,VR.I)
			break
		case "Shifting-SurOeste":
			SSO(file,VR.I)
			break
		case "Normalizar Minimo":
			NorMi(file,VR.I)
			break
		case "Normalizar Maximo":
			NorMa(file,VR.I)
			break
		case "Convolucion":
			Convolucion(file,VR.I)
			break
		case "Filtro Mediana":
			FiltroMediana(file,VR.I)
			break
		case "Filtro Media":
			FiltroMedia(file,VR.I)
			break
		case "Filtro Moda":
			FiltroModa(file,VR.I)
			break
		case "Filtro Orden de Rango":
			FiltroOrdenRango(file,VR.I,Q)
			break
		case "Contraste":
			Contraste(file,VR.I,float64(Q))
			break
		case "Ruido Gaussiano":
			Gaussiano(file,VR.I)
			break
		case "Umbral Adaptativo":
			UmbralAdap(file,VR.I,uint8(Q))
			break
		case "Aclarado B":
			Aclarar(file,VR.I,uint8(Q))
			break
		default: 
			fmt.Println("NONE")
			DataOk = false
	}
	if DataOk ==false{
		return "./SRC/Data/Dummy.jpg","./SRC/Data/Dummy.jpg","./SRC/Data/Dummy.txt","./SRC/Data/Dummy.txt"
	}
	fmt.Println("Action",Program)
	layout1.SetVisible(true)
	go func(){ VR.O = GreyMat(file,VR.O) }()
	time.Sleep(2 * time.Second)
	go func(){ VR.M = GreyMat(VR.I,VR.M) }()
	go Histo(VR.I,VR.H)
	time.Sleep(3 * time.Second)
	return	VR.I, VR.H, VR.M, VR.O
}

func OutputLayout(theme gxui.Theme, layout gxui.LinearLayout, TMNT SuperData, DougOut OutLayout) gxui.LinearLayout {
	layout1 := theme.CreateLinearLayout()
	layout1.SetDirection(gxui.LeftToRight)
	layout2 := theme.CreateLinearLayout()
	layout2.SetDirection(gxui.LeftToRight)
	layout3 := theme.CreateLinearLayout()
	layout3.SetDirection(gxui.TopToBottom)
	DougOut.img1 = Genimage(theme, TMNT.In)
	DougOut.img2 = Genimage(theme, TMNT.Out)
	DougOut.img3 = Genimage(theme, TMNT.Hst)
	layout1.AddChild(DougOut.img1)
	layout1.AddChild(DougOut.img2)
	layout1.AddChild(DougOut.img3)
	
	layout.AddChild(layout1)

	layout21 := theme.CreateScrollLayout()
	layout21.SetScrollAxis(true,true)
	layout22 := theme.CreateScrollLayout()
	layout22.SetScrollAxis(true,true)
	

	fillLabel(DougOut.labelO,TMNT.Ori)

	fillLabel(DougOut.labelN,TMNT.Mat)
	
	holderO := theme.CreatePanelHolder()
	layout21.SetChild(DougOut.labelN)
	layout22.SetChild(DougOut.labelO)

	holderO.AddPanel(layout21, "Matriz Nueva")
	holderO.AddPanel(layout22, "Matriz original")
	
	layout3.AddChild(holderO)
	layout.AddChild(layout2)
	layout.AddChild(layout3)
	return layout
}

func fillLabel(label gxui.Label,In string) {
	label.SetText(In)
	label.SetVerticalAlignment(gxui.AlignMiddle)
	label.SetHorizontalAlignment(gxui.AlignCenter)
	label.SetMultiline(true)
}

func Genimage(theme gxui.Theme, dir string) gxui.Image{
	file := dir
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Failed to open image '%s': %v\n", file, err)
		os.Exit(1)
	}

	source, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Failed to read image '%s': %v\n", file, err)
		os.Exit(1)
	}
	img := theme.CreateImage()
	driver := theme.Driver()
	// Copy the image to a RGBA format before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)
	img.SetExplicitSize(math.Size{W: 256, H: 256})
	img.SetMargin(math.Spacing{L: 10, T: 10, R: 10, B: 10})
	img.SetAspectMode(1)
	return img

}