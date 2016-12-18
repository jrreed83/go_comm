package digitaldesign

type dff struct {
	dataLine   chan byte
	outputLine chan byte
	clkLine    chan byte
}

func (d *dff) Start(state byte) {
	go func() {
		//var state byte = 0
		evtLine := risingEdge(d.clkLine)
		for {
			evt := <-evtLine
			data := <-d.dataLine
			if evt {
				state = data
			}
			d.outputLine <- state
		}
	}()

}

func risingEdge(inputLine chan byte) chan bool {
	outputLine := make(chan bool)
	go func() {
		isRisingEdge := false
		isFirstSample := false
		var prev byte = 0
		for {
			curr := <-inputLine
			if isFirstSample {
				isRisingEdge = false
				isFirstSample = false
			} else {
				isRisingEdge = (curr == 1) && (prev == 0)
			}
			outputLine <- isRisingEdge
			prev = curr
		}
	}()
	return outputLine
}
