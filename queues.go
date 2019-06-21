package queues

import (
	"github.com/felts94/stacks"
)

// Queue implemented with felts94/stacks
type Queue struct {
	Inverted bool //always init as true
	S        stacks.Stack
}

func (q *Queue) flip() {
	temp := stacks.Stack{}
	telem := q.S.Pop()
	for telem != nil {
		temp.Push(telem)
		telem = q.S.Pop()
	}
	q.S = temp
	q.Inverted = !q.Inverted
}

// Enqueue a new item into the queue
func (q *Queue) Enqueue(elem interface{}) {
	if q == nil {
		return
	}
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

// View the contents of the queue without changing them
func (q *Queue) View(start, num *int) []interface{} {

	if q == nil {
		return nil
	}

	if q.S == nil {
		return nil
	}

	// get start and end args
	var n, s int
	if start == nil {
		s = 0
	}
	if num == nil {
		n = len(q.S)
	}

	// flip stack if needed
	if q.Inverted {
		q.flip()
	}

	// create new return var
	var contents []interface{}
	for i := s; i < n; i++ {
		contents = append(contents, q.S[i])
	}

	return contents
}

// NewQueuePointer create empty queue
func NewQueuePointer() *Queue {
	return &Queue{
		Inverted: true,
		S:        stacks.Stack{},
	}
}
