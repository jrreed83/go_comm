package digitaldesign

type dff struct {
	clkLine    <-chan byte
	dataLine   <-chan byte
	outputLine chan<- byte
}

func (d *dff) Start(state byte) {
	go func() {
		risingEdgeLine := risingEdge(d.clkLine)
		for {
			isRisingEdge := <-risingEdgeLine
			data := <-d.dataLine
			if isRisingEdge {
				state = data
			}
			d.outputLine <- state
		}
	}()

}

func risingEdge(inputLine <-chan byte) chan bool {
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
