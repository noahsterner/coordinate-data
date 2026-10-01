package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const (
	ScreenWidth = 800
	ScreenHeight = 800
)

func main() {
	a := app.New()
	window := a.NewWindow("Canvas")
	
	window.SetContent(widget.NewLabel("Hello, World"))
	window.ShowAndRun()
}
