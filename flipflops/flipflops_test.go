package flipflops

import (
	"sync"
	"testing"
	"time"
)

func TestDFlipFlop(t *testing.T) {
	clkChan := make(chan byte)
	dataChan := make(chan byte)
	outputChan := make(chan byte)

	dff := DFlipFlop{dataLine: dataChan, outputLine: outputChan, clkLine: clkChan, state: 0}

	t.Log("Kicking off Flip-Flop")
	dff.Start()

	var wg sync.WaitGroup

	wg.Add(3)

	t.Log("Kicking Off Clock")
	go func() {
		clkChan <- 0
		clkChan <- 1
		clkChan <- 0
		clkChan <- 1
		clkChan <- 0
		clkChan <- 1
		clkChan <- 0
		clkChan <- 1
		clkChan <- 0
		clkChan <- 1
		clkChan <- 0
		wg.Done()
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		dataChan <- 1
		dataChan <- 0
		dataChan <- 0
		dataChan <- 1
		dataChan <- 1
		dataChan <- 1
		dataChan <- 0
		dataChan <- 1
		dataChan <- 1
		dataChan <- 1
		dataChan <- 0
		wg.Done()
	}()

	go func() {
		for {
			select {
			case y := <-outputChan:
				t.Log(y)
			case <-time.After(1):
				t.Log("Times Out ... Quitting")
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()

}
