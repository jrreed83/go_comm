package main

import "fmt"

//import "errors"
import "time"
import "sync"

type Port struct {
	sync.Mutex
	reg byte
}

func NewPort() *Port {
	return &Port{reg: 0}
}

func (p *Port) Write(x byte) {
	p.Lock()
	p.reg = x
	p.Unlock()
}

func (p *Port) Read() byte {
	var x byte
	p.Lock()
	x = p.reg
	p.Unlock()
	return x
}

func main() {

	var x byte

	var wg sync.WaitGroup
	wg.Add(2)
	p := NewPort()

	go func() {
		p.Write(1)
		time.Sleep(10 * time.Millisecond)

		p.Write(2)
		time.Sleep(10 * time.Millisecond)

		p.Write(3)
		time.Sleep(10 * time.Millisecond)

		p.Write(4)
		time.Sleep(10 * time.Millisecond)

		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Millisecond)
		x = p.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		fmt.Println(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		fmt.Println(x)

		wg.Done()
	}()

	wg.Wait()

}
