package digitalcircuits

import (
	"testing"
	"time"
)

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
