package main

import (
	"cronometro-go/chronometer/infraestructure"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cronometro")

	myWindow.SetContent(
		container.New(layout.NewVBoxLayout(),
			infraestructure.NewChronometer(1),
			infraestructure.NewChronometer(3),
			infraestructure.NewChronometer(5)))

	myWindow.ShowAndRun()
}


