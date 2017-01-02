package main

import "fmt"
import "errors"
import "time"

type Register struct {
	inputChan  chan byte
	outputChan chan byte
	state      byte
}

func NewRegister() *Register {
	return &Register{
		inputChan:  make(chan byte),
		outputChan: make(chan byte),
		state:      0}
}

func (r *Register) Write(x byte) error {
	select {
	case r.inputChan <- x:
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("Writing to register timed-out")
	}
}

func (r *Register) Read() (byte, error) {
	var x byte
	select {
	case x = <-r.outputChan:
		return x, nil
	case <-time.After(1 * time.Second):
		return x, errors.New("Read from register timed-out")
	}
}

func (r *Register) Start() {
	go func() {
		for {
			select {
			case x := <-r.inputChan:
				r.state = x
			case r.outputChan <- r.state:
			}
		}
	}()
}

func main() {
	var err error
	var x byte

	r := NewRegister()
	r.Start()

	if err = r.Write(45); err != nil {
		fmt.Println("No Error")
	}
	err = r.Write(65)
	err = r.Write(64)
	err = r.Write(23)
	x, err = r.Read()
	fmt.Println(x)
}
