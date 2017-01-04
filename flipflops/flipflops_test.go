package flipflops

import (
	"testing"
	"time"
)

func TestDFF(t *testing.T) {

	var x byte

	//var x byte
	d := NewDFF()

	t.Log("Kicking off Flip-Flop")
	d.Start()

	d.input.Write(0)
	d.clk.Write(0)
	d.enable <- 0
	time.Sleep(10)
	x = d.output.Read()
	t.Log(x)

	d.input.Write(1)
	d.clk.Write(1)
	d.enable <- 0
	time.Sleep(10)
	x = d.output.Read()
	t.Log(x)

}

func TestDFlipFlop(t *testing.T) {
	var x byte

	d := NewDFlipFlop()

	t.Log("Kicking off Flip-Flop")
	d.Start()

	d.Write(3)
	x = d.Read()
	t.Log(x)

}
