package digitalcircuits

func rising_edge(input chan uint8) chan bool {
	out := make(chan bool)

	go func() {
		var state uint8 = 0
		var status bool
		for {
			curr := <-input
			if curr == 1 && state == 0 {
				status = true
			} else {
				status = false
			}
			out <- status
			state = curr
		}
	}()

	return out
}

func d_flip_flop(clk chan uint8, data chan uint8, output chan uint8) {
	var state uint8 = 0
	var prv_clk uint8 = 0

	for {
		curr_clk := <-clk
		curr_data := <-data

		if 1 == curr_clk && prv_clk == 0 {
			output <- state
			state = curr_data
		}
	}
}

func counter(clk chan uint8, enable chan uint8, reset chan uint8, output chan uint8) {
	var cnt uint8 = 0
	var prv_clk uint8 = 0
	for {
		curr_clk := <-clk
		curr_en := <-enable
		curr_rst := <-reset
		if 1 == curr_clk && prv_clk == 0 {
			if 1 == curr_en {
				cnt = cnt + 1
			} else if 1 == curr_rst {
				cnt = 0
			}
		}
		prv_clk = curr_clk
		output <- cnt
	}
}
