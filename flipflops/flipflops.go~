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
}

type Signal struct {
	Val chan Message
	Clk chan uint32 // Informs reader processes what time-stamp to read
	Evt chan struct{}
	Tbl map[uint32]byte // Shared memory
	Num int             // Number of readers/subscribers
}

func (s *Signal) WaitThenAssign(wait int, d byte) {

	var t uint32
	for i := 0; i < wait; i++ {
		select {
		case t = <-s.Clk:
		case <-time.After(time.Second):
			panic("Time Out")
		}
	}

	s.Val <- Message{Time: t, Data: d}
	s.Evt <- struct{}{}
}

func (s *Signal) Get() Message {
	return <-s.Val
}

func (s *Signal) WaitForEvent() error {
	select {
	case <-s.Evt:
		return nil
	case <-time.After(time.Second):
		return errors.New("No more signal")
	}
}

func main() {
	c := make(chan uint32)
	d := &Signal{
		Clk: c,
		Evt: make(chan struct{}, 1),
		Val: make(chan Message, 1),
	}

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
		d.WaitThenAssign(1, 64)
		d.WaitThenAssign(1, 36)
		d.WaitThenAssign(1, 79)
		wg.Done()
	}()

	go func() {
		for {
			err := d.WaitForEvent()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				return
			}
			fmt.Println(d.Get())
		}

	}()

	wg.Wait()
}
