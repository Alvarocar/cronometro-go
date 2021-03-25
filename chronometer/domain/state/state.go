package state

type ChronoState interface {
	StartChrono()
	StopChrono()
	ReloadChrono()
	FinishChrono()
}
