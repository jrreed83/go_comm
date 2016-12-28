package main

import "fmt"
import . "github.com/jrreed83/go_comm/ringbuffer"

type Uart struct {
	txBuffer RingBuffer
	rxBuffer RingBuffer
	channel  chan byte
}

func newUart() *Uart {
	return &Uart{cmdChan: make(chan byte),
		inputChan:  make(chan byte),
		outputChan: make(chan byte)}
}
func (u *Uart) Start() {
	// Take data off transmit buffer and send across channel
	go func() {
		for {
			x := u.txBuffer.Read()
			u.channel <- x
		}
	}()

	// Take data off channel and put onto receive buffer
	go func() {
		for {
			x := <-u.channel
			u.rxBuffer.Write(x)
		}
	}()
}

func (u *Uart) Put(b byte) {
	// Put byte onto transmit buffer
	u.txBuffer.Write(b)
}

func (u *Uart) Take() byte {
	// Take byte off receive buffer
	return u.rxBuffer.Read()
}

func main() {
	u := newUart()

	u.Start()

	u.Put(1)
	x := u.Take()

	fmt.Println(x)
}
