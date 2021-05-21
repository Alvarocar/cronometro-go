package model

type InitChronoState struct {
	currentChrono *Chronometer
}

func NewInitState(c *Chronometer) *InitChronoState {
	return &InitChronoState{c}
}

func (this *InitChronoState) StartChrono() {
	this.currentChrono.SetIsRunning(true)
	this.currentChrono.RunChrono()
	this.currentChrono.SetState(
		&RunChronoState{
			this.currentChrono})
}

func (this *InitChronoState) StopChrono() {}

func (this *InitChronoState) ReloadChrono() {}

func (this *InitChronoState) FinishChrono() {}
