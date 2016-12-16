package digitaldesign

import (
	"sync"
	"testing"
	"time"
)

func TestDFF(t *testing.T) {
	clk := make(chan byte)
	data := make(chan byte)
	output := make(chan byte)

	dff := DFF{data_port: data, output_port: output, clk_port: clk}

	t.Log("Kicking off Flip-Flop")
	dff.Start()

	var wg sync.WaitGroup

	wg.Add(3)

	t.Log("Kicking Off Clock")
	go func() {
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		wg.Done()
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		data <- 23
		data <- 23
		data <- 65
		data <- 54
		data <- 41
		data <- 42
		data <- 12
		data <- 12
		data <- 75
		data <- 46
		data <- 34
		wg.Done()
	}()

	go func() {
		for {
			select {
			case y := <-output:
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

func TestClockEvent(t *testing.T) {
	// Initialize Clock Channel
	clk := make(chan uint8)

	event := rising_edge(clk)

	t.Log("Kicking off Clock")

	go func() {
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 1
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
	}()

	go func() {
		for {
			select {
			case evt := <-event:
				if evt {
					t.Log("y")
				} else {
					t.Log("n")
				}
			case <-time.After(1):
				t.Log("Timed Out ... Quitting")
				return
			}
		}

	}()

	time.Sleep(1)

}
