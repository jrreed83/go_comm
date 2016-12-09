package digitalcircuits

func rising_edge(input_chan chan uint8) chan bool {
	output_chan := make(chan bool)

	go func() {
		var state uint8

		// Pull the first sample off the channel rather
		state = <-input_chan
		for {
			curr := <-input_chan
			if curr == 1 && state == 0 {
				output_chan <- true // Event only gets triggered here
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

		curr_data := <-data_chan
		select {
		case <-evt_chan:
			state = curr_data
		default:
			continue
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

		select {
		case <-evt_chan:
			if 1 == enable {
				count = count + 1
			} else if 1 == reset {
				count = 0
			}
		default:
			continue
		}
		output_chan <- count
	}
}
