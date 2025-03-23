package kwaymerge

import (
	"container/heap"
	"math"
)

// ListElement represents an element in one of the lists
type ListElement struct {
	Value     int // The actual value
	ListIndex int // Index of the list this element belongs to
	ElemIndex int // Index of this element in its list
}

// ListElementMinHeap implementation
type ListElementMinHeap []ListElement

func (h ListElementMinHeap) Len() int           { return len(h) }
func (h ListElementMinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h ListElementMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListElementMinHeap) Push(x interface{}) {
	*h = append(*h, x.(ListElement))
}

func (h *ListElementMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SmallestRange finds the smallest range that includes at least one number
// from each of the k lists
// Time Complexity: O(N log K) where N is the total number of elements across all lists
// and K is the number of lists
// Space Complexity: O(K) for the heap
func SmallestRange(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{0, 0}
	}

	// Create a min heap
	h := &ListElementMinHeap{}
	heap.Init(h)

	// Initialize variables to track the range
	minRange := math.MaxInt32
	rangeStart, rangeEnd := 0, 0

	// Track the maximum element in the heap
	maxVal := math.MinInt32

	// Insert the first element from each list into the heap
	for i, list := range nums {
		if len(list) > 0 {
			heap.Push(h, ListElement{
				Value:     list[0],
				ListIndex: i,
				ElemIndex: 0,
			})
			maxVal = max(maxVal, list[0])
		}
	}

	// Process the heap until any list is exhausted
	for h.Len() == len(nums) {
		// Get the minimum element
		minElement := heap.Pop(h).(ListElement)

		// Calculate the current range
		currentRange := maxVal - minElement.Value

		// Update the smallest range if necessary
		if currentRange < minRange {
			minRange = currentRange
			rangeStart = minElement.Value
			rangeEnd = maxVal
		}

		// If there's a next element in the same list, add it to the heap
		if minElement.ElemIndex+1 < len(nums[minElement.ListIndex]) {
			nextElement := ListElement{
				Value:     nums[minElement.ListIndex][minElement.ElemIndex+1],
				ListIndex: minElement.ListIndex,
				ElemIndex: minElement.ElemIndex + 1,
			}
			heap.Push(h, nextElement)

			// Update max value if necessary
			maxVal = max(maxVal, nextElement.Value)
		}
	}

	return []int{rangeStart, rangeEnd}
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SmallestRangeOptimized finds the smallest range with an alternative approach
// This implementation tracks next elements differently and might be more efficient
// in some cases
// Time Complexity: O(N log K) where N is the total number of elements across all lists
// and K is the number of lists
// Space Complexity: O(K) for the heap and pointers
func SmallestRangeOptimized(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{0, 0}
	}

	k := len(nums) // Number of lists

	// Create a min heap to store the current elements
	h := &ListElementMinHeap{}
	heap.Init(h)

	// Track the maximum value among current elements
	currentMax := math.MinInt32

	// Insert the first element from each list into the heap
	for i := 0; i < k; i++ {
		if len(nums[i]) > 0 {
			heap.Push(h, ListElement{
				Value:     nums[i][0],
				ListIndex: i,
				ElemIndex: 0,
			})
			currentMax = max(currentMax, nums[i][0])
		}
	}

	// Initialize result with a large range
	rangeStart, rangeEnd := 0, math.MaxInt32

	// Process the heap until any list is exhausted
	for h.Len() == k {
		// Get the minimum element
		minElement := heap.Pop(h).(ListElement)

		// Calculate the current range
		if currentMax-minElement.Value < rangeEnd-rangeStart {
			rangeStart = minElement.Value
			rangeEnd = currentMax
		}

		// Move to the next element in the list
		listIdx := minElement.ListIndex
		elemIdx := minElement.ElemIndex + 1

		// If we've reached the end of this list, we're done
		if elemIdx >= len(nums[listIdx]) {
			break
		}

		// Add the next element to the heap
		heap.Push(h, ListElement{
			Value:     nums[listIdx][elemIdx],
			ListIndex: listIdx,
			ElemIndex: elemIdx,
		})

		// Update the current maximum
		currentMax = max(currentMax, nums[listIdx][elemIdx])
	}

	return []int{rangeStart, rangeEnd}
}

// Example usage:
//
// Input: nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
// Output: [20,24]
// Explanation:
// List 1: [4, 10, 15, 24, 26], 24 is in range [20,24].
// List 2: [0, 9, 12, 20], 20 is in range [20,24].
// List 3: [5, 18, 22, 30], 22 is in range [20,24].
//
// Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
// Output: [1,1]
// Explanation: The smallest range here is [1,1], covering one element from each list.
//
// Input: nums = [[10,10],[11,11]]
// Output: [10,11]
// Explanation: The range [10,11] covers 10 from the first list and 11 from the second list.
//
// Input: nums = [[10],[11]]
// Output: [10,11]
// Explanation: The range [10,11] covers 10 from the first list and 11 from the second list.
