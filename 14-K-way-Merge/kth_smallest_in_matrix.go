package kwaymerge

import (
	"container/heap"
)

// MatrixElement represents an element in the matrix with its position
type MatrixElement struct {
	Value int // The value of the element
	Row   int // The row index of the element
	Col   int // The column index of the element
}

// MatrixMinHeap implementation
type MatrixMinHeap []MatrixElement

func (h MatrixMinHeap) Len() int           { return len(h) }
func (h MatrixMinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MatrixMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MatrixMinHeap) Push(x interface{}) {
	*h = append(*h, x.(MatrixElement))
}

func (h *MatrixMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KthSmallestInMatrix finds the kth smallest element in a n x n matrix
// where each row and column is sorted in ascending order
// Time Complexity: O(min(K, N) + K log(min(K, N))) where N is the row/column size
// Space Complexity: O(min(K, N)) for the heap
func KthSmallestInMatrix(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	// Create a min heap
	h := &MatrixMinHeap{}
	heap.Init(h)

	// Insert the first element from each row
	for i := 0; i < n && i < k; i++ {
		heap.Push(h, MatrixElement{
			Value: matrix[i][0],
			Row:   i,
			Col:   0,
		})
	}

	// Pop elements until we reach the kth element
	var current MatrixElement
	for i := 0; i < k; i++ {
		if h.Len() == 0 {
			break
		}

		// Get the smallest element
		current = heap.Pop(h).(MatrixElement)

		// If there's a next element in the same row, add it to the heap
		if current.Col+1 < n {
			heap.Push(h, MatrixElement{
				Value: matrix[current.Row][current.Col+1],
				Row:   current.Row,
				Col:   current.Col + 1,
			})
		}
	}

	return current.Value
}

// KthSmallestInMatrixBinarySearch uses binary search to find the kth smallest element
// Time Complexity: O(N * log(max-min)) where N is the row/column size
// and max-min is the range of values in the matrix
// Space Complexity: O(1)
func KthSmallestInMatrixBinarySearch(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	// Define search range
	left, right := matrix[0][0], matrix[n-1][n-1]

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		// Count how many elements are less than or equal to mid
		count := 0
		j := n - 1 // Start from the rightmost column

		for i := 0; i < n; i++ {
			// For each row, move left until finding an element <= mid
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += j + 1 // Add count of elements <= mid in this row
		}

		// Adjust search range
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// CountLessEqual counts the number of elements less than or equal to target in matrix
func CountLessEqual(matrix [][]int, target int) int {
	n := len(matrix)
	count := 0
	col := n - 1

	// Start from the top-right corner
	for row := 0; row < n; row++ {
		// Move left until finding an element <= target
		for col >= 0 && matrix[row][col] > target {
			col--
		}
		// Add count of elements <= target in this row
		count += col + 1
	}

	return count
}

// Example usage:
//
// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15],
// and the 8th smallest number is 13
//
// Input: matrix = [[-5]], k = 1
// Output: -5
//
// Input: matrix = [[1,2],[1,3]], k = 2
// Output: 1
