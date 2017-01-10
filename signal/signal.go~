package main

import (
	"fmt"
	"sync"
	//"time"
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
	sync.RWMutex
	readPtr  *TVar
	writePtr *TVar
	event    chan struct{}
}

func NewSignal() *Signal {
	return &Signal{readPtr: nil, writePtr: nil, event: make(chan struct{})}
}

func (s *Signal) Get() interface{} {
	var x interface{}
	s.RLock()
	if s.readPtr != nil {
		x = s.readPtr.value
	}
	s.RUnlock()
	return x
}

func (s *Signal) Put(x interface{}) {
	s.RLock()
	s.writePtr = NewTVar(x, s.writePtr)
	s.RUnlock()
}

func (s *Signal) Commit() {
	s.Lock()
	s.readPtr = s.writePtr
	s.Unlock()
	s.event <- struct{}{}
}

func (s *Signal) Assign(s1 *Signal) {
	x := s1.Get()
	s.Put(x)
}

func main() {
	s1 := NewSignal()
	s2 := NewSignal()

	//go func() {
	//	s1.Put(0)
	//	s1.Put(1)
	//	s1.Commit()

	//s1.Put(0)
	//s1.Put(1)
	//s1.Commit()
	//}()

	go func() {
		<-s1.event
		s2.Put(1)
		//s2.Commit()

		//<-s1.event
		//s2.Put(1)
		//s2.Commit()
	}()

	s1.Put(0)
	s1.Put(1)
	s1.Commit()

	fmt.Println(s2.writePtr.value)
}
