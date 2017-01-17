package flipflops

import (
//"time"
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
			c.State = (c.State + 1) % 2
			c.Time++
			for i = 0; i < c.NumReaders; i++ {
				c.Out <- Message{Time: c.Time, Data: c.State}
			}
		}
	}()
}

func (c *SystemClock) Tick() {
	select {
	case c.Event <- struct{}{}:
	default:
	}
}
