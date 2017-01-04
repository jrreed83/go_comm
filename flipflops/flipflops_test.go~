package flipflops

import (
	"testing"
	//	"time"
)

func TestDFlipFlop(t *testing.T) {
	var x byte

	d := NewDFlipFlop()

	t.Log("Kicking off Flip-Flop")
	d.Start()

	d.Write(3)
	x = d.Read()
	t.Log(x)

	d.Write(10)
	x = d.Read()
	t.Log(x)

}
