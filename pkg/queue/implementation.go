package queue

import (
	"time"
)

var queue = make([]Message, 0)

func NewEnqueuer() Enqueue {
	return enqueuer{}
}

type enqueuer struct{}

func (e enqueuer) Enqueue(m Message) {
	queue = append(queue, m)
}

func NewDequeuer() Dequeue {
	return dequeuer{}
}

type dequeuer struct{}

func (e dequeuer) Dequeue() (message Message, hasMessage bool) {
	if len(queue) == 0 {
		return nil, false
	}

	msg := queue[0]
	queue = queue[1:]
	return msg, true
}

func ListenFor(dequeue Dequeue, message Message, handler func()) {
	for {
		m, has := dequeue.Dequeue()
		if !has {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if m == message {
			handler()
		}
	}
}
