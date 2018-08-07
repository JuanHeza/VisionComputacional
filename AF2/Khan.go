package main
/*
cls && go build Kamala.go F5O.go Khan.go getfiles.go umbrales.go
*/
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
	In, Out string
}
var TMNT SuperData
var progressBar gxui.ProgressBar

type OutLayout struct{
	img1, img2 gxui.Image
}

var inf, sup int = 0,0

// Number picker uses the gxui.DefaultAdapter for driving a list
func numberPicker(theme gxui.Theme, overlay gxui.BubbleOverlay, window gxui.Window, driver gxui.Driver) gxui.Control {
	var items  []string
	var GOs = []string {"Semejanza","Discontinuiodad"}
	//Data := make(chan bool)
	go func(){
		items=GetFiles("./SRC/Input",true) 
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

	progressBar = theme.CreateProgressBar()
	progressBar.SetDesiredSize(math.Size{W: 400, H: 20})
	progressBar.SetTarget(100)
	
	layout2.AddChild(progressBar)
	layout2.SetVisible(true)
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
	labelM.SetText("AI:")
	textoM := theme.CreateTextBox()
	textoM.SetText("1")
	
	Textos.AddChild(labelM)
	Textos.AddChild(textoM)
	Textos.SetVisible(true)
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
			}else{
			fmt.Println("ALL OK")
			layout2.RemoveAll()
			driver.Call(func() {
				progressBar.SetProgress(0)
			})
			TMNT.Out = Ejecution(Digit,Selection,TMNT.In, layout2)
			if (TMNT.In != ""){
					img1 := theme.CreateImage()
					img2 := theme.CreateImage()
					DougOut := OutLayout{img1,img2}
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
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}

func Ejecution(Q int,Program string, file string, layout1 gxui.LinearLayout) (string){
	fmt.Println("Program:",Program,"Image:",file)
	VR := GetNames()
	DataOk:=true
	switch(Program){
		case "Discontinuiodad":
			fmt.Println("Discontinuiodad")
			break
		case "Semejanza":
			fmt.Println("Semejanza")
			Kamala(file, VR.I,Q,1)
			break
		
		default: 
			fmt.Println("NONE")
			DataOk = false
	}
	if DataOk == false{
		return "./SRC/Data/Dummy.jpg"
	}
	fmt.Println("Action",Program)
	layout1.SetVisible(true)
	return	VR.I	
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
	layout1.AddChild(DougOut.img1)
	layout1.AddChild(DougOut.img2)
	
	layout.AddChild(layout1)

	layout21 := theme.CreateScrollLayout()
	layout21.SetScrollAxis(true,true)
	layout22 := theme.CreateScrollLayout()
	layout22.SetScrollAxis(true,true)

	layout.AddChild(layout3)
	layout.AddChild(layout2)
	return layout
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

func Progress(X float64){
	progress := int(X)
		progressBar.SetProgress(progress)
}