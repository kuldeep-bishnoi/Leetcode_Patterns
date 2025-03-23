package binarysearch

// FindMaxInBitonicArray finds the maximum value in a bitonic array
// A bitonic array is an array that monotonically increases and then decreases
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMaxInBitonicArray(arr []int) int {
	if len(arr) == 0 {
		return -1 // Invalid input
	}

	if len(arr) == 1 {
		return arr[0] // Only one element
	}

	start, end := 0, len(arr)-1

	for start < end {
		mid := start + (end-start)/2

		// Check if we're in the increasing part of the array
		if arr[mid] < arr[mid+1] {
			start = mid + 1 // Move towards the peak
		} else {
			end = mid // Mid could be the peak or in the decreasing part
		}
	}

	// At this point, start == end, pointing to the peak
	return arr[start]
}

// FindMaxIndexInBitonicArray finds the index of the maximum value in a bitonic array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMaxIndexInBitonicArray(arr []int) int {
	if len(arr) == 0 {
		return -1 // Invalid input
	}

	if len(arr) == 1 {
		return 0 // Only one element
	}

	start, end := 0, len(arr)-1

	for start < end {
		mid := start + (end-start)/2

		// Check if we're in the increasing part of the array
		if arr[mid] < arr[mid+1] {
			start = mid + 1 // Move towards the peak
		} else {
			end = mid // Mid could be the peak or in the decreasing part
		}
	}

	// At this point, start == end, pointing to the peak
	return start
}

// SearchInBitonicArray searches for a target in a bitonic array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchInBitonicArray(arr []int, target int) int {
	if len(arr) == 0 {
		return -1 // Invalid input
	}

	// First, find the peak of the bitonic array
	peakIndex := FindMaxIndexInBitonicArray(arr)

	// If the target is the peak element, return its index
	if arr[peakIndex] == target {
		return peakIndex
	}

	// Search in the ascending part (before peak)
	result := orderAgnosticBinarySearch(arr, 0, peakIndex-1, target, true)

	// If not found in the ascending part, search in the descending part (after peak)
	if result == -1 {
		result = orderAgnosticBinarySearch(arr, peakIndex+1, len(arr)-1, target, false)
	}

	return result
}

// orderAgnosticBinarySearch performs binary search on a sorted array segment
// isAscending: whether the array segment is sorted in ascending order
func orderAgnosticBinarySearch(arr []int, start, end, target int, isAscending bool) int {
	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid // Target found
		}

		if isAscending {
			// Ascending order binary search
			if arr[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			// Descending order binary search
			if arr[mid] > target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1 // Target not found
}

// IsBitonicArray checks if an array is bitonic
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsBitonicArray(arr []int) bool {
	if len(arr) <= 2 {
		return true // Arrays with 0, 1, or 2 elements are trivially bitonic
	}

	// First, find where the array stops increasing
	i := 1
	for i < len(arr) && arr[i-1] < arr[i] {
		i++
	}

	// If we reached the end, it's a strictly increasing array (bitonic)
	if i == len(arr) {
		return true
	}

	// Continue and check if the array is strictly decreasing after the peak
	for i < len(arr) && arr[i-1] > arr[i] {
		i++
	}

	// If we reached the end, it's bitonic
	return i == len(arr)
}

// Example usage:
//
// Find Maximum in Bitonic Array:
// Input: arr = [1, 3, 8, 12, 4, 2]
// Output: 12 (maximum value in the bitonic array)
//
// Find Maximum Index in Bitonic Array:
// Input: arr = [1, 3, 8, 12, 4, 2]
// Output: 3 (index of the maximum value 12)
//
// Search in Bitonic Array:
// Input: arr = [1, 3, 8, 12, 4, 2], target = 4
// Output: 4 (index of 4 in the array)
//
// Input: arr = [1, 3, 8, 12, 4, 2], target = 10
// Output: -1 (target not found)
//
// Is Bitonic Array:
// Input: arr = [1, 3, 8, 12, 4, 2]
// Output: true (the array is bitonic: it increases and then decreases)
