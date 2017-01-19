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
	Event      chan struct{}
	Out        chan Message
	NumReaders uint8
}

func NewSystemClock(numReaders uint8) *SystemClock {
	return &SystemClock{
		Time:       0,
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
				c.Out <- Message{Time: c.Time}
			}
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

func (c *SystemClock) wait(pause uint8, data byte) Message {

	var cnt uint8
	var msg Message

	cnt = 0
	for {
		select {
		case msg = <-c.Out:
			if cnt == pause {
				return Message{Time: msg.Time, Data: data}
			}
			cnt++
		case <-time.After(time.Second):
			panic("wait method timed out")
		}

	}

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
		d <- sysClk.wait(2, 1)
		d <- sysClk.wait(2, 4)
		wg.Done()
	}()

	go func() {
		for {
			select {
			case msg := <-d:
				fmt.Printf("%d %d\n", msg.Time, msg.Data)
			case <-time.After(time.Millisecond):
				wg.Done()
				return

			}
		}
	}()

	wg.Wait()

}
