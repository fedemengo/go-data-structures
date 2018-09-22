package main

import (
	"testing"

	"github.com/fedemengo/go-data-structures/queue"
	"github.com/fedemengo/go-data-structures/xorlist"
)

func XorList(size int) {
	SIZE := size
	S2 := SIZE / 100

	xl := xorlist.NewXorList()

	// contains 0, 1, 2, 3, ..., 1000000
	for i := 0; i < SIZE; i++ {
		xl.PushBack(i)
	}

	last := xl.Size()
	times := 0

	for xl.Size() > 0 {
		xl.PopFront()

		if times < 10 && xl.Size() < S2 {
			for i := last; i < last+SIZE; i++ {
				xl.PushBack(i)
			}
			times++
			last = last + SIZE
		}
	}
}

func benchmarkXorList(size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		XorList(size)
	}
}

func BenchmarkXorList1000(b *testing.B)   { benchmarkXorList(1000, b) }
func BenchmarkXorList10000(b *testing.B)  { benchmarkXorList(10000, b) }
func BenchmarkXorList100000(b *testing.B) { benchmarkXorList(100000, b) }

func Queue(size int) {
	SIZE := size
	S2 := SIZE / 100

	q := queue.NewQueue()

	// contains 0, 1, 2, 3, ..., 1000000
	for i := 0; i < SIZE; i++ {
		q.Push(i)
	}

	last := q.Size()
	times := 0

	for q.Size() > 0 {
		q.Pop()

		if times < 10 && q.Size() < S2 {
			for i := last; i < last+SIZE; i++ {
				q.Push(i)
			}
			times++
			last = last + SIZE
		}
	}
}

func benchmarkQueue(size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Queue(size)
	}
}

func BenchmarkQueue1000(b *testing.B)   { benchmarkQueue(1000, b) }
func BenchmarkQueue10000(b *testing.B)  { benchmarkQueue(10000, b) }
func BenchmarkQueue100000(b *testing.B) { benchmarkQueue(100000, b) }
