package main

import "fmt"
import "time"
import "errors"

type Uart struct {
	txBuffer  chan byte
	rxBuffer  chan byte
	txChannel chan byte
	rxChannel chan byte
}

func NewUart() *Uart {
	return &Uart{txBuffer: make(chan byte, 16),
		rxBuffer:  make(chan byte, 16),
		txChannel: make(chan byte),
		rxChannel: make(chan byte)}
}
func (u *Uart) Start() {
	// Transmitter
	go func() {
		var i byte
		for {
			x := <-u.txBuffer
			for i = 0; i < 8; i++ {
				bit := (x >> i) & 0x01
				u.txChannel <- bit
			}
		}
	}()

	// Receiver
	go func() {
		var x byte
		var i byte
		for {
			x = 0
			for i = 0; i < 8; i++ {
				bit := <-u.txChannel
				x |= (bit << i)
			}
			u.rxBuffer <- x
		}
	}()
}

func (u *Uart) Put(x byte) error {
	// Put byte onto transmit buffer
	select {
	case u.txBuffer <- x:
	case <-time.After(1):
		return errors.New("Put method failed due to timeout")
	}
	return nil
}

func (u *Uart) Get() (byte, error) {
	// Put byte onto transmit buffer
	var x byte
	select {
	case x = <-u.rxBuffer:
	case <-time.After(1):
		return x, errors.New("Put method failed due to timeout")
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
