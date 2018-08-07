package main

import (
	//"strconv"
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func GenCogMix(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	labelM := theme.CreateLabel()
	labelM.SetText("Modulo \"M\":")
	textoM := theme.CreateTextBox()
	textoM.Text()
	
	layout.AddChild(labelM)
	layout.AddChild(textoM)
	

	button := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action()})
		layout.AddChild(b)
		return b
	}

	Final := theme.CreateLabel()
	button("Aceptar",
		func() { fmt.Println("ALL OK")/*Res=Mixto(textoA.Text(),textoX.Text(),textoC.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
			adapter.SetItems(Res)
			list.SetAdapter(adapter)
			list.SetOrientation(gxui.Vertical)*/
		},
	)
	layout.AddChild(Final)
	fmt.Println(Final.Color())
	
	resultado := theme.CreateScrollLayout ()
	resultado.SetScrollAxis(false,true)
	//resultado.SetChild(list)
	table := theme.CreateTableLayout()
	table.SetGrid(7,4) // columns, rows
	table.SetChildAt(0, 0, 1, 4, layout)
	table.SetChildAt(1, 0, 2, 2, resultado)
//	table.SetChildAt(1, 2, 2, 2, resultado)
//	table.SetChildAt(3, 0, 4, 2, resultado)
//	table.SetChildAt(3, 2, 4, 2, resultado)
	return table
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
	holder.AddPanel(GenCogMix(theme, overlay), "Exponencial")

	window := theme.CreateWindow(800, 600, "Lists")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}