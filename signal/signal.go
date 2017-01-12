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
	sync.RWMutex
	readPtr  *Node
	writePtr *Node
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
	s.writePtr = AddNode(x, s.writePtr)
	s.RUnlock()
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

	wg.Add(3)

	go func() {
		s1.Put(1)
		s1.Put(2)
		s1.Commit()
		s1.event <- struct{}{}

		s1.Put(3)
		s1.Put(4)
		s1.Commit()
		s1.event <- struct{}{}

		wg.Done()
	}()

	go func() {
		<-s1.event
		fmt.Println(s1.Get())
		s2.Assign(s1)
		s2.Commit()

		<-s1.event
		fmt.Println(s1.Get())
		s2.Assign(s1)
		s2.Commit()

		wg.Done()
	}()

	go func() {
		//		<-s2.event
		//	fmt.Println(s2.Get())

		//		<-s2.event
		//	fmt.Println(s2.Get())

		wg.Done()
	}()
	wg.Wait()
	//<-s2.event
	//<-s2.event
	//	time.Sleep(time.Second)

	//fmt.Println(s2.readPtr.value)
	//fmt.Println(s2.readPtr.prvPtr.value)
	//fmt.Println(s1.writePtr.prvPtr.prvPtr.value)

}
