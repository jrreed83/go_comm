package flipflops

import "github.com/jrreed83/go_comm/port"

type DFF struct {
	clk        port.Port
	data       port.Port
	output     port.Port
	reqChan    chan struct{}
	outputChan chan byte
}

func NewDFF() *DFF {
	return &DFF{
		clk:        *port.NewPort(),
		data:       *port.NewPort(),
		output:     *port.NewPort(),
		reqChan:    make(chan struct{}),
		outputChan: make(chan byte)}
}

func (d *DFF) Sample() byte {
	d.reqChan <- struct{}{}
	return <-d.outputChan
}

func (d *DFF) Start() {
	go func() {
		risingEdge := risingEdgeAction()
		for {
			select {
			case <-d.reqChan:
				clk := d.clk.Read()
				if risingEdge(clk) {
					data := d.data.Read()
					d.output.Write(data)
				}
				d.outputChan <- d.output.Read()
			}
		}
	}()
}

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
				d.state = data
				d.outputLine <- data
			}
			d.outputLine <- d.state
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
