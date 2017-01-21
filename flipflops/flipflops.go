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
	Msg chan Message    // Informs reader processes what time-stamp to read
	Tbl map[uint32]byte // Shared memory
	Num int             // Number of readers/subscribers

	CurrTime uint32
	CurrVal  byte
	NxtTime  uint32 // Most Recent Time
	NxtVal   byte   // Most recent value
}

func NewSignal(n int) *Signal {
	return &Signal{
		Msg: make(chan Message),
		Tbl: make(map[uint32]byte, n),
		Num: n,
	}
}

func (s *Signal) Put(t uint32, d byte) {
	if t < s.NxtTime {
		panic("Must be a later time")
	}
	s.NxtTime = t
	s.NxtVal = d

	s.Tbl[t] = d
}

func (s *Signal) Get() byte {
	return 0
}

// Send most recent data to subscribers
func (s *Signal) SyncMsg() {
	var i int
	for i = 0; i < s.Num; i++ {
		s.Msg <- Message{Time: s.NxtTime, Data: s.NxtVal}
	}
}

func (s *Signal) Wait() (Message, error) {
	select {
	case msg := <-s.Msg:
		return msg, nil
	case <-time.After(time.Second):
		return Message{}, errors.New("No more signal")
	}
}

func (s *Signal) Raise() error {
	select {
	case s.Msg <- Message{}:
		return nil
	case <-time.After(time.Second):
		return errors.New("Error")
	}
}

type Clock struct {
	Time   uint32
	Event  *Signal
	Driver *Signal
}

func (c *Clock) Start() {
	go func() {
		for {
			_, err := c.Event.Wait()
			if err != nil {
				continue
			}
			pulse := uint8(c.Time % 2)
			c.Driver.Put(c.Time, pulse)

			// Synchronize
			c.Driver.SyncMsg()
			c.Time++

		}
	}()
}

func (c *Clock) Tick() {
	c.Event.Raise()
}

func main() {
	drv := NewSignal(1)
	evt := NewSignal(1)

	clk := Clock{Time: 0, Event: evt, Driver: drv}
	var wg sync.WaitGroup

	wg.Add(2)

	clk.Start()

	go func() {
		clk.Tick()
		clk.Tick()
		clk.Tick()
		wg.Done()
	}()

	go func() {
		for {
			msg, err := clk.Driver.Wait()
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
