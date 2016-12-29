package main

import "fmt"
import "time"
import "errors"

type Uart struct {
	txBuffer chan byte
	rxBuffer chan byte
	txChan   chan byte
	rxChan   chan byte
}

type Bus struct {
	txChan1 chan byte
	rxChan1 chan byte
	txChan2 chan byte
	rxChan2 chan byte
}

func (b *Bus) Start() {
	go func() {
		for {
			b.rxChan2 <- <-b.txChan1
		}
	}()
}

func NewUart() *Uart {
	return &Uart{
		txBuffer: make(chan byte, 8),
		rxBuffer: make(chan byte, 8),
		txChan:   make(chan byte),
		rxChan:   make(chan byte)}
}

func (u *Uart) Start() {
	go func() {
		var i byte
		for {
			x := <-u.txBuffer
			for i = 0; i < 8; i++ {
				bit := (x >> i) & 0x01
				u.txChan <- bit
			}
		}
	}()

	go func() {
		var x byte
		var i byte
		for {
			x = 0
			for i = 0; i < 8; i++ {
				bit := <-u.rxChan
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
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("Put method failed due to timeout")
	}
}

func (u *Uart) Get() (byte, error) {
	// Put byte onto transmit buffer
	var x byte
	select {
	case x = <-u.rxBuffer:
		return x, nil
	case <-time.After(1 * time.Second):
		return x, errors.New("Get method failed due to timeout")
	}
}

func main() {

	u1 := NewUart()
	u2 := NewUart()
	bus := Bus{
		txChan1: u1.txChan,
		rxChan1: u1.rxChan,
		txChan2: u2.txChan,
		rxChan2: u2.rxChan}

	u1.Start()
	u2.Start()
	bus.Start()

	var err error
	var x byte

	if err = u1.Put(86); err != nil {
		fmt.Println(err)
		return
	}

	if err = u1.Put(76); err != nil {
		fmt.Println(err)
		return
	}

	if x, err = u2.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\n", x)

	if x, err = u2.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\n", x)
}
