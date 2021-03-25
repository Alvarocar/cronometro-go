package model

type ChronoState interface {
	StartChrono()
	StopChrono()
	ReloadChrono()
	FinishChrono()
}
