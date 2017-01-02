package flipflops

type DFlipFlop struct {
	clkLine    chan byte
	dataLine   chan byte
	outputLine chan byte
	state      byte
}

func NewDFlipFlop() *DFlipFlop {
	return &DFlipFlop{
		clkLine:    make(chan byte, 10),
		dataLine:   make(chan byte, 10),
		outputLine: make(chan byte, 10),
		state:      0}
}
func (d *DFlipFlop) Start() {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			clk := <-d.clkLine
			data := <-d.dataLine
			if risingEdge(clk) {
				//				d.state = data
				d.outputLine <- data
			}
			//			d.outputLine <- d.state
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
