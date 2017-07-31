package main

import "fmt"
import "sync"
import "time"

type FlipFlop struct {
	Clk    chan byte
	Enable chan byte
	Input  chan byte
	Output chan byte
}

func NewFlipFlop() *FlipFlop {
	return &FlipFlop{
		Clk:    make(chan byte),
		Enable: make(chan byte),
		Input:  make(chan byte),
		Output: make(chan byte),
	}
}

func (F *FlipFlop) Start(prvClk, storage uint8) {
	go func() {
		for {
			// Collect input signals
			currClk := <-F.Clk
			enable := <-F.Enable
			input := <-F.Input

			// Set output signal
			F.Output <- storage

			// Modify storage register based on input signals
			if prvClk == 0 && currClk == 1 && enable == 1 {
				storage = input
			}

			prvClk = currClk
		}
	}()
}

func main() {
	f := NewFlipFlop()
	f.Start(0, 0)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		f.Clk <- 1
		f.Enable <- 1
		f.Input <- 1

		f.Clk <- 0
		f.Enable <- 1
		f.Input <- 1

		f.Clk <- 1
		f.Enable <- 1
		f.Input <- 1

		wg.Done()
	}()

	go func() {
		for {
			select {
			case o := <-f.Output:
				fmt.Println(o)
			case <-time.After(5 * time.Second):
				fmt.Println("Quitting...")
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}
