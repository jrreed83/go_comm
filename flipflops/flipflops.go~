package flipflops

import (
	"fmt"
	"github.com/jrreed83/go_comm/port"
	"time"
)

type DFF struct {
	input  port.Port
	output port.Port
	clk    port.Port
	enable chan byte
}

func NewDFF() *DFF {
	d := DFF{
		input:  *port.NewPort(),
		output: *port.NewPort(),
		clk:    *port.NewPort(),
		enable: make(chan byte)}

	return &d
}

func (d *DFF) Start() {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			<-d.enable
			clk := d.clk.Read()
			if risingEdge(clk) {
				fmt.Println("Rising Edge")
				input := d.input.Read()
				d.output.Write(input)
			}
		}
	}()
}

type DFlipFlop struct {
	reqLine    chan struct{}
	enableLine chan struct{}
	clkLine    chan byte
	inputLine  chan byte
	outputLine chan byte
	state      byte
}

func NewDFlipFlop() *DFlipFlop {
	d := DFlipFlop{
		reqLine:    make(chan struct{}),
		enableLine: make(chan struct{}),
		clkLine:    make(chan byte),
		inputLine:  make(chan byte),
		outputLine: make(chan byte),
		state:      0}

	return &d
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
				case <-time.After(time.Millisecond):
				}

				select {
				case <-d.reqLine:
					d.outputLine <- d.state
				case <-time.After(time.Millisecond):
				}

			}

		}
	}()

}

func (d *DFlipFlop) Write(x byte) {
	d.clkLine <- 0
	d.clkLine <- 1
	d.enableLine <- struct{}{}
	d.inputLine <- x
}

func (d *DFlipFlop) Read() byte {
	d.clkLine <- 0
	d.clkLine <- 1
	d.reqLine <- struct{}{}
	x := <-d.outputLine
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
