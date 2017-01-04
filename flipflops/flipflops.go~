package flipflops

import (
//"time"
)

type DFlipFlop struct {
	reqLine    chan byte
	enableLine chan byte
	clkLine    chan byte
	inputLine  chan byte
	outputLine chan byte
	state      byte
}

func NewDFlipFlop() *DFlipFlop {
	return &DFlipFlop{
		reqLine:    make(chan byte),
		enableLine: make(chan byte),
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

				select {
				case <-d.enableLine:
					d.state = <-d.inputLine
					<-d.enableLine
				case <-d.reqLine:
					d.outputLine <- d.state
					<-d.reqLine
				}
			}
		}
	}()

}

func (d *DFlipFlop) Write(x byte) {
	d.clkLine <- 0
	d.clkLine <- 1
	d.enableLine <- 0
	d.inputLine <- x
	d.enableLine <- 1

}

func (d *DFlipFlop) Read() byte {
	d.clkLine <- 0
	d.clkLine <- 1
	d.reqLine <- 0
	x := <-d.outputLine
	d.reqLine <- 1
	return x
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
