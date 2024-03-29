package xorlist

import (
	"os"
	"strconv"
	"testing"
)

func TestXorList(t *testing.T) {
	xl := NewXorList()
	//debug.SetGCPercent(-1)

	// contains 1000000, ..., 3, 2, 1, 0
	SIZE, _ := strconv.Atoi(os.Args[1])
	var S2 = SIZE / 100
	for i := 0; i < SIZE; i++ {
		if xl.Size() != i {
			t.Error("PUSH error - Expected size", i+1, "get", xl.Size())
		}
		xl.PushFront(i)
	}

	// contains 0, 1, 2, 3, ..., 1000000
	xl.Reverse()
	xl.print()
	xl.Reverse()
	xl.print()
	xl.Reverse()

	size := xl.Size()
	next := SIZE - 1
	times := 0
	for xl.Size() > 0 {
		if xl.Size() != size {
			t.Error("POP error - Expected size", size, "get", xl.Size())
		}
		y := xl.PopBack().(int)
		if y != next {
			t.Error("POP error - Expected value", next, "get", y)
		}
		next--
		size--

		if times < 10 && xl.Size() < S2 {
			for i := S2 - 1; i < SIZE; i++ {
				xl.PushBack(i)
			}
			times++
			next = SIZE - 1
			size = SIZE
		}
	}

}
