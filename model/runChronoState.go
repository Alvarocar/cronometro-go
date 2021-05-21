package model

type RunChronoState struct {
	currentChrono *Chronometer
}

func (this *RunChronoState) StartChrono() {}

func (this *RunChronoState) StopChrono() {
	this.currentChrono.SetIsRunning(false)
	this.currentChrono.SetState(
		&StopChronoState{
			this.currentChrono})
}
func (this *RunChronoState) ReloadChrono() {}

func (this *RunChronoState) FinishChrono() {
	this.currentChrono.SetIsRunning(false)
	this.currentChrono.Clean()
	this.currentChrono.SetState(
		&InitChronoState{
			this.currentChrono})
}
