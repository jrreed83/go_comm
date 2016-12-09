package digitalcircuits

import (
	"testing"
	"time"
)

func TestDFlipFlop(t *testing.T) {
	clk := make(chan uint8)
	data := make(chan uint8)
	output := make(chan uint8)

	t.Log("Kicking off Flip Flop")
	go d_flip_flop(clk, data, output)

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
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		//	clk <- 0
		data <- 23
		//	clk <- 1
		data <- 23
		//	clk <- 0
		data <- 65
		//	clk <- 1
		data <- 65
		//	clk <- 0
		data <- 41
		//	clk <- 1
		data <- 42
		//	clk <- 0
		data <- 12
		//	clk <- 1
		data <- 12
		//	clk <- 0
		data <- 75
		//	clk <- 1
		data <- 46
		//	clk <- 0
		data <- 34
		//	clk <- 1
	}()

	time.Sleep(1)

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
	time.Sleep(2)
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
				//case <-time.After(1):
				//	t.Log("Timed Out ... Quitting")
				//	return
			}
		}

	}()

	time.Sleep(1)

}
