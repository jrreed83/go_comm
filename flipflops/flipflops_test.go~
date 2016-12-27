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

	dff := DFlipFlop{dataLine: dataChan, outputLine: outputChan, clkLine: clkChan}

	t.Log("Kicking off Flip-Flop")
	dff.Start(0)

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
		dataChan <- 23
		dataChan <- 23
		dataChan <- 65
		dataChan <- 54
		dataChan <- 41
		dataChan <- 42
		dataChan <- 12
		dataChan <- 12
		dataChan <- 75
		dataChan <- 46
		dataChan <- 34
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
