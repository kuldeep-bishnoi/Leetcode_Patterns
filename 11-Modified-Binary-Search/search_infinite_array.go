package binarysearch

// ArrayReader interface provides a way to access an array of unknown size
type ArrayReader interface {
	Get(index int) int // Returns the value at index, or a large number if index is out of bounds
}

// concreteArrayReader is a concrete implementation of ArrayReader for testing
type concreteArrayReader struct {
	arr []int
}

// Get returns the value at index, or a large number if index is out of bounds
func (r *concreteArrayReader) Get(index int) int {
	if index >= len(r.arr) {
		return 1<<31 - 1 // Return a very large number (INT_MAX)
	}
	return r.arr[index]
}

// NewArrayReader creates a new ArrayReader from a given array
func NewArrayReader(arr []int) ArrayReader {
	return &concreteArrayReader{arr: arr}
}

// SearchInInfiniteArray searches for a target in a sorted array of unknown size
// Time Complexity: O(log n) where n is the position of the target
// Space Complexity: O(1)
func SearchInInfiniteArray(reader ArrayReader, target int) int {
	// First, find a range where the target might exist
	// Start with a small range and exponentially increase it
	start := 0
	end := 1

	// Double the range until we find a range that might contain the target
	// We know we've gone too far when the element at 'end' is greater than the target
	for reader.Get(end) < target {
		newStart := end + 1
		end = end + (end-start+1)*2 // Double the size of the range
		start = newStart
	}

	// Now perform binary search in the range [start, end]
	return binarySearchInReader(reader, target, start, end)
}

// binarySearchInReader performs binary search in a specified range of the ArrayReader
func binarySearchInReader(reader ArrayReader, target, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		midVal := reader.Get(mid)

		if midVal == target {
			return mid
		}

		if midVal > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1 // Target not found
}

// SearchInInfiniteArrayAlternative is an alternative approach using a simpler
// range-finding mechanism
// Time Complexity: O(log n) where n is the position of the target
// Space Complexity: O(1)
func SearchInInfiniteArrayAlternative(reader ArrayReader, target int) int {
	// Find the bounds
	left := 0
	right := 1

	// While the target is greater than the element at 'right'
	for reader.Get(right) < target {
		left = right
		right *= 2 // Exponentially increase the bound
	}

	// Perform binary search
	return binarySearchInReader(reader, target, left, right)
}

// Example usage:
//
// Input: array = [4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30], target = 16
// Output: 6 (index of 16 in the array)
//
// Input: array = [4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30], target = 11
// Output: -1 (target not found)
//
// To use these functions:
// 1. Create an ArrayReader: reader := NewArrayReader([]int{4, 6, 8, 10, 12, 14, 16, 18, 20})
// 2. Call the search function: result := SearchInInfiniteArray(reader, 10)
// The result will be the index of the target in the array, or -1 if not found
