package main

import (
	"fmt"
	"sync"
)

type TVar struct {
	value  interface{}
	nxtPtr *TVar
	prvPtr *TVar
}

func NewTVar(val interface{}, previous *TVar) *TVar {
	newVal := TVar{value: val, nxtPtr: nil, prvPtr: nil}
	newVal.prvPtr = previous
	if previous != nil {
		previous.nxtPtr = &newVal
	}
	return &newVal
}

type Signal struct {
	sync.Mutex
	readPtr  *TVar
	writePtr *TVar
}

func NewSignal() *Signal {
	return &Signal{readPtr: nil, writePtr: nil}
}

func (s *Signal) Get() interface{} {
	var x interface{}
	s.Lock()
	if s.readPtr != nil {
		x = s.readPtr.value
	}
	s.Unlock()
	return x
}

func (s *Signal) Put(x interface{}) {
	s.Lock()
	s.writePtr = NewTVar(x, s.writePtr)
	if s.readPtr == nil {
		s.readPtr = s.writePtr
	}
	s.Unlock()
}

func (s *Signal) Update() {
	s.Lock()
	s.readPtr = s.writePtr
	s.Unlock()
}

func (s *Signal) Assign(s1 *Signal) {
	x := s1.Get()
	s.Put(x)
}

func main() {
	s1 := NewSignal()
	s1.Put(0)
	s1.Put("hello")
	s1.Update()

	s2 := NewSignal()
	s2.Assign(s1)
	s2.Update()
	fmt.Println(s2.Get())
}
