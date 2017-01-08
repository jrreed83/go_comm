package main

import (
	"fmt"
)

type TVal struct {
	val byte
	nxt *TVal
	prv *TVal
}

func NewTVal(x byte, prv *TVal) *TVal {
	newVal := TVal{val: x, nxt: nil, prv: nil}

	newVal.prv = prv

	if prv != nil {
		prv.nxt = &newVal
	}

	return &newVal
}

func main() {
	x := NewTVal(4, nil)
	y := NewTVal(6, x)
	fmt.Println(y.prv.val)
}
