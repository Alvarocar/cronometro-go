package state

import (
	"context"
	"cronometro-go/chronometer/domain"
)

type StopChronoState struct {
	currentChrono *domain.Chronometer
}

func (this *StopChronoState) StartChrono(seg chan<- int, min chan<- int) {}

func (this *StopChronoState) StopChrono() {}

func (this *StopChronoState) ReloadChrono(seg chan<- int, min chan<- int) {
	ctx, cancel := context.WithCancel(context.Background())
	this.currentChrono.RunChrono(seg, min, ctx)
	this.currentChrono.SetState(
		&runChronoState{
			this.currentChrono, &cancel})
}
func (this *StopChronoState) FinishChrono(seg chan<- int, min chan<- int) {
	this.currentChrono.Clean(seg, min)

	this.currentChrono.SetState(
		&initChronoState{
			this.currentChrono,
			context.Background()})
}