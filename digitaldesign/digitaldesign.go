package digitaldesign

type DFF struct {
	input_port  chan uint8
	output_port chan uint8
	clk_port    chan uint8
}

func (d DFF) Start() {
	go func() {

		var state uint8 = 0

		for evt := range rising_edge(d.clk_port) {

			// Get data off input lines
			data := <-d.input_port

			// Event based logic
			if evt {
				state = data
			}

			// Put data onto output line
			d.output_port <- state
		}
	}()

}

func rising_edge(input_chan chan uint8) chan bool {

	output_chan := make(chan bool)

	go func() {
		var event bool = false
		var prev uint8 = 0
		var curr uint8 = 0
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

func d_flip_flop(clk_chan chan uint8, data_chan chan uint8, output_chan chan uint8) {
	var state uint8 = 0

	evt_chan := rising_edge(clk_chan)

	for {
		// Get data off input lines
		data := <-data_chan
		evt := <-evt_chan

		// Event based logic
		if evt {
			state = data
		}

		// Put data onto output line
		output_chan <- state

	}
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
