package forzago

type Queue interface {
	Enqueue(data []byte)
	Dequeue() ([]byte, bool)
}
type ChannelQueue struct {
	c chan []byte
}

func NewChannelQueue() *ChannelQueue {
	return &ChannelQueue{
		c: make(chan []byte, 1000),
	}
}

func (q *ChannelQueue) Enqueue(data []byte) {
	select {
	case q.c <- data:
	}
}

func (q *ChannelQueue) Dequeue() ([]byte, bool) {
	select {
	case data := <-q.c:
		return data, true
	}
}
