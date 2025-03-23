package kwaymerge

import (
	"container/heap"
	"sort"
)

// Element with distance represents an element with its distance to target
type ElementWithDistance struct {
	Value    int
	Distance int
}

// MaxDistanceHeap is a max heap based on distance
type MaxDistanceHeap []ElementWithDistance

func (h MaxDistanceHeap) Len() int           { return len(h) }
func (h MaxDistanceHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance } // Max heap
func (h MaxDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(ElementWithDistance))
}

func (h *MaxDistanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindKClosestElements finds the k closest elements to x in the array arr
// Time Complexity: O(N log K) where N is the array size and K is the parameter
// Space Complexity: O(K) for the heap
func FindKClosestElements(arr []int, k int, x int) []int {
	// Create a max heap
	h := &MaxDistanceHeap{}
	heap.Init(h)

	// Process each element in the array
	for _, num := range arr {
		// Calculate the distance to x
		distance := abs(num - x)

		// If heap size is less than k, add the element
		if h.Len() < k {
			heap.Push(h, ElementWithDistance{Value: num, Distance: distance})
		} else if distance < (*h)[0].Distance ||
			(distance == (*h)[0].Distance && num < (*h)[0].Value) {
			// If current element is closer than the farthest in heap,
			// or if the distance is the same but the value is smaller,
			// remove the farthest and add the current element
			heap.Pop(h)
			heap.Push(h, ElementWithDistance{Value: num, Distance: distance})
		}
	}

	// Extract the k closest elements from the heap
	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		element := heap.Pop(h).(ElementWithDistance)
		result[i] = element.Value
	}

	// Sort the result
	sort.Ints(result)

	return result
}

// FindKClosestElementsBinarySearch finds the k closest elements using binary search
// Time Complexity: O(log N + K) where N is the array size and K is the parameter
// Space Complexity: O(K) for the result array
func FindKClosestElementsBinarySearch(arr []int, k int, x int) []int {
	// Edge cases
	if len(arr) == k {
		return arr
	}

	// Binary search to find the closest index to x
	left, right := 0, len(arr)-k

	for left < right {
		mid := left + (right-left)/2

		// Compare distances
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Return k elements starting from the found index
	return arr[left : left+k]
}

// FindKClosestElementsTwoPointers uses two pointers approach
// Time Complexity: O(N - K) where N is the array size and K is the parameter
// Space Complexity: O(K) for the result array
func FindKClosestElementsTwoPointers(arr []int, k int, x int) []int {
	// Initialize two pointers
	left, right := 0, len(arr)-1

	// Shrink the window until we have k elements
	for right-left >= k {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}

	// Return the k elements in the window
	return arr[left : right+1]
}

// abs returns the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Example usage:
//
// Input: arr = [1,2,3,4,5], k = 4, x = 3
// Output: [1,2,3,4]
// Explanation: The 4 closest elements to 3 are 1, 2, 3, and 4.
//
// Input: arr = [1,2,3,4,5], k = 4, x = -1
// Output: [1,2,3,4]
// Explanation: The 4 closest elements to -1 are 1, 2, 3, and 4.
//
// Input: arr = [1,1,1,10,10,10], k = 1, x = 9
// Output: [10]
// Explanation: The closest element to 9 is 10.
