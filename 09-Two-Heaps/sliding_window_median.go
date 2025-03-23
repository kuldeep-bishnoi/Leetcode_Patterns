package twoheaps

import (
	"container/heap"
)

// MedianSlidingWindow calculates the median of all sliding windows of size k in the array
// Time Complexity: O(n * log k) where n is the length of the array and k is the window size
// Space Complexity: O(k) for storing elements in both heaps
func MedianSlidingWindow(nums []int, k int) []float64 {
	result := make([]float64, 0, len(nums)-k+1)
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return result
	}

	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	// HashMap to count elements to be removed
	removed := make(map[int]int)

	// Process the first k elements
	for i := 0; i < k; i++ {
		addNum(nums[i], maxHeap, minHeap)
	}

	// Calculate the median of the first window
	result = append(result, findMedian(maxHeap, minHeap))

	// Process the rest of the elements
	for i := k; i < len(nums); i++ {
		// Add the element being moved into the window
		addNum(nums[i], maxHeap, minHeap)

		// Mark the element leaving the window for removal
		outNum := nums[i-k]
		removed[outNum]++

		// Balance the heaps after removal
		// We delay the actual removal until necessary to keep the code efficient
		balance(maxHeap, minHeap, removed)

		// Calculate the median of the current window
		result = append(result, findMedian(maxHeap, minHeap))
	}

	return result
}

// addNum adds a number to the appropriate heap
func addNum(num int, maxHeap *MaxHeap, minHeap *MinHeap) {
	if maxHeap.Len() == 0 || num <= maxHeap.Peek() {
		heap.Push(maxHeap, num)
	} else {
		heap.Push(minHeap, num)
	}

	// Balance heaps
	rebalanceHeaps(maxHeap, minHeap)
}

// balance both heaps and remove marked elements as needed
func balance(maxHeap *MaxHeap, minHeap *MinHeap, removed map[int]int) {
	// Clean up maxHeap, removing marked elements from the top
	for maxHeap.Len() > 0 && removed[maxHeap.Peek()] > 0 {
		removed[maxHeap.Peek()]--
		heap.Pop(maxHeap)
	}

	// Clean up minHeap, removing marked elements from the top
	for minHeap.Len() > 0 && removed[minHeap.Peek()] > 0 {
		removed[minHeap.Peek()]--
		heap.Pop(minHeap)
	}

	// Rebalance the heaps
	rebalanceHeaps(maxHeap, minHeap)
}

// rebalanceHeaps ensures that maxHeap has either the same number of elements as minHeap,
// or one more element than minHeap
func rebalanceHeaps(maxHeap *MaxHeap, minHeap *MinHeap) {
	// If maxHeap has more than one extra element than minHeap
	if maxHeap.Len() > minHeap.Len()+1 {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}
	// If minHeap has more elements than maxHeap
	if minHeap.Len() > maxHeap.Len() {
		heap.Push(maxHeap, heap.Pop(minHeap))
	}
}

// findMedian calculates the median from the two heaps
func findMedian(maxHeap *MaxHeap, minHeap *MinHeap) float64 {
	if maxHeap.Len() == minHeap.Len() {
		return float64(maxHeap.Peek()+minHeap.Peek()) / 2.0
	}
	return float64(maxHeap.Peek())
}

// Example usage:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
// Explanation:
// Window 1 = [1,3,-1] -> median = 1
// Window 2 = [3,-1,-3] -> median = -1
// Window 3 = [-1,-3,5] -> median = -1
// Window 4 = [-3,5,3] -> median = 3
// Window 5 = [5,3,6] -> median = 5
// Window 6 = [3,6,7] -> median = 6
//
// Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
// Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]
