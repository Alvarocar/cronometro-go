package application

import "cronometro-go/chronometer/domain"

type ChronoService struct {
	chronometer *domain.Chronometer
}

func (this *ChronoService) StartChrono(){
	this.chronometer.StartChrono()
}

func (this *ChronoService) StopChrono() {
	this.chronometer.StopChrono()
}

func (this *ChronoService) ReloadChrono() {
	this.chronometer.ReloadChrono()
}

func (this *ChronoService) FinishChrono(){
	this.chronometer.FinishChrono()
}