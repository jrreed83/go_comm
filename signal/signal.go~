package main

import (
	"fmt"
	"sync"
	"time"
)

type Sample struct {
	tick uint32
	x    byte
}

type Signal struct {
	next   Sample
	curr   Sample
	prev   Sample
	buffer chan Sample
	change chan struct{}
	read   chan struct{}
	write  chan struct{}
}

// Just waits for change event to occur
func process(s *Signal) {
	<-s.change
}

// Check whether signal is at rising edge
func rising_edge(s *Signal) bool {
	var x bool
	<-s.read
	x = (s.curr == 1) && (s.prev == 0)
	return x
}

// Check whether signal is at falling edge
func rising_edge(s *Signal) bool {
	var x bool
	<-s.read
	x = (s.curr == 0) && (s.prev == 1)
	return x
}

func (s *Signal) Read() Sample {
	s.Lock()
	samp := s.curr
	s.Unlock()
	return samp
}

func (s *Signal) Write(samp Sample) {
	s.Lock()
	s.next = samp
	s.Unlock()

}

func (s *Signal) Update() {
	s.Lock()
	s.curr = s.next
	s.Unlock()
}
func main() {
	fmt.Println("vim-go")
}
