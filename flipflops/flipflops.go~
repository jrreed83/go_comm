package flipflops

import (
//"time"
)

type Wire chan byte

func NewWire() Wire {
	return make(chan byte, 1)
}

func send(w Wire, x byte) {
	w <- x
}

func receive(w Wire) byte {
	return <-w
}

type DFlipFlop struct {
	reqLine    Wire
	enableLine Wire
	clkLine    Wire
	inputLine  Wire
	outputLine Wire
	state      byte
}

func NewDFlipFlop() *DFlipFlop {
	return &DFlipFlop{
		reqLine:    NewWire(),
		enableLine: NewWire(),
		clkLine:    NewWire(),
		inputLine:  NewWire(),
		outputLine: NewWire(),
		state:      0}

}

func (d *DFlipFlop) Start() {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			clk := receive(d.clkLine)
			if risingEdge(clk) {
				enable := receive(d.enableLine)
				if enable == 0 {
					d.state = receive(d.inputLine)
				}
			}
			send(d.outputLine, d.state)
		}
	}()

}

func (d *DFlipFlop) Write(x byte) {
	send(d.clkLine, 0)
	send(d.clkLine, 1)
	send(d.enableLine, 0)
	send(d.inputLine, x)
	return y

}

func (d *DFlipFlop) Read() byte {
	send(d.reqLine, 0)
	y := receive(d.outputLine)
	return y
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
