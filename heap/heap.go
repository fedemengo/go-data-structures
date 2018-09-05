package heap

// Elem is the wrapper around each heap elements
type Elem struct {
	Key interface{}
	Val interface{}
}

// Heap is a type that represents a heap data structure
type Heap struct {
	container []Elem
	cmp       func(Elem, Elem) bool
	size      int
}

// NewHeap return the pointer to a Heap object
func NewHeap(cmp func(e1, e2 Elem) bool) *Heap {
	h := new(Heap)
	h.container = make([]Elem, 32)
	h.cmp = cmp
	h.size = 0
	return h
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func (h *Heap) updateKey(i int) {
	tmp := h.container[i]
	for i > 0 && !h.cmp(h.container[parent(i)], tmp) {
		h.container[i] = h.container[parent(i)]
		i = parent(i)
	}
	h.container[i] = tmp
}

func (h *Heap) heapify(i int) {
	index := i
	l := left(i)
	r := l + 1

	if l < h.size && h.cmp(h.container[l], h.container[index]) {
		index = l
	}

	if r < h.size && h.cmp(h.container[r], h.container[index]) {
		index = r
	}

	if index != i {
		tmp := h.container[i]
		h.container[i] = h.container[index]
		h.container[index] = tmp
		h.heapify(index)
	}
}

func (h *Heap) resize() {
	c := make([]Elem, len(h.container)<<1)
	copy(c, h.container)
	h.container = c
}

// Push an element in the Heap
func (h *Heap) Push(x Elem) {
	h.container[h.size] = x
	h.updateKey(h.size)
	h.size++
	if h.size == len(h.container) {
		h.resize()
	}
}

// Pop an element from the Heap
func (h *Heap) Pop() Elem {
	x := h.Front()
	h.size--
	h.container[0] = h.container[h.size]
	h.heapify(0)
	if h.size == (len(h.container) >> 2) {
		h.resize()
	}
	return x
}

// Front return the first element in the Heap
func (h *Heap) Front() Elem {
	x := h.container[0]
	return x
}

// BuildHeap builds a heap in linear time from a slice of Elem
func BuildHeap(elems *[]Elem, cmp func(e1, e2 Elem) bool) *Heap {
	h := new(Heap)
	h.container = make([]Elem, len(*elems))
	copy(h.container, *elems)
	h.cmp = cmp
	h.size = len(*elems)

	for i := h.size / 2; i > -1; i-- {
		h.heapify(i)
	}
	return h
}

// Size return the size of Heap
func (h *Heap) Size() (size int) {
	return h.size
}
