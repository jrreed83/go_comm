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
	Clk  chan uint32 // Informs reader processes what time-stamp to read
	Evt  chan struct{}
	Tbl  map[uint32]byte // Shared memory
	Num  int             // Number of readers/subscribers
}

func (s *Signal) WaitThenAssign(wait int, d byte) {

	var t uint32
	for i := 0; i < wait; i++ {
		t = <-s.Clk
	}
	s.Val = d
	s.Tbl[t] = d
	s.Time = t
	s.Evt <- struct{}{}
}

func (s *Signal) Get() byte {
	return s.Val
}

func (s *Signal) Wait() error {
	select {
	case <-s.Evt:
		return nil
	case <-time.After(time.Second):
		return errors.New("No more signal")
	}
}

func main() {
	c := make(chan uint32)
	d := NewSignal(c)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		c <- 0
		c <- 1
		c <- 2
		c <- 3
		c <- 4
	}()

	go func() {
		clk.WaitThenAssign(1, 64)
		clk.WaitThenAssign(1, 36)
		clk.WaitThenAssign(1, 79)
		wg.Done()
	}()

	go func() {
		for {
			err := clk.Wait()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				return
			}
			fmt.Println(clk.Get())
		}

	}()

	wg.Wait()
}
