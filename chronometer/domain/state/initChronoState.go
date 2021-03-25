package state

import (
	"cronometro-go/chronometer/domain"
)

type InitChronoState struct {
	currentChrono *domain.Chronometer
}

func NewInitState(c *domain.Chronometer) *InitChronoState {
	return &InitChronoState{c}
}

func (this *InitChronoState) StartChrono() {
	this.currentChrono.SetIsRunning(true)
	this.currentChrono.RunChrono()
	this.currentChrono.SetState(
		&RunChronoState{
			this.currentChrono})
}
func (this *InitChronoState) StopChrono() {
	//No hay nada por aqui
}
func (this *InitChronoState) ReloadChrono() {
	//No hay nada por aca
}
func (this *InitChronoState) FinishChrono() {
	//No hay nada tampoco
}
