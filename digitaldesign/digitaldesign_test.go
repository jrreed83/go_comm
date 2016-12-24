package digitaldesign

import (
	"sync"
	"testing"
	"time"
)

func TestDFlipFlop(t *testing.T) {
	clkLine := make(chan byte)
	dataLine := make(chan byte)
	outputLine := make(chan byte)

	circuit := DFlipFlop{dataLine: dataLine, outputLine: outputLine, clkLine: clkLine}

	t.Log("Kicking off Flip-Flop")
	circuit.Start(0)

	var wg sync.WaitGroup

	wg.Add(3)

	t.Log("Kicking Off Clock")
	go func() {
		clkLine <- 0
		clkLine <- 1
		clkLine <- 0
		clkLine <- 1
		clkLine <- 0
		clkLine <- 1
		clkLine <- 0
		clkLine <- 1
		clkLine <- 0
		clkLine <- 1
		clkLine <- 0
		wg.Done()
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		dataLine <- 23
		dataLine <- 23
		dataLine <- 65
		dataLine <- 54
		dataLine <- 41
		dataLine <- 42
		dataLine <- 12
		dataLine <- 12
		dataLine <- 75
		dataLine <- 46
		dataLine <- 34
		wg.Done()
	}()

	go func() {
		for {
			select {
			case y := <-outputLine:
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
