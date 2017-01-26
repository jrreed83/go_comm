package main

import (
	//"errors"
	"fmt"
	//"sync"
	//"time"
)

func StartDriver(sig *Signal) {
	go func() {
		var tr uint32 = 0
		var tw uint32 = 0

		for {
			select {
			case msg := <-sig.PutReqChan:

				prv := sig.Tbl[tr]

				sig.Tbl[tw] = msg.X
				tr = tw
				tw++

				if prv != msg.X {
					sig.EvtChan <- tr
				}

			case <-sig.GetReqChan:
				x := sig.Tbl[tr]
				sig.GetChan <- Msg{X: x, T: tr}
			}

		}

	}()
}

func NewSignal() *Signal {
	return &Signal{
		GetReqChan: make(chan uint32),
		PutReqChan: make(chan Msg),
		GetChan:    make(chan Msg),
		EvtChan:    make(chan uint32),
		Tbl:        make(map[uint32]byte),
	}
}

type Msg struct {
	T uint32
	X byte
}

type Signal struct {
	GetReqChan chan uint32
	PutReqChan chan Msg
	GetChan    chan Msg
	EvtChan    chan uint32
	Tbl        map[uint32]byte // Shared memory
	Num        int             // Number of readers/subscribers
}

func (s *Signal) Assign(x byte) {
	s.PutReqChan <- Msg{X: x}
}

func (s *Signal) Get() Msg {
	s.GetReqChan <- 0
	msg := <-s.GetChan
	return msg
}

func main() {

	s := NewSignal()

	StartDriver(s)

	go func() {
		for {
			t := <-s.EvtChan
			fmt.Println(t)
		}
	}()

	s.Assign(1)
	s.Assign(1)
	s.Assign(1)
	s.Assign(2)
	x := s.Get()
	fmt.Println(x)

	fmt.Println("Hello")
}
