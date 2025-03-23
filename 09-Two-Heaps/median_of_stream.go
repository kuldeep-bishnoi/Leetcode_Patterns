package twoheaps

import (
	"container/heap"
)

// MedianFinder keeps track of the median in a stream of numbers
type MedianFinder struct {
	maxHeap *MaxHeap // Stores the smaller half of numbers
	minHeap *MinHeap // Stores the larger half of numbers
}

// NewMedianFinder initializes a new MedianFinder
func NewMedianFinder() *MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	return &MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

// AddNum adds a number to the data stream
// Time Complexity: O(log n) where n is the number of elements
// Space Complexity: O(n) for storing all elements
func (mf *MedianFinder) AddNum(num int) {
	// Step 1: Add the number to the appropriate heap
	if mf.maxHeap.Len() == 0 || num <= mf.maxHeap.Peek() {
		heap.Push(mf.maxHeap, num)
	} else {
		heap.Push(mf.minHeap, num)
	}

	// Step 2: Balance the heaps
	// If maxHeap has more than one extra element than minHeap
	if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	}
	// If minHeap has more elements than maxHeap
	if mf.minHeap.Len() > mf.maxHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

// FindMedian returns the median of all elements so far
// Time Complexity: O(1)
// Space Complexity: O(1)
func (mf *MedianFinder) FindMedian() float64 {
	// If both heaps have the same size, return the average of their tops
	if mf.maxHeap.Len() == mf.minHeap.Len() {
		return float64(mf.maxHeap.Peek()+mf.minHeap.Peek()) / 2.0
	}
	// Otherwise, the median is the top of maxHeap (which has one more element)
	return float64(mf.maxHeap.Peek())
}

// Example usage:
//
// Input: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//        [[], [1], [2], [], [3], []]
// Output: [null, null, null, 1.5, null, 2.0]
// Explanation:
// mf = NewMedianFinder()
// mf.AddNum(1)    // arr = [1]
// mf.AddNum(2)    // arr = [1, 2]
// mf.FindMedian() // return 1.5 (i.e., (1 + 2) / 2)
// mf.AddNum(3)    // arr = [1, 2, 3]
// mf.FindMedian() // return 2.0
//
// Input: ["MedianFinder", "addNum", "findMedian", "addNum", "findMedian"]
//        [[], [5], [], [2], []]
// Output: [null, null, 5.0, null, 3.5]
// Explanation:
// mf = NewMedianFinder()
// mf.AddNum(5)    // arr = [5]
// mf.FindMedian() // return 5.0
// mf.AddNum(2)    // arr = [2, 5]
// mf.FindMedian() // return 3.5
