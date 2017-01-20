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

func (s *Signal) Notify(t uint32) {
	var i int
	for i = 0; i < s.Num; i++ {
		s.Msg <- Message{Time: t, Data: 0}
	}
}

func (s *Signal) Wait() Message {
	return <-s.Msg
}

type Clock struct {
	Time uint32
	In   *Signal
	Out  *Signal
}

func (c *Clock) Start() {
	go func() {
		for {
			c.In.Wait()
			pulse := uint8(c.Time % 2)
			c.Out.Put(c.Time, pulse)
			c.Out.Notify(c.Time)
			c.Time++

		}
	}()
}

func (c *Clock) Tick() {
	c.In.Msg <- Message{}
}

func main() {
	s := NewSignal(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		s.Put(0, 0)
		s.Put(1, 4)
		s.Notify(1)
		wg.Done()
	}()

	go func() {
		fmt.Println(s.Wait())
		wg.Done()
	}()

	wg.Wait()
}
