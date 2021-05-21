package model

type StopChronoState struct {
	currentChrono *Chronometer
}

func (this *StopChronoState) StartChrono() {}

func (this *StopChronoState) StopChrono() {}

func (this *StopChronoState) ReloadChrono() {
	this.currentChrono.SetIsRunning(true)
	this.currentChrono.RunChrono()
	this.currentChrono.SetState(
		&RunChronoState{
			this.currentChrono})
}
func (this *StopChronoState) FinishChrono() {
	this.currentChrono.Clean()
	this.currentChrono.SetState(
		&InitChronoState{
			this.currentChrono})
}
