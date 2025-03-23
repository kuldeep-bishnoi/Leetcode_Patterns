package twoheaps

import (
	"container/heap"
)

// Interval represents a start and end interval
type Interval struct {
	Start int
	End   int
	Index int // Original index in the input array
}

// StartMinHeap implements a min heap for intervals, sorted by start time
type StartMinHeap []Interval

func (h StartMinHeap) Len() int           { return len(h) }
func (h StartMinHeap) Less(i, j int) bool { return h[i].Start < h[j].Start }
func (h StartMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StartMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *StartMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *StartMinHeap) Peek() Interval {
	if h.Len() == 0 {
		return Interval{}
	}
	return (*h)[0]
}

// EndMaxHeap implements a max heap for intervals, sorted by end time
type EndMaxHeap []Interval

func (h EndMaxHeap) Len() int           { return len(h) }
func (h EndMaxHeap) Less(i, j int) bool { return h[i].End > h[j].End } // We want max heap for end times
func (h EndMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EndMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *EndMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *EndMaxHeap) Peek() Interval {
	if h.Len() == 0 {
		return Interval{}
	}
	return (*h)[0]
}

// FindRightInterval finds the right interval for each interval in the input
// Time Complexity: O(n log n) where n is the number of intervals
// Space Complexity: O(n) for storing intervals in heaps
func FindRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	// Initialize heaps
	startMinHeap := &StartMinHeap{}
	endMaxHeap := &EndMaxHeap{}
	heap.Init(startMinHeap)
	heap.Init(endMaxHeap)

	// Fill heaps with all intervals
	for i := 0; i < n; i++ {
		interval := Interval{
			Start: intervals[i][0],
			End:   intervals[i][1],
			Index: i,
		}
		heap.Push(startMinHeap, interval)
		heap.Push(endMaxHeap, interval)
	}

	// Create the result array with -1 as default for each interval
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}

	// Process all intervals
	for endMaxHeap.Len() > 0 {
		// Get interval with the largest end time
		endInterval := heap.Pop(endMaxHeap).(Interval)

		// If the largest start time is less than end time, no right interval exists
		if startMinHeap.Len() == 0 || startMinHeap.Peek().Start < endInterval.End {
			result[endInterval.Index] = -1
			continue
		}

		// Find the first interval that starts after the end time
		// We need to pop intervals from the start heap until we find one
		// that meets our criteria, then reinsert them back
		var rightInterval Interval
		tempHeap := &StartMinHeap{}
		heap.Init(tempHeap)

		found := false
		for startMinHeap.Len() > 0 && !found {
			interval := heap.Pop(startMinHeap).(Interval)
			if interval.Start >= endInterval.End {
				rightInterval = interval
				found = true
			}
			heap.Push(tempHeap, interval)
		}

		// Put all intervals back into the start heap
		for tempHeap.Len() > 0 {
			heap.Push(startMinHeap, heap.Pop(tempHeap))
		}

		if found {
			result[endInterval.Index] = rightInterval.Index
		}
	}

	return result
}

// FindRightIntervalOptimized provides a more efficient solution using sorting
// Time Complexity: O(n log n) where n is the number of intervals
// Space Complexity: O(n) for storing intervals and results
func FindRightIntervalOptimized(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	// Create a slice of intervals with their original indices
	indexedIntervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		indexedIntervals[i] = Interval{
			Start: intervals[i][0],
			End:   intervals[i][1],
			Index: i,
		}
	}

	// Sort the intervals by start time
	startHeap := &StartMinHeap{}
	heap.Init(startHeap)
	for _, interval := range indexedIntervals {
		heap.Push(startHeap, interval)
	}

	// Create a sorted slice of intervals by start time
	sortedByStart := make([]Interval, 0, n)
	for startHeap.Len() > 0 {
		sortedByStart = append(sortedByStart, heap.Pop(startHeap).(Interval))
	}

	// Initialize result array
	result := make([]int, n)

	// For each interval, find its right interval using binary search
	for _, interval := range indexedIntervals {
		result[interval.Index] = binarySearch(sortedByStart, interval.End)
	}

	return result
}

// binarySearch finds the index of the first interval with start >= target
// If not found, returns -1
func binarySearch(intervals []Interval, target int) int {
	left, right := 0, len(intervals)-1
	answer := -1

	for left <= right {
		mid := left + (right-left)/2
		if intervals[mid].Start >= target {
			answer = intervals[mid].Index
			right = mid - 1 // Continue searching for a smaller index
		} else {
			left = mid + 1
		}
	}

	return answer
}

// Example usage:
//
// Input: intervals = [[1,2]]
// Output: [-1]
// Explanation: There is only one interval in the collection, so it outputs -1.
//
// Input: intervals = [[3,4],[2,3],[1,2]]
// Output: [-1,0,1]
// Explanation:
// For [3,4], there is no right interval.
// For [2,3], the right interval is [3,4] with index 0.
// For [1,2], the right interval is [2,3] with index 1.
//
// Input: intervals = [[1,4],[2,3],[3,4]]
// Output: [-1,2,-1]
// Explanation:
// For [1,4], there is no right interval.
// For [2,3], the right interval is [3,4] with index 2.
// For [3,4], there is no right interval.
