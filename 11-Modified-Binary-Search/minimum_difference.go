package binarysearch

import "math"

// FindMinDifferenceElement finds the element in the sorted array that has
// the minimum difference with the given target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMinDifferenceElement(arr []int, target int) int {
	n := len(arr)

	if n == 0 {
		return -1 // Invalid input
	}

	// Edge cases
	if target <= arr[0] {
		return arr[0] // Target is less than or equal to the first element
	}
	if target >= arr[n-1] {
		return arr[n-1] // Target is greater than or equal to the last element
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return arr[mid] // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, start > end
	// The element at index 'start' is the ceiling of the target
	// The element at index 'end' is the floor of the target
	// We need to compare both to find the one with the minimum difference

	// Check which one has the minimum difference
	if (arr[start] - target) < (target - arr[end]) {
		return arr[start]
	}
	return arr[end]
}

// FindMinDifferenceElementAbs finds the element in the sorted array that has
// the minimum absolute difference with the given target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMinDifferenceElementAbs(arr []int, target int) int {
	n := len(arr)

	if n == 0 {
		return -1 // Invalid input
	}

	// Edge cases
	if target <= arr[0] {
		return arr[0] // Target is less than or equal to the first element
	}
	if target >= arr[n-1] {
		return arr[n-1] // Target is greater than or equal to the last element
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return arr[mid] // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, start > end
	// start is the index of the ceiling and end is the index of the floor

	// Use absolute difference to find the minimum
	diffStart := int(math.Abs(float64(arr[start] - target)))
	diffEnd := int(math.Abs(float64(arr[end] - target)))

	if diffStart < diffEnd {
		return arr[start]
	}
	return arr[end]
}

// FindKClosestElements finds k elements in the sorted array that are closest to the target
// Time Complexity: O(log n + k) where n is the array size and k is the number of closest elements
// Space Complexity: O(k) for the result array
func FindKClosestElements(arr []int, k int, target int) []int {
	n := len(arr)

	if n == 0 || k <= 0 {
		return []int{} // Invalid input
	}

	if k >= n {
		return arr // If k is greater than or equal to array size, return the entire array
	}

	// First, find the index of the element closest to the target
	// or the floor of the target if exact match doesn't exist
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			left = mid
			break
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// At this point, right <= left
	// right is the index of the floor of the target
	// left is the index of the ceiling of the target

	// Initialize two pointers for the sliding window
	left = max(0, right)
	right = min(n-1, left+1)

	// Expand the window until we have k elements
	result := []int{}
	for right-left-1 < k {
		// If left pointer goes out of bounds, expand right
		if left < 0 {
			right++
			continue
		}

		// If right pointer goes out of bounds, expand left
		if right >= n {
			left--
			continue
		}

		// Compare the differences and decide which way to expand
		if target-arr[left] <= arr[right]-target {
			left-- // Expand left
		} else {
			right++ // Expand right
		}
	}

	// Collect the k closest elements
	for i := left + 1; i < right; i++ {
		result = append(result, arr[i])
	}

	return result
}

// Helper functions for min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage:
//
// Find Minimum Difference Element:
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 10 (10 has the minimum difference with 12)
//
// Input: arr = [4, 6, 10], target = 7
// Output: 6 (6 has the minimum difference with 7)
//
// Find K Closest Elements:
// Input: arr = [1, 2, 3, 4, 5], k = 4, target = 3
// Output: [1, 2, 3, 4] (the 4 closest elements to 3)
//
// Input: arr = [1, 2, 3, 4, 5], k = 4, target = -1
// Output: [1, 2, 3, 4] (the 4 closest elements to -1)
