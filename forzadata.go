package forzago

import "sync"

type Queue interface {
	Enqueue(data []byte)
	Dequeue() ([]byte, bool)
}

type Inputter interface {
	StartAccepting(q Queue)
}

type Outputter interface {
	StartOutputting(q Queue)
}

type Pipeline struct {
	i  Inputter
	q  Queue
	o  Outputter
	wg *sync.WaitGroup
}

func (p *Pipeline) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

func NewPipeline(i Inputter, q Queue, o Outputter) *Pipeline {
	var wg sync.WaitGroup
	wg.Add(1)
	return &Pipeline{i: i, q: q, o: o, wg: &wg}
}
