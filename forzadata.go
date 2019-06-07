// Package forzago provides a library, set of interfaces and a client for interacting with the
// DataOut feature of Forza Motorsport 7.

package forzago

import "sync"

// Queue is an interface that groups the Dequeuer and Enqueuer methods
type Queue interface {
	Dequeuer
	Enqueuer
}

// Dequeuer wraps the Dequeue method.
//
// Dequeue retrieves a []byte from the the underlying Queue
type Dequeuer interface {
	Dequeue() ([]byte, bool)
}

// Enqueuer wraps the Enqueue method.
//
// Enqueue places a []byte into the underlying Queue
type Enqueuer interface {
	Enqueue(data []byte)
}

// Inputter wraps the method StartAccepting
//
// StartAccepting will cause the assocaiated inputter to start
// placing items on to the Queue q.
type Inputter interface {
	StartAccepting(q Queue)
}

// Outputter wraps the method StartOutputting
//
// StartOutputting will cause the associated outputter to start
// consuming items from the Queue q so they may be output as
// appropriate.
type Outputter interface {
	StartOutputting(q Queue)
}

// Pipeline is a collection of a specific inputter, outputter and queue
// used for pipelining of a strean of forza data
type Pipeline struct {
	i  Inputter
	q  Queue
	o  Outputter
	wg *sync.WaitGroup
}

// Start is used to start the processing of data by the pipeline.
// It will start the inputter and outputter by in thier own goroutines
// by calling their respective StartAccepting and StartOutputting. A wait group
// is used to track the goroutines.
func (p *Pipeline) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

// NewPipeline returns a new Pipeline using i, q, and o as the respective
// Inputter, Queue and Outputter.
func NewPipeline(i Inputter, q Queue, o Outputter) *Pipeline {
	var wg sync.WaitGroup
	wg.Add(1)
	return &Pipeline{i: i, q: q, o: o, wg: &wg}
}
