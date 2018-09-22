package queue

// Queue is a type represent a queue data structure
type Queue struct {
	container []interface{}
	head      int
	tail      int
}

// NewQueue return the pointer to a Queue object
func NewQueue() *Queue {
	q := new(Queue)
	q.container = make([]interface{}, 32)
	q.head = 0
	q.tail = 0
	return q
}

// Push an element in the queue
func (q *Queue) Push(x interface{}) {
	if q.nextIndex(q.tail) == q.head {
		q.resize()
	}
	q.container[q.tail] = x
	q.tail = q.nextIndex(q.tail)
}

// Pop an element form the queue
func (q *Queue) Pop() interface{} {
	x := q.Front()
	q.head = q.nextIndex(q.head)
	if q.Size() == (len(q.container) >> 2) {
		q.resize()
	}
	return x
}

// Front return the first element in the queue
func (q *Queue) Front() interface{} {
	x := q.container[q.head]
	return x
}

// Size return the size of queue
func (q *Queue) Size() (size int) {
	if q.tail >= q.head {
		size = q.tail - q.head
	} else {
		size = len(q.container) - q.head + q.tail
	}
	return size
}

func (q *Queue) nextIndex(index int) int {
	return (index + 1) & (len(q.container) - 1)
}

func (q *Queue) resize() {
	c := make([]interface{}, len(q.container)<<1)
	if q.tail >= q.head {
		copy(c, q.container[q.head:q.tail])
	} else {
		n := copy(c, q.container[q.head:])
		copy(c[n:], q.container[:q.tail])
	}
	q.head, q.tail = 0, q.Size()
	q.container = c
}
