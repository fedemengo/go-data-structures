package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap(func(e1, e2 Elem) bool {
		return e1.Key.(int) > e2.Key.(int)
	})

	elems := make([]Elem, 0)
	for i := 0; i < 10000; i++ {
		if h.Size() != i {
			t.Error("PUSH error - Expected size", i+1, "get", h.Size())
		}
		h.Push(Elem{Key: i, Val: "hello"})
		elems = append(elems, Elem{Key: i, Val: "hi"})
	}

	size := h.Size()
	last := h.Size() - 1
	for h.Size() > 0 {
		if h.Size() != size {
			t.Error("POP error - Expected size", size, "get", h.Size())
		}
		y := h.Front().Key.(int)
		x := h.Pop().Key.(int)
		if y != last {
			t.Error("Front error - Expected value", last, "get", y)
		}
		if x != last {
			t.Error("POP error - Expected value", last, "get", x)
		}
		last--
		size--
	}

	h2 := BuildHeap(&elems, func(e1, e2 Elem) bool {
		return e1.Key.(int) < e2.Key.(int)
	})

	last = 0
	for h2.Size() > 0 {
		x := h2.Pop().Key.(int)
		if x != last {
			t.Error("POP error - Expected value", last, "get", x)
		}
		last++
	}
}
