package model

type StopChronoState struct {
	currentChrono *Chronometer
}

func (this *StopChronoState) StartChrono() {
	//No hay nada por aqui
}
func (this *StopChronoState) StopChrono() {
	//No hay nada por aca
}
func (this *StopChronoState) ReloadChrono() {
	this.currentChrono.SetIsRunning(true)
	this.currentChrono.RunChrono()
	this.currentChrono.SetState(
		&RunChronoState{
			this.currentChrono})
}
func (this *StopChronoState) FinishChrono() {
	this.currentChrono.SetIsRunning(false)
	this.currentChrono.Clean()
	this.currentChrono.SetState(
		&InitChronoState{
			this.currentChrono})
}