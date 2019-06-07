package forzago

import "sync"

type inputter interface {
	StartAccepting(q Queue)
}

type outputter interface {
	StartOutputting(q Queue)
}

type ForzaDataIn struct {
	i  inputter
	q  Queue
	o  outputter
	wg *sync.WaitGroup
}

func (p *ForzaDataIn) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

func NewForzaDataIn(i inputter, q Queue, o outputter) *ForzaDataIn {
	var wg sync.WaitGroup
	wg.Add(1)
	return &ForzaDataIn{i: i, q: q, o: o, wg: &wg}
}
