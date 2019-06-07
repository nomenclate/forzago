package forzago

import "sync"

type inputter interface {
	StartAccepting(q Queue)
}

type outputter interface {
	StartOutputting(q Queue)
}

type DataIn struct {
	i  inputter
	q  Queue
	o  outputter
	wg *sync.WaitGroup
}

func (p *DataIn) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

func NewDataIn(i inputter, q Queue, o outputter) *DataIn {
	var wg sync.WaitGroup
	wg.Add(1)
	return &DataIn{i: i, q: q, o: o, wg: &wg}
}
