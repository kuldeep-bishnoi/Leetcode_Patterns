package kwaymerge

import (
	"container/heap"
)

// Element represents an element in one of the arrays
type Element struct {
	Value        int // The actual value
	ArrayIndex   int // Index of the array this element belongs to
	ElementIndex int // Index of this element in its array
}

// MinHeap implementation
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKSortedArrays merges K sorted arrays into one sorted array
// Time Complexity: O(N log K) where N is the total number of elements across all arrays
// and K is the number of arrays
// Space Complexity: O(K) for storing the heap with one element from each of the K arrays
func MergeKSortedArrays(arrays [][]int) []int {
	// Calculate the total size for the result array
	totalSize := 0
	for _, array := range arrays {
		totalSize += len(array)
	}

	// Initialize the result array with the calculated size
	result := make([]int, 0, totalSize)

	// Create a min heap
	h := &MinHeap{}
	heap.Init(h)

	// Insert the first element from each array into the heap
	for i, array := range arrays {
		if len(array) > 0 {
			heap.Push(h, Element{
				Value:        array[0],
				ArrayIndex:   i,
				ElementIndex: 0,
			})
		}
	}

	// Extract the minimum element and add the next element from the same array
	for h.Len() > 0 {
		// Get the minimum element
		minElement := heap.Pop(h).(Element)
		result = append(result, minElement.Value)

		// If there are more elements in the same array, add the next one
		if minElement.ElementIndex+1 < len(arrays[minElement.ArrayIndex]) {
			heap.Push(h, Element{
				Value:        arrays[minElement.ArrayIndex][minElement.ElementIndex+1],
				ArrayIndex:   minElement.ArrayIndex,
				ElementIndex: minElement.ElementIndex + 1,
			})
		}
	}

	return result
}

// MergeKSortedArraysBruteForce merges K sorted arrays using a simpler approach
// This is less efficient but easier to understand
// Time Complexity: O(N log N) where N is the total number of elements
// Space Complexity: O(N) for storing all elements
func MergeKSortedArraysBruteForce(arrays [][]int) []int {
	// Gather all elements into a single array
	allElements := []int{}
	for _, array := range arrays {
		allElements = append(allElements, array...)
	}

	// Sort the combined array
	// Note: In real implementation, use the sort package
	// This is a placeholder for the sorting operation
	// sort.Ints(allElements)

	// Instead, let's implement a quick sort for demo purposes
	quickSort(allElements, 0, len(allElements)-1)

	return allElements
}

// quickSort is a helper function for sorting an array
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array
		pi := partition(arr, low, high)

		// Sort the partitions
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

// partition is a helper function for quickSort
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Example usage:
//
// Input: arrays = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
//
// Input: arrays = [[]]
// Output: []
//
// Input: arrays = [[1]]
// Output: [1]
