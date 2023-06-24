package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("HBoxLayoutRatio")

	content := container.NewAppTabs()
	content.SetTabLocation(container.TabLocationBottom)
	content.Append(container.NewTabItem("HBoxRatioLayout", CreateHBoxRatioLayoutWidget()))
	content.Append(container.NewTabItem("DateInputWidget", CreateDateInputWidget()))

	w.SetContent(content)
	w.ShowAndRun()
}

