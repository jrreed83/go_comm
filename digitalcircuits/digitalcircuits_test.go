package digitalcircuits

import (
	"testing"
	"time"
)

func TestDFlipFlop(t *testing.T) {
	clk := make(chan uint8)
	data := make(chan uint8)
	output := make(chan uint8)

	t.Log("Kicking Off Clock")
	go func() {
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
		clk <- 0
		clk <- 1
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		data <- 23
		data <- 45
		data <- 65
		data <- 52
		data <- 12
		data <- 10
	}()

	t.Log("Kicking off flip flop")
	go d_flip_flop(clk, data, output)

	go func() {
		for {
			select {
			case y := <-output:
				t.Log(y)
			case <-time.After(1):
				t.Log("Times Out ... Quitting")
				return
			}
		}
	}()

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
			case <-event:
				t.Log("rising_edge")
			case <-time.After(1):
				t.Log("Timed Out ... Quitting")
				return
			}
		}

	}()

	time.Sleep(1)

}
