package main

import (
	//"errors"
	"fmt"
	"sync"
	//"time"
)

type Signal struct {
	Clk <-chan uint32 // Informs reader processes what time-stamp to read
	Evt chan uint32
	Tbl map[uint32]byte // Shared memory
	Num int             // Number of readers/subscribers
}

func (s *Signal) Get(t uint32) byte {
	return s.Tbl[t]
}

func (s *Signal) Put(t uint32, x byte) {
	s.Tbl[t] = x
}

type Process struct {
	Time uint32
}

func (p *Process) Wait(s *Signal) {
	select {
	case t := <-s.Evt:
		p.Time = t
		return
	}

}

func (p *Process) Get(s *Signal) byte {
	return s.Get(p.Time)
}

func (p *Process) Put(s *Signal, x byte) {
	s.Put(p.Time, x)
}

func (p *Process) Start(CLK *Signal, D *Signal, Q *Signal) {
	go func() {
		for {
			p.Wait(CLK)
			x := p.Get(D)
			p.Put(Q, x)
		}
	}()
}

func main() {

	CLK := &Signal{Evt: make(chan uint32), Tbl: make(map[uint32]byte)}
	D := &Signal{Evt: make(chan uint32), Tbl: make(map[uint32]byte)}
	Q := &Signal{Evt: make(chan uint32), Tbl: make(map[uint32]byte)}

	dff := &Process{}

	dff.Start(CLK, D, Q)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		CLK.Put(0, 0)
		CLK.Put(1, 1)
		CLK.Evt <- 1
		CLK.Put(2, 0)
		CLK.Put(3, 1)
		CLK.Evt <- 3
		CLK.Put(4, 0)
		wg.Done()
	}()

	go func() {
		D.Put(0, 32)
		D.Put(1, 42)
		D.Put(2, 98)
		D.Put(3, 76)
		D.Put(4, 34)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(Q.Tbl)
}
