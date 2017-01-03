package flipflops

import (
	"sync"
	"testing"
	"time"
)

func TestDFF(t *testing.T) {

	var x byte
	d := NewDFF()

	t.Log("Kicking off Flip-Flop")
	d.Start()

	d.clk.Write(0)
	d.data.Write(0)
	x = d.Sample()
	t.Log(x)

	d.clk.Write(1)
	d.data.Write(1)
	x = d.Sample()
	t.Log(x)

	d.clk.Write(0)
	d.data.Write(1)
	x = d.Sample()
	t.Log(x)

	d.clk.Write(1)
	d.data.Write(0)
	x = d.Sample()
	t.Log(x)

}

func TestDFlipFlop(t *testing.T) {

	d := NewDFlipFlop()

	t.Log("Kicking off Flip-Flop")
	d.Start()

	var wg sync.WaitGroup

	wg.Add(3)

	t.Log("Kicking Off Clock")
	go func() {
		d.clkLine <- 0
		d.clkLine <- 1
		d.clkLine <- 0
		d.clkLine <- 1
		d.clkLine <- 0
		d.clkLine <- 1
		d.clkLine <- 0
		d.clkLine <- 1
		d.clkLine <- 0
		d.clkLine <- 1
		d.clkLine <- 0
		wg.Done()
	}()
	t.Log("Kicking Off Data Line")
	go func() {
		d.dataLine <- 1
		d.dataLine <- 0
		d.dataLine <- 0
		d.dataLine <- 1
		d.dataLine <- 1
		d.dataLine <- 1
		d.dataLine <- 0
		d.dataLine <- 1
		d.dataLine <- 1
		d.dataLine <- 1
		d.dataLine <- 0
		wg.Done()
	}()

	go func() {
		for {
			select {
			case y := <-d.outputLine:
				t.Log(y)
			case <-time.After(1):
				t.Log("Times Out ... Quitting")
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()

}
