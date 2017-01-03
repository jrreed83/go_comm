package main

import "fmt"

//import "errors"
import "time"
import "sync"

type Reg struct {
	sync.Mutex
	state byte
}

func NewReg() *Reg {
	return &Reg{state: 0}
}

func (r *Reg) Write(x byte) error {
	r.Lock()
	r.state = x
	r.Unlock()
	return nil
}

func (r *Reg) Read() (byte, error) {
	var x byte
	r.Lock()
	x = r.state
	r.Unlock()
	return x, nil
}

type Register struct {
	lockChan chan struct{}
	state    byte
}

func NewRegister() *Register {
	r := Register{
		lockChan: make(chan struct{}, 1),
		state:    0}
	r.lockChan <- struct{}{}
	return &r
}

func (r *Register) Write(x byte) error {
	<-r.lockChan
	r.state = x
	r.lockChan <- struct{}{}
	return nil
}

func (r *Register) Read() (byte, error) {
	var x byte
	<-r.lockChan
	x = r.state
	r.lockChan <- struct{}{}
	return x, nil
}

func main() {

	var x byte

	var wg sync.WaitGroup
	wg.Add(2)
	r := NewRegister()

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
