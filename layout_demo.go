package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func CreateHBoxRatioLayoutWidget() *fyne.Container {
	lbl1 := widget.NewLabel("Name:")
	entry1 := widget.NewEntry()
	lbl2 := widget.NewLabel("Number:")
	entry2 := widget.NewEntry()
	btn  := widget.NewButton("okay", nil)
	content := container.New(NewHBoxRatioLayout([]float32{10, 20, 10, 20, 5, 15}), lbl1, entry1, lbl2, entry2, widget.NewSeparator(), btn)
	return content
}
