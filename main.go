package main

import (
	"cronometro-go/model"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cronometro")

	myWindow.SetContent(
		container.New(layout.NewVBoxLayout(),
			model.NewChronometer(1),
			model.NewChronometer(3),
			model.NewChronometer(5)))

	myWindow.ShowAndRun()
}


