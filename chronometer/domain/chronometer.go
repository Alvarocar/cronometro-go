package domain

import (
	"context"
	"cronometro-go/chronometer/domain/state"
	"time"
)

type Chronometer struct{
	currentState state.ChronoState
	min int
	seg int
	delay int
}

func (this *Chronometer) StartChrono(seg chan<- int, min chan<- int) {
	this.currentState.StartChrono(seg, min)
}

func (this *Chronometer) StopChrono() {
	this.currentState.StopChrono()
}

func (this *Chronometer) ReloadChrono(seg chan<- int, min chan<- int) {
	this.currentState.ReloadChrono(seg, min)
}

func (this *Chronometer) FinishChrono(seg chan<- int, min chan<- int){
	this.currentState.FinishChrono(seg, min)
}

func (this *Chronometer) SetState(s state.ChronoState){
	this.currentState = s
}

func (this *Chronometer) RunChrono(
	segChan chan<- int, minChan chan<- int, ctx context.Context){
	go this.chronoRutine(segChan, minChan, ctx)
}

func (this *Chronometer) chronoRutine(
	segChan chan<- int, minChan chan<- int, ctx context.Context){
	defer close(segChan)
	defer close(minChan)

	for this.min < 999 {
		for this.seg < 60 {
			select {
				case <-ctx.Done():
					return
				default:
					if this.seg % this.delay == 0 {
						segChan <- this.seg
					}
					time.Sleep(time.Second)
					this.seg++
			}
		}
		this.seg = 0
		this.min++
		minChan <- this.min
	}
}
/**

func (this *Chronometer) refreshRutine(
	segChan <-chan int, minChan <-chan int){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(segChan <-chan int){
		defer wg.Done()
		for seg := range segChan {

			if seg % this.delay != 0 {
				continue
			}
			if seg < 10 {
				this.segLabel.Set(
					fmt.Sprint("0", seg))
			} else {
				this.segLabel.Set(
					fmt.Sprint(seg))
			}
		}
	}(segChan)

	go func(minChan <-chan int){
		defer wg.Done()
		for min := range minChan {
			this.minLabel.Set(
				fmt.Sprint(min))
		}
	}(minChan)

	wg.Wait()
}

 */
func (this *Chronometer) Clean(seg chan<- int, min chan<- int){
	this.min = 0
	this.seg = 0
	seg<- this.seg
	min <-this.min
}