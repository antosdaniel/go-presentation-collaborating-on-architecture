package queue

type Message any

type Enqueue interface {
	Enqueue(message Message)
}

type Dequeue interface {
	Dequeue() (message Message, hasMessage bool)
}
