package main

import (
	"fmt"
	"sync"
	//	"time"
)

type Message struct {
	Time uint32
	Data byte
}

type Signal struct {
	Msg chan Message    // Informs reader processes what time-stamp to read
	Tbl map[uint32]byte // Shared memory
	Num int             // Number of readers/subscribers
}

func NewSignal(n int) *Signal {
	return &Signal{
		Msg: make(chan Message),
		Tbl: make(map[uint32]byte, n),
		Num: n,
	}
}

func (s *Signal) Put(t uint32, d byte) {
	s.Tbl[t] = d
}

func (s *Signal) Get() byte {
	return 0
}

func (s *Signal) SyncMsg(t uint32) {
	var i int
	for i = 0; i < s.Num; i++ {
		s.Msg <- Message{Time: t}
	}
}

func (s *Signal) Wait() Message {
	return <-s.Msg
}

func (s *Signal) Raise() {
	s.Msg <- Message{}
}

type Clock struct {
	Time   uint32
	Event  *Signal
	Driver *Signal
}

func (c *Clock) Start() {
	go func() {
		for {
			c.Event.Wait()
			pulse := uint8(c.Time % 2)
			c.Driver.Put(c.Time, pulse)
			c.Driver.SyncMsg(c.Time)
			c.Time++

		}
	}()
}

func (c *Clock) Tick() {
	c.Event.Raise()
}

func main() {
	s := NewSignal(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		s.Put(0, 0)
		s.Put(1, 4)
		s.SyncMsg(1)
		wg.Done()
	}()

	go func() {
		fmt.Println(s.Wait())
		wg.Done()
	}()

	wg.Wait()
}
