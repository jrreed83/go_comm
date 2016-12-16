package digitaldesign

type DFF struct {
	data_port   chan byte
	output_port chan byte
	clk_port    chan byte
}

func (d *DFF) Start() {
	go func() {

		var state byte = 0

		// Look for rising-edge event
		for is_rising_edge := range rising_edge(d.clk_port) {

			// Event based logic
			if is_rising_edge {
				state = <-d.data_port
			} else {
				<-d.data_port // drop on the floor
			}

			// Put data onto output line
			d.output_port <- state
		}
	}()

}

func rising_edge(input_chan chan byte) chan bool {

	output_chan := make(chan bool)

	go func() {
		var event bool = false
		var prev byte = 0
		var curr byte = 0
		var first bool = false
		for {
			// Pull data off input line
			curr = <-input_chan

			if first {
				event = false
				first = false
			} else {

				// Determine if rising edge event is triggered
				event = (curr == 1) && (prev == 0)
			}

			// Put data onto output channel
			output_chan <- event

			// Self explanatory
			prev = curr
		}
	}()

	return output_chan
}

func counter(clk_chan chan uint8, enable_chan chan uint8, reset_chan chan uint8, output_chan chan uint8) {
	var count uint8 = 0
	evt_chan := rising_edge(clk_chan)

	for {
		enable := <-enable_chan
		reset := <-reset_chan
		event := <-evt_chan

		if event {
			if 1 == enable {
				count = count + 1
			} else if 1 == reset {
				count = 0
			}
		}
		output_chan <- count
	}
}
