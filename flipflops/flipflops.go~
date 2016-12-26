package flipflops

type DFlipFlop struct {
	clkLine    <-chan byte
	dataLine   <-chan byte
	outputLine chan<- byte
}

func (d *DFlipFlop) Start(state byte) {
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

type SpiMaster struct {
	cmdLine  chan<- byte
	sclkLine chan<- byte
	mosiLine chan<- byte
	misoLine <-chan byte
	ssLine   <-chan byte
}

type SpiSlave struct {
	sclkLine   <-chan byte
	mosiLine   <-chan byte
	misoLine   chan<- byte
	ssLine     chan<- byte
	outputLine chan<- byte
}

type SpiBus struct {
	master SpiMaster
	slave  SpiSlave
}
