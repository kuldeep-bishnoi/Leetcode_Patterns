package topk

import (
	"container/heap"
)

// MinHeap implementation for finding Kth largest element
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

// MaxHeap implementation for finding Kth smallest element
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
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

// FindKthLargest finds the Kth largest element in an array
// Time Complexity: O(N log K) where N is the array size and K is the parameter
// Space Complexity: O(K) for storing the min heap
func FindKthLargest(nums []int, k int) int {
	// Create a min heap
	h := &MinHeap{}
	heap.Init(h)

	// Process each element in the array
	for _, num := range nums {
		// If heap size is less than k, add the element
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			// If current element is larger than the smallest in heap,
			// remove the smallest and add the current element
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	// The root of the min heap is the kth largest element
	return (*h)[0]
}

// FindKthSmallest finds the Kth smallest element in an array
// Time Complexity: O(N log K) where N is the array size and K is the parameter
// Space Complexity: O(K) for storing the max heap
func FindKthSmallest(nums []int, k int) int {
	// Create a max heap
	h := &MaxHeap{}
	heap.Init(h)

	// Process each element in the array
	for _, num := range nums {
		// If heap size is less than k, add the element
		if h.Len() < k {
			heap.Push(h, num)
		} else if num < (*h)[0] {
			// If current element is smaller than the largest in heap,
			// remove the largest and add the current element
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	// The root of the max heap is the kth smallest element
	return (*h)[0]
}

// FindKthLargestQuickSelect uses the QuickSelect algorithm to find the Kth largest element
// Time Complexity: O(N) average case, O(N²) worst case
// Space Complexity: O(1) as we're doing it in-place
func FindKthLargestQuickSelect(nums []int, k int) int {
	// Convert k to 0-indexed (k-1 th largest = n-k th smallest)
	k = len(nums) - k

	return quickSelect(nums, 0, len(nums)-1, k)
}

// quickSelect is a helper function that implements the QuickSelect algorithm
func quickSelect(nums []int, start, end, k int) int {
	// Base case: if the list contains only one element
	if start == end {
		return nums[start]
	}

	// Partition the array and get the pivot index
	pivotIndex := partition(nums, start, end)

	// If the pivot is the kth element, return it
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		// If k is less than the pivot index, search the left part
		return quickSelect(nums, start, pivotIndex-1, k)
	} else {
		// If k is greater than the pivot index, search the right part
		return quickSelect(nums, pivotIndex+1, end, k)
	}
}

// partition is a helper function used by quickSelect
func partition(nums []int, start, end int) int {
	// Choose the rightmost element as pivot
	pivot := nums[end]

	// Index of smaller element
	i := start

	// Traverse through all elements
	// compare each element with pivot
	for j := start; j < end; j++ {
		// If current element is smaller than the pivot
		if nums[j] < pivot {
			// Swap elements at i and j
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// Swap the pivot element with the element at i
	nums[i], nums[end] = nums[end], nums[i]

	// Return the pivot index
	return i
}

// Example usage:
//
// Find Kth Largest:
// Input: nums = [3, 2, 1, 5, 6, 4], k = 2
// Output: 5 (the 2nd largest element is 5)
//
// Input: nums = [3, 2, 3, 1, 2, 4, 5, 5, 6], k = 4
// Output: 4 (the 4th largest element is 4)
//
// Find Kth Smallest:
// Input: nums = [3, 2, 1, 5, 6, 4], k = 2
// Output: 2 (the 2nd smallest element is 2)
//
// Input: nums = [3, 2, 3, 1, 2, 4, 5, 5, 6], k = 4
// Output: 3 (the 4th smallest element is 3)
