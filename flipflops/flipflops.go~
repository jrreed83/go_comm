package flipflops

import (
//"time"
)

type DFlipFlop struct {
	enableLine chan byte
	clkLine    chan byte
	inputLine  chan byte
	outputLine chan byte
	state      byte
}

func NewDFlipFlop() *DFlipFlop {
	return &DFlipFlop{
		enableLine: make(chan byte, 1),
		clkLine:    make(chan byte),
		inputLine:  make(chan byte),
		outputLine: make(chan byte),
		state:      0}

}

func (d *DFlipFlop) Start() {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			clk := <-d.clkLine
			if risingEdge(clk) {
				enable := <-d.enableLine
				if enable == 0 {
					d.state = <-d.inputLine
				}
			}
			d.outputLine <- d.state
		}
	}()

}

func (d *DFlipFlop) Write(x byte) byte {
	d.enableLine <- 0
	d.clkLine <- 0
	_ = <-d.outputLine
	d.clkLine <- 1
	d.inputLine <- x
	y := <-d.outputLine
	return y

}

func (d *DFlipFlop) Read() byte {
	d.enableLine <- 1
	d.clkLine <- 0
	_ = <-d.outputLine
	d.clkLine <- 1
	y := <-d.outputLine
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
