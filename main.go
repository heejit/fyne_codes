package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MyDateWidget")

	dateEntry := NewMyDateEntry()
	label := widget.NewLabel("Current Selection")
	button := widget.NewButton("Show", func() { on_show_click(dateEntry, label, w) } )

	helpText := `
	1. Press UP/Down Arrow to change the part of date.
	2. Space to set current date.
    3. Delete to clear
    4. Press enter key to update the date
	4. You can enter part of date (like Only day, day-month)
    5. *Assuming first part is Day*
    `
	helpTextWidget := widget.NewTextGridFromString(helpText)
	w.SetContent(container.NewVBox(dateEntry, button, label, widget.NewSeparator(), helpTextWidget))
	w.ShowAndRun()
}

func on_show_click(d *MyDateEntry, l *widget.Label, w fyne.Window) {
	var msg string
	if d.ToDate().IsZero() == true {
		msg = "No Input"
	} else {
		msg = fmt.Sprintf("Your input is : %s", d.ToDate().Format(time.DateOnly))
	}
	l.SetText(msg)
}