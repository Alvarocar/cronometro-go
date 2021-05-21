package state

import (
	"context"
	"cronometro-go/chronometer/domain"
)

type initChronoState struct {
	currentChrono *domain.Chronometer
	ctx context.Context
}

func NewInitState(c *domain.Chronometer) *initChronoState {
	return &initChronoState{c, context.Background()}
}

func (this *initChronoState) StartChrono(seg chan<- int, min chan<- int) {
	ctx, cancel := context.WithCancel(this.ctx)
	this.currentChrono.RunChrono(seg, min, ctx)
	this.currentChrono.SetState(
		&runChronoState{
			this.currentChrono, &cancel})
}

func (this *initChronoState) StopChrono() {}
func (this *initChronoState) ReloadChrono(seg chan<- int, min chan<- int) {}
func (this *initChronoState) FinishChrono(seg chan<- int, min chan<- int) {}
