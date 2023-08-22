package main

import (
	"fmt"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func CreateDateInputWidget(app fyne.App, win fyne.Window) *fyne.Container {

	dateEntry := NewJDateInputWidget(app, win)
	label := widget.NewLabel("Current Selection")
	button := widget.NewButton("Show", func() { on_show_click(dateEntry, label) } )

	helpText := `
	1. Press UP/Down/Left/Right Arrow to change the date.
	2. Space to set current date.
    `
	helpTextWidget := widget.NewTextGridFromString(helpText)

	return container.NewVBox(dateEntry, button, label, widget.NewSeparator(), helpTextWidget)

}

func on_show_click(d *JDateInputWidget, l *widget.Label) {
	var msg string
	if d.GetDate().IsZero() == true {
		msg = "No Input"
	} else {
		msg = fmt.Sprintf("Your input is : %s", d.GetDate().Format(time.DateOnly))
	}
	l.SetText(msg)
}
