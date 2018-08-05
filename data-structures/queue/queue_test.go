package queue

import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 1000000; i++ {
		if q.Size() != i {
			t.Error("PUSH error - Expected size", i+1, "get", q.Size())
		}
		q.Push(i)
	}
	size := q.Size()
	last := -1
	fmt.Println(size)
	for q.Size() > 0 {
		if q.Size() != size {
			t.Error("POP error - Expected size", size, "get", q.Size())
		}
		y := q.Front().(int)
		x := q.Pop().(int)
		if y != last+1 {
			t.Error("Front error - Expected value", last+1, "get", y)
		}
		if x != last+1 {
			t.Error("POP error - Expected value", last+1, "get", x)
		}
		last = x
		size--
	}
}
