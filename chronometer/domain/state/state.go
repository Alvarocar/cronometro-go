package state

type ChronoState interface {
	StartChrono(seg chan<- int, min chan<- int)
	StopChrono()
	ReloadChrono(seg chan<- int, min chan<- int)
	FinishChrono(seg chan<- int, min chan<- int)
}
