package main

import (
	//"errors"
	"fmt"
	//"sync"
	"time"
)

// Probably need to protect the read pointer.  Although
// we probably only really will read it when we get an event
func StartDriver(sig *Signal) {
	go func() {
		var clk uint32 = 0

		for {
			select {
			case msg := <-sig.PutReq:
		
				x := msg.X 
				t := msg.T + 1

				prv := sig.Tbl[clk]
 
				sig.Tbl[clk + t] = x

				clk += t

				if prv != msg.X {
					sig.Evt <- clk 
				}

			case <-sig.GetReq:
				x := sig.Tbl[clk]
				sig.GetResp <- Msg{X: x, T: clk}
			}

		}

	}()
}

func NewSignal() *Signal {
	return &Signal{
		GetReq:  make(chan uint32),
		PutReq:  make(chan Msg),
		GetResp: make(chan Msg),
		Evt:     make(chan uint32),
		Tbl:     make(map[uint32]byte),
	}
}

type Msg struct {
	T uint32
	X byte
}

type Signal struct {
	GetReq  chan uint32
	PutReq  chan Msg
	GetResp chan Msg
	Evt     chan uint32
	Tbl     map[uint32]byte // Shared memory
	Num     int             // Number of readers/subscribers
}

func (s *Signal) Assign(x byte) {
	select{
	case s.PutReq <- Msg{X: x}:
	case <-time.After(time.Second):
		panic("Assign timed-out")
	}
}

// Here 't' denotes wait time
func (s *Signal) WaitThenAssign(t uint32, x byte) {
	s.PutReq <- Msg{X: x, T: t}
}

func (s *Signal) Get() Msg {
	s.GetReq <- 0
	msg := <-s.GetResp
	return msg
}

func main() {

	s := NewSignal()

	StartDriver(s)

	go func() {
		for {
			t := <-s.Evt
			fmt.Println(t)
		}
	}()

	s.WaitThenAssign(1,1)
	s.WaitThenAssign(2,1)
	s.Assign(1)
	s.Assign(2)
	x := s.Get()
	fmt.Println(x)

	fmt.Println("Hello")
}
