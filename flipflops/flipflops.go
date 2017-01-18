package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Time uint32
	Data byte
}

type SystemClock struct {
	Time       uint32
	State      byte
	Event      chan struct{}
	Out        chan Message
	NumReaders uint8
}

func NewSystemClock(numReaders uint8) *SystemClock {
	return &SystemClock{
		Time:       0,
		State:      0,
		Event:      make(chan struct{}),
		Out:        make(chan Message, numReaders),
		NumReaders: numReaders,
	}
}

func (c *SystemClock) Start() {
	go func() {
		var i int
		for {
			<-c.Event
			for i = 0; i < int(c.NumReaders); i++ {
				c.Out <- Message{Time: c.Time, Data: c.State}
			}
			c.State = (c.State + 1) % 2
			c.Time++

		}
	}()
}

func (c *SystemClock) Tick() {
	for {
		select {
		case c.Event <- struct{}{}:
			return
		default:
			continue
		}
	}
}

func wait(clk chan Message, n uint8, x byte) Message {
	var i int
	var msg Message
	for i = 0; i <= int(n); i++ {
		msg = <-clk
	}

	return Message{Time: msg.Time, Data: x}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	sysClk := NewSystemClock(1)
	sysClk.Start()

	d := make(chan Message)

	go func() {

		sysClk.Tick()
		sysClk.Tick()
		sysClk.Tick()
		sysClk.Tick()
		sysClk.Tick()
		sysClk.Tick()

		wg.Done()
	}()

	go func() {
		d <- wait(sysClk.Out, 2, 1)
		d <- wait(sysClk.Out, 2, 4)
		wg.Done()
	}()

	go func() {
		for {
			select {
			case msg := <-d:
				fmt.Printf("%d %d\n", msg.Time, msg.Data)
			case <-time.After(time.Second):
				wg.Done()
				return

			}
		}
	}()

	wg.Wait()

}
