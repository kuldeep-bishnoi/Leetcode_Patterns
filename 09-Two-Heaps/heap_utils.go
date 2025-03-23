package twoheaps

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Peek() int {
	if h.Len() == 0 {
		return 0 // Or a suitable default value
	}
	return (*h)[0]
}

// MaxHeap implementation
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Note the > instead of <
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() int {
	if h.Len() == 0 {
		return 0 // Or a suitable default value
	}
	return (*h)[0]
}

// Generic heap implementation for more complex data types

// Item is an interface that implements priority comparison
type Item interface {
	Less(other Item) bool
}

// GenericMinHeap implements a min heap for any type that satisfies the Item interface
type GenericMinHeap []Item

func (h GenericMinHeap) Len() int           { return len(h) }
func (h GenericMinHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h GenericMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GenericMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *GenericMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *GenericMinHeap) Peek() Item {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

// GenericMaxHeap implements a max heap for any type that satisfies the Item interface
type GenericMaxHeap []Item

func (h GenericMaxHeap) Len() int           { return len(h) }
func (h GenericMaxHeap) Less(i, j int) bool { return !h[i].Less(h[j]) } // Invert the comparison
func (h GenericMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GenericMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *GenericMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *GenericMaxHeap) Peek() Item {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}
