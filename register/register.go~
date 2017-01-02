package main

import "fmt"
import "errors"
import "time"
import "sync"

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
	//var err error
	var x byte

	var wg sync.WaitGroup
	wg.Add(2)

	r := NewRegister()
	r.Start()

	go func() {
		_ = r.Write(1)
		time.Sleep(10 * time.Millisecond)

		_ = r.Write(2)
		time.Sleep(10 * time.Millisecond)

		_ = r.Write(3)
		time.Sleep(10 * time.Millisecond)

		_ = r.Write(4)
		time.Sleep(10 * time.Millisecond)

		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Millisecond)
		x, _ = r.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x, _ = r.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x, _ = r.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x, _ = r.Read()
		fmt.Println(x)

		wg.Done()
	}()

	wg.Wait()
}
