package application

import (
	"cronometro-go/chronometer/domain"
	"cronometro-go/chronometer/domain/state"
)

type chronoService struct {
	chronometer *domain.Chronometer
}

func NewChronoService () *chronoService {
	var myChrono *domain.Chronometer
	myChrono.SetState(state.NewInitState(myChrono))
	return &chronoService{myChrono}
}

func (this *chronoService) StartChrono(){
	this.chronometer.StartChrono()
}

func (this *chronoService) StopChrono() {
	this.chronometer.StopChrono()
}

func (this *chronoService) ReloadChrono() {
	this.chronometer.ReloadChrono()
}

func (this *chronoService) FinishChrono(){
	this.chronometer.FinishChrono()
}

func setContext(){

}