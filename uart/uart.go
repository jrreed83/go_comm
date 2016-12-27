package main

import "fmt"

type Uart struct {
	cmdChan    chan byte
	inputChan  chan byte
	outputChan chan byte
}

func newUart() *Uart {
	return &Uart{cmdChan: make(chan byte),
		inputChan:  make(chan byte),
		outputChan: make(chan byte)}
}
func (u *Uart) Start() {
	// Start transmitter.  Forwards data
	// from 'outside' to receiver
	go func() {
		for {
			u.inputChan <- <-u.cmdChan
		}
	}()

	// Start receiver
	go func() {
		for {
			u.outputChan <- <-u.inputChan
		}
	}()
}

func (u *Uart) Put(b byte) {
	u.cmdChan <- b
}

func (u *Uart) Take() byte {
	return <-u.outputChan
}

func main() {
	u := newUart()

	u.Start()

	u.Put(1)
	x := u.Take()

	fmt.Println(x)
}
