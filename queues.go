package queues

import (
	"github.com/felts94/stacks"
)

// Queue implemented with felts94/stacks
type Queue struct {
	Inverted bool //always init as true
	S        stacks.Stack
}

// Enqueue a new item into the queue
func (q *Queue) Enqueue(elem interface{}) {
	if q.Inverted {
		q.S.Push(elem)
	} else {
		temp := stacks.Stack{}
		telem := q.S.Pop()
		for telem != nil {
			temp.Push(telem)
			telem = q.S.Pop()
		}
		temp.Push(elem)
		q.S = temp
		q.Inverted = !q.Inverted
	}
}

// Dequeue an item and return it
func (q *Queue) Dequeue() interface{} {
	if q.Inverted {
		temp := stacks.Stack{}
		elem := q.S.Pop()
		for elem != nil {
			temp.Push(elem)
			elem = q.S.Pop()
		}
		elem = temp.Pop()
		q.S = temp
		q.Inverted = !q.Inverted
		return elem
	}
	return q.S.Pop()
}

// NewQueuePointer create empty queue
func NewQueuePointer() *Queue {
	return &Queue{
		Inverted: true,
		S:        stacks.Stack{},
	}
}
