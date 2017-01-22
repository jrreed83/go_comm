package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type MsgType byte

type Message struct {
	Time uint32
	Data byte
	Type MsgType
}

type Signal struct {
	Time uint32
	Val  byte
	Msg  chan Message    // Informs reader processes what time-stamp to read
	Tbl  map[uint32]byte // Shared memory
	Num  int             // Number of readers/subscribers
}

func NewSignal(n int) *Signal {
	return &Signal{
		Time: 0,
		Val:  0,
		Msg:  make(chan Message),
		Tbl:  make(map[uint32]byte, n),
		Num:  n,
	}
}

func (s *Signal) WaitThenAssign(wait uint32, d byte) {
	t := s.Time + wait
	s.Tbl[t] = d
	s.Time = t

	if s.Val != d {
		s.Msg <- Message{Time: t, Data: d}
	}

	s.Time = t
	s.Val = d

}

func (s *Signal) Get() byte {
	return 0
}

func (s *Signal) Wait() (Message, error) {
	select {
	case msg := <-s.Msg:
		return msg, nil
	case <-time.After(time.Second):
		return Message{}, errors.New("No more signal")
	}
}

func main() {
	clk := NewSignal(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		clk.WaitThenAssign(1, 1)
		clk.WaitThenAssign(1, 0)
		clk.WaitThenAssign(1, 1)
		wg.Done()
	}()

	go func() {
		for {
			msg, err := clk.Wait()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				return
			}
			fmt.Println(msg)
		}

	}()

	wg.Wait()
}
