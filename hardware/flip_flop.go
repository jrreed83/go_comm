package main

import "fmt"

func flip_flop(port0 chan uint8, port1 chan uint8) {
	var state uint8 = 0
	for {
		in := <-port0
		port1 <- state
		state = in
	}
}

func main() {
	port0 := make(chan uint8)
	port1 := make(chan uint8)

	go flip_flop(port0, port1)

	bits := []uint8{0, 1, 0, 1, 0, 1, 0, 1}

	for _, x := range bits {
		port0 <- x
		y := <-port1
		fmt.Print(y)
	}
}
