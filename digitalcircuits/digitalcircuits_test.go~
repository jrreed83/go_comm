package digitalcircuits

import "testing"

func TestClockEvent(t *testing.T) {
	// Initialize Clock Channel
	clk := make(chan uint8)

	event := rising_edge(clk)

	t.Log("Kicking off Clock")

	go func() {
		clk <- 0
		y := <-event
		t.Log(y)

		clk <- 1
		y = <-event
		t.Log(y)

	}()

}
