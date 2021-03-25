package state

import (
	"cronometro-go/chronometer/domain"
)

type RunChronoState struct {
	currentChrono *domain.Chronometer
}

func (this *RunChronoState) StartChrono() {
	//No hay nada por aqui
}
func (this *RunChronoState) StopChrono() {
	this.currentChrono.SetIsRunning(false)
	this.currentChrono.SetState(
		&StopChronoState{
			this.currentChrono})
}
func (this *RunChronoState) ReloadChrono() {
	//No hay nada por aca
}
func (this *RunChronoState) FinishChrono() {
	this.currentChrono.SetIsRunning(false)
	this.currentChrono.Clean()
	this.currentChrono.SetState(
		&InitChronoState{
		this.currentChrono})
}