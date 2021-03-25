package infraestructure

import (
	"cronometro-go/chronometer/domain"
	"cronometro-go/chronometer/domain/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//Create a new Chronometer
func NewChronometer(delay int) *fyne.Container {

	var myChrono domain.Chronometer

	myChrono.SetState(state.NewInitState(&myChrono))

	myChrono.minLabel = binding.NewString()
	myChrono.segLabel = binding.NewString()
	myChrono.minLabel.Set("0")
	myChrono.segLabel.Set("00")

	if delay <= 0 {
		delay = 1
	}
	myChrono.delay = delay
	//Chronometer Label
	contentChronoLabel := container.New(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			widget.NewLabelWithData(myChrono.minLabel),
			widget.NewLabel(":"),
			widget.NewLabelWithData(myChrono.segLabel),
			layout.NewSpacer())

	//Chronometer Buttons
	myChrono.start = widget.NewButtonWithIcon(
		"Iniciar", theme.MediaPlayIcon(),
		myChrono.startChrono)

	myChrono.stop = widget.NewButtonWithIcon(
		"Pausar", theme.MediaPauseIcon(),
		myChrono.stopChrono)

	myChrono.reload = widget.NewButtonWithIcon(
		"Reanudar", theme.MediaReplayIcon(),
		myChrono.reloadChrono)

	myChrono.finish = widget.NewButtonWithIcon(
		"Terminar", theme.MediaStopIcon(),
		myChrono.finishChrono)

	buttons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), myChrono.start, myChrono.stop, myChrono.reload, myChrono.finish, layout.NewSpacer())

	return container.New(layout.NewVBoxLayout(),contentChronoLabel, buttons)
}
