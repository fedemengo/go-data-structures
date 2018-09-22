package xorlist

import (
	"fmt"
	"unsafe"

	"github.com/fedemengo/go-data-structures/queue"
)

// XorList is a type representing a Xor Linked List
type XorList struct {
	head   *xorListNode
	tail   *xorListNode
	antiGC []*xorListNode
	free   *queue.Queue
	size   int
}

type xorListNode struct {
	elem  interface{}
	np    *xorListNode
	index int
}

// NewXorList return the pointer to a XorList object
func NewXorList() *XorList {
	xl := new(XorList)
	xl.head = nil
	xl.tail = nil
	// avoid GC to clear nodes
	xl.antiGC = make([]*xorListNode, 0)
	xl.free = queue.NewQueue()
	xl.size = 0
	return xl
}

func newNode(x interface{}) *xorListNode {
	n := new(xorListNode)
	n.elem = x
	n.np = nil
	return n
}

func xor(l, r *xorListNode) *xorListNode {
	left := unsafe.Pointer(l)
	right := unsafe.Pointer(r)
	return (*xorListNode)(unsafe.Pointer(uintptr(left) ^ uintptr(right)))
}

func (xl *XorList) insert(left, right *xorListNode, x interface{}) *xorListNode {
	n := newNode(x)
	n.np = xor(left, right)
	if left != nil {
		left.np = xor(xor(left.np, right), n)
	}

	if right != nil {
		right.np = xor(xor(right.np, left), n)
	}

	if xl.free.Size() > 0 {
		n.index = xl.free.Pop().(int)
		xl.antiGC[n.index] = n
	} else {
		n.index = len(xl.antiGC)
		xl.antiGC = append(xl.antiGC, n)
	}
	xl.size++
	return n
}

func (xl *XorList) remove(left, right, node *xorListNode) *xorListNode {
	xl.free.Push(node.index)
	if left != nil {
		left.np = xor(xor(left.np, right), node)
	}

	if right != nil {
		right.np = xor(xor(right.np, left), node)
	}
	xl.size--
	return left
}

// PushFront pushes an element in front of the list
func (xl *XorList) PushFront(x interface{}) {
	xl.head = xl.insert(nil, xl.head, x)
	if xl.tail == nil {
		xl.tail = xl.head
	}
}

// PushBack pushes an element in the back of the list
func (xl *XorList) PushBack(x interface{}) {
	xl.tail = xl.insert(xl.tail, nil, x)
	if xl.head == nil {
		xl.head = xl.tail
	}
}

// PopFront pops an element from the front of the list
func (xl *XorList) PopFront() (x interface{}) {
	x = xl.head.elem
	xl.head = xl.remove(xl.head.np, nil, xl.head)
	return
}

// PopBack pops an element from the back of the list
func (xl *XorList) PopBack() (x interface{}) {
	x = xl.tail.elem
	xl.tail = xl.remove(xl.tail.np, nil, xl.tail)
	return
}

// Size return the size of the list
func (xl *XorList) Size() int {
	return xl.size
}

// Reverse a list in constant time
func (xl *XorList) Reverse() {
	xl.head, xl.tail = xl.tail, xl.head
}

func (xl *XorList) print() {
	fmt.Println("Printing")
	n := xl.head
	var last *xorListNode
	for n != nil {
		fmt.Println(n.elem.(int), n.index)
		last, n = n, xor(n.np, last)
	}
}
