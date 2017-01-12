package main

import (
	"fmt"
	"sync"
	//"time"
)

type Node struct {
	value  interface{}
	nxtPtr *Node
	prvPtr *Node
}

func NewNode(val interface{}) *Node {
	return &Node{
		value:  val,
		nxtPtr: nil,
		prvPtr: nil,
	}
}

func AddNode(val interface{}, previous *Node) *Node {
	newVal := NewNode(val)
	newVal.prvPtr = previous
	if previous != nil {
		previous.nxtPtr = newVal
	}
	return newVal
}

type Signal struct {
	sync.Mutex
	readPtr  *Node
	writePtr *Node
	event    chan struct{}
}

func NewSignal() *Signal {
	return &Signal{readPtr: nil, writePtr: nil, event: make(chan struct{})}
}

func (s *Signal) Get() interface{} {
	var x interface{}
	s.Lock()
	defer s.Unlock()
	if s.readPtr != nil {
		x = s.readPtr.value
	}
	return x
}

func (s *Signal) Put(x interface{}) {
	s.Lock()
	s.writePtr = AddNode(x, s.writePtr)
	s.Unlock()
}

func (s *Signal) Commit() {
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
	s2 := NewSignal()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		s1.Put(1)
		s1.Put(2)
		s1.Commit()
		s1.event <- struct{}{}
		<-s2.event

		s1.Put(3)
		s1.Put(4)
		s1.Commit()
		s1.event <- struct{}{}
		<-s2.event

		wg.Done()
	}()

	go func() {
		<-s1.event
		s2.Assign(s1)
		s2.event <- struct{}{}
		s2.Commit()
		fmt.Println(s2.Get())

		<-s1.event
		s2.Assign(s1)
		s2.event <- struct{}{}
		s2.Commit()
		fmt.Println(s2.Get())

		wg.Done()
	}()

	wg.Wait()

}
