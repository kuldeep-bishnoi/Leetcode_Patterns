package binarysearch

// OrderAgnosticBinarySearch performs binary search on a sorted array
// without knowing if it's sorted in ascending or descending order
// Time Complexity: O(log n)
// Space Complexity: O(1)
func OrderAgnosticBinarySearch(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}

	start, end := 0, len(arr)-1

	// Determine if the array is sorted in ascending or descending order
	isAscending := arr[start] < arr[end]

	for start <= end {
		// Calculate mid point to avoid integer overflow
		mid := start + (end-start)/2

		// If the key is found at mid, return mid
		if arr[mid] == key {
			return mid
		}

		if isAscending {
			// Ascending order binary search
			if key < arr[mid] {
				end = mid - 1 // search in the left half
			} else {
				start = mid + 1 // search in the right half
			}
		} else {
			// Descending order binary search
			if key > arr[mid] {
				end = mid - 1 // search in the left half
			} else {
				start = mid + 1 // search in the right half
			}
		}
	}

	// Key not found
	return -1
}

// OrderAgnosticBinarySearchRecursive performs binary search recursively
// on a sorted array without knowing if it's ascending or descending
// Time Complexity: O(log n)
// Space Complexity: O(log n) due to recursion stack
func OrderAgnosticBinarySearchRecursive(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}

	isAscending := arr[0] < arr[len(arr)-1]
	return orderAgnosticBSRecursiveHelper(arr, key, 0, len(arr)-1, isAscending)
}

// orderAgnosticBSRecursiveHelper is a helper function for recursive binary search
func orderAgnosticBSRecursiveHelper(arr []int, key, start, end int, isAscending bool) int {
	// Base case: element not found
	if start > end {
		return -1
	}

	// Calculate mid point
	mid := start + (end-start)/2

	// If the key is found at mid, return mid
	if arr[mid] == key {
		return mid
	}

	if isAscending {
		// Ascending order binary search
		if key < arr[mid] {
			return orderAgnosticBSRecursiveHelper(arr, key, start, mid-1, isAscending)
		}
		return orderAgnosticBSRecursiveHelper(arr, key, mid+1, end, isAscending)
	}

	// Descending order binary search
	if key > arr[mid] {
		return orderAgnosticBSRecursiveHelper(arr, key, start, mid-1, isAscending)
	}
	return orderAgnosticBSRecursiveHelper(arr, key, mid+1, end, isAscending)
}

// Example usage:
//
// Ascending order array:
// Input: arr = [1, 2, 3, 4, 5, 6, 7], key = 5
// Output: 4 (index of 5 in the array)
//
// Descending order array:
// Input: arr = [7, 6, 5, 4, 3, 2, 1], key = 5
// Output: 2 (index of 5 in the array)
//
// Key not present:
// Input: arr = [1, 2, 3, 4, 5, 6, 7], key = 8
// Output: -1 (key not found)
