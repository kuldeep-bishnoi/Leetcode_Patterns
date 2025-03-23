package topk

import (
	"container/heap"
	"math"
)

// Point represents a point in 2D space
type Point struct {
	X, Y     int
	Distance float64
}

// MaxDistanceHeap is a max heap based on distance from origin
type MaxDistanceHeap []Point

func (h MaxDistanceHeap) Len() int           { return len(h) }
func (h MaxDistanceHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance } // Max heap
func (h MaxDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxDistanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KClosestPoints finds the k closest points to the origin (0, 0)
// Time Complexity: O(N log K) where N is the number of points and K is the parameter
// Space Complexity: O(K) for storing the heap
func KClosestPoints(points [][]int, k int) [][]int {
	// Create a max heap based on distance from origin
	h := &MaxDistanceHeap{}
	heap.Init(h)

	// Process each point
	for _, point := range points {
		x, y := point[0], point[1]
		// Calculate Euclidean distance from origin
		distance := math.Sqrt(float64(x*x + y*y))

		// If heap size is less than k, add the point
		if h.Len() < k {
			heap.Push(h, Point{X: x, Y: y, Distance: distance})
		} else if distance < (*h)[0].Distance {
			// If current point is closer than the farthest point in heap,
			// remove the farthest and add the current point
			heap.Pop(h)
			heap.Push(h, Point{X: x, Y: y, Distance: distance})
		}
	}

	// Extract the k closest points from the heap
	result := make([][]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		point := heap.Pop(h).(Point)
		result[i] = []int{point.X, point.Y}
	}

	return result
}

// KClosestPointsQuickSelect uses the QuickSelect algorithm to find the K closest points
// Time Complexity: O(N) average case, O(N²) worst case
// Space Complexity: O(1) for in-place partitioning (excluding output array)
func KClosestPointsQuickSelect(points [][]int, k int) [][]int {
	// Create a slice to store distances
	distances := make([]Point, len(points))
	for i, point := range points {
		x, y := point[0], point[1]
		distance := float64(x*x + y*y) // Using distance squared to avoid sqrt (not necessary for comparison)
		distances[i] = Point{X: x, Y: y, Distance: distance}
	}

	// Use quickSelect to find the K closest points
	quickSelectPoints(distances, 0, len(distances)-1, k)

	// Extract the k closest points
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = []int{distances[i].X, distances[i].Y}
	}

	return result
}

// quickSelectPoints partitions the points around a pivot until the kth closest point is in its position
func quickSelectPoints(points []Point, start, end, k int) {
	if start >= end {
		return
	}

	// Partition the array and get the pivot index
	pivotIndex := partitionPoints(points, start, end)

	// If the pivot is at the kth position, we're done
	if pivotIndex == k-1 {
		return
	} else if pivotIndex > k-1 {
		// If pivot is after k, search left side
		quickSelectPoints(points, start, pivotIndex-1, k)
	} else {
		// If pivot is before k, search right side
		quickSelectPoints(points, pivotIndex+1, end, k)
	}
}

// partitionPoints partitions the points array around a pivot
func partitionPoints(points []Point, start, end int) int {
	// Choose the last element as pivot
	pivot := points[end].Distance

	// Index of smaller element
	i := start

	// Traverse through all elements
	for j := start; j < end; j++ {
		// If current element is smaller than the pivot
		if points[j].Distance < pivot {
			// Swap elements
			points[i], points[j] = points[j], points[i]
			i++
		}
	}

	// Swap pivot element with element at i
	points[i], points[end] = points[end], points[i]

	return i
}

// Example usage:
//
// Input: points = [[1,3],[-2,2]], k = 1
// Output: [[-2,2]]
// Explanation: The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
//
// Input: points = [[3,3],[5,-1],[-2,4]], k = 2
// Output: [[3,3],[-2,4]]
// The closest 2 points are [3,3] and [-2,4].
