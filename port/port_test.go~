package port

import (
	"sync"
	"testing"
	"time"
)

func TestPort(t *testing.T) {
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
		t.Log(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		t.Log(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		t.Log(x)

		time.Sleep(10 * time.Millisecond)
		x = p.Read()
		t.Log(x)

		wg.Done()
	}()

	wg.Wait()
}
