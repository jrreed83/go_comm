package digitaldesign

type dff struct {
	clkLine    <-chan byte
	dataLine   <-chan byte
	outputLine chan<- byte
}

func (d *dff) Start(state byte) {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			clk := <-d.clkLine
			data := <-d.dataLine
			if risingEdge(clk) {
				state = data
			}
			d.outputLine <- state
		}
	}()

}

func risingEdgeAction() func(byte) bool {
	isFirstSample := true
	var prevClk byte
	action := func(currClk byte) bool {
		if isFirstSample {
			isFirstSample = false
			prevClk = currClk
			return false
		}
		risingEdge := (currClk == 1) && (prevClk == 0)
		prevClk = currClk
		return risingEdge
	}
	return action
}

func risingEdgeGen(inputLine <-chan byte) chan bool {
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
