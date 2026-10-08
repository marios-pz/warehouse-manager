package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Warehouse Manager")
	w.SetContent(widget.NewLabel("Warehouse Manager"))
	w.ShowAndRun()
}
