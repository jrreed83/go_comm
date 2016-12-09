package digitalcircuits

func rising_edge(input_chan chan uint8) chan bool {
	output_chan := make(chan bool)

	go func() {
		var state uint8

		// Pull the first sample off the channel rather
		state = <-input_chan
		output_chan <- false
		for {
			curr := <-input_chan
			if curr == 1 && state == 0 {
				output_chan <- true
			} else {
				output_chan <- false
			}
			state = curr
		}
	}()

	return output_chan
}

func d_flip_flop(clk_chan chan uint8, data_chan chan uint8, output_chan chan uint8) {
	var state uint8 = 0

	evt_chan := rising_edge(clk_chan)

	for {

		data := <-data_chan
		evt := <-evt_chan
		if evt {
			state = data
		}
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
