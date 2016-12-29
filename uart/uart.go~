package main

import "fmt"
import "time"

//import . "github.com/jrreed83/go_comm/ringbuffer"

type Uart struct {
	txBuffer chan byte
	rxBuffer chan byte
	channel  chan byte
}

func newUart() *Uart {
	return &Uart{txBuffer: make(chan byte, 8),
		rxBuffer: make(chan byte, 8),
		channel:  make(chan byte)}
}
func (u *Uart) Start() {
	// Take data off transmit buffer and send across channel
	go func() {
		for {
			x := <-u.txBuffer
			u.channel <- x
		}
	}()

	// Take data off channel and put onto receive buffer
	go func() {
		for {
			x := <-u.channel
			u.rxBuffer <- x
		}
	}()
}

func (u *Uart) Put(x byte) {
	// Put byte onto transmit buffer
	select {
	case u.txBuffer <- x:
	case <-time.After(1):
	}
}

func (u *Uart) Get() byte {
	// Take byte off receive buffer
	select {
	case x := <-u.rxBuffer:
		return x
	case <-time.After(1):
		return 0
	}
}

func main() {
	u := newUart()
	u.Start()
	u.Put(1)
	u.Put(5)

	x := u.Get()
	fmt.Println(x)

	y := u.Get()
	fmt.Println(y)
}
