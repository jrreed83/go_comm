package port

import "sync"

type Port struct {
	sync.Mutex
	reg byte
}

func NewPort() *Port {
	return &Port{reg: 0}
}

func (p *Port) Write(x byte) {
	p.Lock()
	p.reg = x
	p.Unlock()
}

func (p *Port) Read() byte {
	var x byte
	p.Lock()
	x = p.reg
	p.Unlock()
	return x
}
