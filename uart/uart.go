package main

import "fmt"
import "time"
import "errors"

type Uart struct {
	txBuffer chan byte
	rxBuffer chan byte
	channel  chan byte
}

func NewUart() *Uart {
	return &Uart{txBuffer: make(chan byte, 16),
		rxBuffer: make(chan byte, 16),
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

func (u *Uart) Put(x byte) error {
	// Put byte onto transmit buffer
	var i byte
	for i = 0; i < 8; i++ {
		bit := (x >> i) & 0x01
		select {
		case u.txBuffer <- bit:
		case <-time.After(1):
			return errors.New("Put method failed due to timeout")
		}
	}
	return nil
}

func (u *Uart) Get() (byte, error) {
	// Put byte onto transmit buffer
	var x byte
	var i byte
	for i = 0; i < 8; i++ {
		select {
		case bit := <-u.rxBuffer:
			x |= (bit << i)
		case <-time.After(1):
			return x, errors.New("Put method failed due to timeout")
		}
	}
	return x, nil
}

func main() {
	u := NewUart()
	u.Start()
	var err error
	var x byte

	if err = u.Put(86); err != nil {
		fmt.Println(err)
		return
	}

	if x, err = u.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%c\n", x)
}
