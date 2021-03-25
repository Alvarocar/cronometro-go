package model

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
	"time"
)

type Chronometer struct{
	minLabel binding.String
	segLabel binding.String

	start       *widget.Button
	stop        *widget.Button
	reload      *widget.Button
	finish      *widget.Button

	isRunning bool

	currentState ChronoState

	min int
	seg int
	delay int
}

func (this *Chronometer) startChrono () {
	this.currentState.StartChrono()
}

func (this *Chronometer) stopChrono () {
	this.currentState.StopChrono()
}

func (this *Chronometer) reloadChrono() {
	this.currentState.ReloadChrono()
}

func (this *Chronometer) finishChrono(){
	this.currentState.FinishChrono()
}

func (this *Chronometer) SetState(s ChronoState){
	this.currentState = s
}

func (this *Chronometer) RunChrono(){
	segChan := make(chan int)
	minChan := make(chan int)
	go this.chronoRutine(segChan, minChan)
	go this.refreshRutine(segChan, minChan)
}

func (this *Chronometer) chronoRutine(
	segChan chan<- int, minChan chan<- int){
	defer close(segChan)
	defer close(minChan)

	for this.min < 999 {
		for this.seg < 60 {
			if this.isRunning {
				segChan <- this.seg
				time.Sleep(time.Second)
				this.seg++
			} else {
				return
			}
		}
		this.seg = 0
		this.min++
		minChan <- this.min
	}
}

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

func (this *Chronometer) Clean(){
	this.min = 0
	this.seg = 0
	this.segLabel.Set("00")
	this.minLabel.Set("0")
}


func (this *Chronometer) SetIsRunning(isRunning bool){
	this.isRunning = isRunning
}
