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
		var event bool = false
		var prev byte = 0
		var curr byte = 0
		var first bool = false
		for {
			curr = <-inputLine
			if first {
				event = false
				first = false
			} else {
				event = (curr == 1) && (prev == 0)
			}
			outputLine <- event
			prev = curr
		}
	}()

	return outputLine
}
