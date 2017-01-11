package main

import (
	"fmt"
	"sync"
	"time"
)

type TVar struct {
	value  interface{}
	nxtPtr *TVar
	prvPtr *TVar
}

func NewTVar(val interface{}) *TVar {
	return &TVar{
		value:  val,
		nxtPtr: nil,
		prvPtr: nil,
	}
}

func AddTVar(val interface{}, previous *TVar) *TVar {
	newVal := NewTVar(val)
	newVal.prvPtr = previous
	if previous != nil {
		previous.nxtPtr = newVal
	}
	return newVal
}

type Signal struct {
	sync.RWMutex
	readPtr  *TVar
	writePtr *TVar
	event    chan struct{}
}

func NewSignal() *Signal {
	return &Signal{readPtr: nil, writePtr: nil, event: make(chan struct{}, 2)}
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
	s.writePtr = AddTVar(x, s.writePtr)
	s.RUnlock()
}

func (s *Signal) Commit() {
	s.Lock()
	s.readPtr = s.writePtr
	s.Unlock()
	select {
	case s.event <- struct{}{}:
	case <-time.After(time.Millisecond):
	}

}

func (s *Signal) Notify() {
	select {
	case s.event <- struct{}{}:
	case <-time.After(time.Millisecond):
	}
}

func (s *Signal) Assign(s1 *Signal) {
	x := s1.Get()
	s.Put(x)
}

func main() {
	s1 := NewSignal()
	s2 := NewSignal()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		s1.Put(0)
		s1.Put(1)
		s1.Commit()

		s1.Put(0)
		s1.Put(1)
		s1.Commit()

		wg.Done()
	}()

	go func() {
		<-s1.event
		s2.Put(5)
		s2.Commit()
		//s2.Notify()

		<-s1.event
		s2.Put(8)
		s2.Commit()
		//s2.Notify()

		wg.Done()
	}()

	wg.Wait()
	//<-s2.event
	//<-s2.event
	//	time.Sleep(time.Second)

	fmt.Println(s2.writePtr.value)
	fmt.Println(s2.writePtr.prvPtr.value)
}
