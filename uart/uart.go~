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

func (u *Uart) Put(x byte) error {
	// Put byte onto transmit buffer
	select {
	case u.txBuffer <- x:
		return nil
	case <-time.After(1):
		return errors.New("Put method failed due to timeout")
	}
}

func (u *Uart) Get() (byte, error) {
	// Take byte off receive buffer
	select {
	case x := <-u.rxBuffer:
		return x, nil
	case <-time.After(1):
		return 0, errors.New("Get method failed due to timeout")
	}
}

func main() {
	u := NewUart()
	u.Start()
	var err error
	var x byte
	var i int

	for i = 0; i < 10; i++ {
		if err = u.Put(byte(i)); err != nil {
			fmt.Println(err)
			return
		}
	}

	for i = 0; i < 10; i++ {
		if x, err = u.Get(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(x)
	}

}
