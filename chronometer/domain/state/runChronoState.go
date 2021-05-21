package state

import (
	"context"
	"cronometro-go/chronometer/domain"
)

type runChronoState struct {
	currentChrono *domain.Chronometer
	cancelChrono *context.CancelFunc
}

func (this *runChronoState) StartChrono(seg chan<- int, min chan<- int) {
	//No hay nada por aqui
}
func (this *runChronoState) StopChrono() {
	cancel := *this.cancelChrono
	cancel()
	this.currentChrono.SetState(
		&StopChronoState{
			this.currentChrono})
}
func (this *runChronoState) ReloadChrono(seg chan<- int, min chan<- int) {
	//No hay nada por aca
}
func (this *runChronoState) FinishChrono(seg chan<- int, min chan<- int) {
	cancel := *this.cancelChrono
	cancel()
	this.currentChrono.Clean(seg, min)
	this.currentChrono.SetState(
		&initChronoState{
		this.currentChrono, context.Background()})
}