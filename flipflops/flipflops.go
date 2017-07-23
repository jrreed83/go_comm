package main

import (
	//"errors"
	"fmt"
	//"sync"
	"time"
)

type Request struct {
	T      uint32
	X      byte
	Sender chan byte
}

type SignalServer struct {
	wave       map[uint32]byte
	readQueue  chan Request
	writeQueue chan Request
}

type SignalListener struct {
	readQueue chan Request
	Response  chan byte
}

type SignalDriver struct {
	writeQueue chan Request
	Response   chan byte
}

func (s *SignalServer) Start() {
	go func() {
		for {
			select {
			case request := <-s.writeQueue:
				x := request.X
				t := request.T
				s.wave[t] = x
				request.Sender <- 0
			case request := <-s.readQueue:
				t := request.T
				x := s.wave[t]
				request.Sender <- x
			}

		}

	}()
}

// Probably need to protect the read pointer.  Although
// we probably only really will read it when we get an event
func StartDriver(sig *Signal) {
	go func() {
		var clk uint32 = 0

		for {
			select {
			case msg := <-sig.PutReq:
				x := msg.X
				t := msg.T
				prv := sig.Wave[clk]
				newClk := clk + t + 1
				sig.Wave[newClk] = x
				clk = newClk
				if prv != msg.X {
					sig.Evt <- clk
				}
			case <-sig.GetReq:
				x := sig.Wave[clk]
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
		Wave:    make(map[uint32]byte),
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
	Wave    map[uint32]byte // Waveform.  Not directly accessible by processes.
}

func (s *Signal) Assign(x byte) {
	select {
	case s.PutReq <- Msg{X: x}:
	case <-time.After(time.Second):
		panic("Assign timed-out")
	}
}

// Get ...
func (s *Signal) Get() Msg {
	select {
	case s.GetReq <- 0:
		msg := <-s.GetResp
		return msg
	case <-time.After(time.Second):
		panic("Get timed-out")
	}
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

	s.Assign(1)
	s.Assign(2)
	x := s.Get()
	fmt.Println(x)

	fmt.Println("Hello")
}
