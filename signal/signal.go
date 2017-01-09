package main

import (
	"fmt"
	"sync"
)

type Sample struct {
	t uint32
	x byte
}

type TVal struct {
	val    byte
	nxtPtr *TVal
	prvPtr *TVal
}

func NewTVal(val byte, previous *TVal) *TVal {
	newVal := TVal{val: val, nxtPtr: nil, prvPtr: nil}
	newVal.prvPtr = previous
	if previous != nil {
		previous.nxtPtr = &newVal
	}
	return &newVal
}

type Signal struct {
	sync.Mutex
	readPtr  *TVal
	writePtr *TVal
}

func NewSignal(x byte) *Signal {
	tval := NewTVal(x, nil)
	return &Signal{readPtr: tval, writePtr: tval}
}

func (s *Signal) Read() byte {
	var x byte
	s.Lock()
	x = s.readPtr.val
	s.Unlock()
	return x
}

func (s *Signal) Write(x byte) {
	s.Lock()
	s.writePtr = NewTVal(x, s.writePtr)
	s.Unlock()
}

func (s *Signal) Update() {
	s.Lock()
	s.readPtr = s.writePtr
	s.Unlock()
}

func (s *Signal) Assign(s1 *Signal) {
	x := s1.Read()
	s.Write(x)
}

func main() {
	s1 := NewSignal(0)
	s1.Write(1)
	s1.Update()

	s2 := NewSignal(0)
	s2.Assign(s1)
	s2.Update()
	fmt.Println(s2.Read())
}
