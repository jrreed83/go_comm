package hardware

func counter(clk chan uint8, reset chan uint8, enable chan uint8, output chan uint8) {
	var cnt uint = 0
	for {
		c := <-clk
		e := <-enable
		r := <-reset
		if 1 == c {
			if 1 == e {
				cnt = cnt + 1
			} else if 1 == r {
				cnt = 0
			}
		}
		output <- cnt
	}
}
